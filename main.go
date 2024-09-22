package main

import (
	"fmt"
)

func main() {
	video := "./clips/ravus1.mp4"
	audio := "./clips/ravus1.wav"
	vttFile := "./clips/ravus1.vtt"
	srtFile := "./clips/ravus1.srt"
	outputVideo := "./clips/ravus1_with_subtitles.mp4"

	// Extract audio from video
	fmt.Println("Extracting audio...")
	if err := extractAudioFromVideo(video, audio); err != nil {
		fmt.Println("Error extracting audio:", err)
		return
	}

	// Convert audio to text (WebVTT format)
	fmt.Println("Converting speech to text...")
	actualVTTFile, err := speechToText(audio, vttFile)
	if err != nil {
		fmt.Println("Error converting speech to text:", err)
		return
	}

	// Convert WebVTT to SRT
	fmt.Println("Converting WebVTT to SRT...")
	if err := convertVTTtoSRT(actualVTTFile, srtFile); err != nil {
		fmt.Println("Error converting WebVTT to SRT:", err)
		return
	}

	// Add subtitles to video
	fmt.Println("Adding subtitles to video...")
	if err := addSubtitlesToVideo(video, srtFile, outputVideo); err != nil {
		fmt.Println("Error adding subtitles to video:", err)
		return
	}

	fmt.Println("Process completed successfully!")
}
