package main

import (
	"fmt"
	"os/exec"
)

// Extract audio from video
func extractAudioFromVideo(videoPath string, audioPath string) error {
	// ffmpeg command: ffmpeg -i input.mp4 -ar 16000 -q:a 0 -map a output.wav
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ar", "16000", "-q:a", "0", "-map", "a", audioPath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// Convert audio to text using Whisper
func speechToText(audioPath string, textPath string) error {
	// Construct the command to run inside WSL
	whisperCmd := fmt.Sprintf("/home/leo/whisper.cpp/build/bin/main -m /home/leo/whisper.cpp/models/ggml-medium.bin -f %s --output-file %s --language fr -otxt",
		audioPath, textPath)

	// Execute the command inside WSL
	cmd := exec.Command("wsl", "bash", "-c", whisperCmd)
	output, err := cmd.CombinedOutput() // Capture both stdout and stderr
	if err != nil {
		fmt.Println("Command output:", string(output)) // Print the output for debugging
		return err
	}
	return nil
}

func main() {
	video := "./clips/ravus1.mp4"
	audio := "./clips/ravus1.wav"
	transcript := "./clips/ravus1.txt"

	// Extract audio from video
	err := extractAudioFromVideo(video, audio)
	if err != nil {
		fmt.Println("Error extracting audio:", err)
		return
	}
	fmt.Println("Audio extracted successfully")

	// Convert audio to text
	err = speechToText(audio, transcript)
	if err != nil {
		fmt.Println("Error converting speech to text:", err)
		return
	}
	fmt.Println("Speech converted to text successfully")
}
