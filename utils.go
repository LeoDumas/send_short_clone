package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Extract audio from video
func extractAudioFromVideo(videoPath string, audioPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ar", "16000", "-q:a", "0", "-map", "a", audioPath)
	return cmd.Run()
}

// Convert audio to text using Whisper
func speechToText(audioPath string, vttPath string) (string, error) {
	whisperCmd := fmt.Sprintf("/home/leo/whisper.cpp/build/bin/main -m /home/leo/whisper.cpp/models/ggml-medium.bin -f %s --output-file %s --language fr -ovtt",
		audioPath, vttPath)
	cmd := exec.Command("wsl", "bash", "-c", whisperCmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command output:", string(output))
		return "", err
	}

	actualVTTPath := vttPath + ".vtt"
	if _, err := os.Stat(actualVTTPath); err == nil {
		return actualVTTPath, nil
	}

	return vttPath, nil
}

func convertVTTtoSRT(vttPath, srtPath string) error {
	vttFile, err := os.Open(vttPath)
	if err != nil {
		return err
	}
	defer vttFile.Close()

	srtFile, err := os.Create(srtPath)
	if err != nil {
		return err
	}
	defer srtFile.Close()

	scanner := bufio.NewScanner(vttFile)
	writer := bufio.NewWriter(srtFile)
	lineCount := 1
	isSubtitle := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-->") {
			isSubtitle = true
			fmt.Fprintf(writer, "%d\n", lineCount)
			fmt.Fprintf(writer, "%s\n", convertVTTTimeToSRTTime(line))
		} else if isSubtitle && line == "" {
			fmt.Fprintf(writer, "\n")
			lineCount++
			isSubtitle = false
		} else if isSubtitle {
			fmt.Fprintf(writer, "%s\n", line)
		}
	}

	writer.Flush()
	return scanner.Err()
}

func convertVTTTimeToSRTTime(vttTime string) string {
	return strings.Replace(vttTime, ".", ",", -1)
}

func addSubtitlesToVideo(inputVideo, subtitlesFile, outputVideo string) error {
	cmd := exec.Command("ffmpeg",
		"-i", inputVideo,
		"-vf", fmt.Sprintf("subtitles=%s:force_style='FontSize=24,FontName=Arial'", subtitlesFile),
		"-c:a", "copy",
		outputVideo)
	return cmd.Run()
}
