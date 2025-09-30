package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("gifi")

	label := widget.NewLabel("Drag and drop a video file to convert it to a GIF.")
	outputLabel := widget.NewLabel("")

	resolutions := []string{"320", "480", "640", "800", "1024"}
	resolutionSelect := widget.NewSelect(resolutions, func(value string) {})
	resolutionSelect.SetSelected("320")

	frameRates := []string{"8 fps", "12 fps", "15 fps", "24 fps", "30 fps"}
	frameRateSelect := widget.NewSelect(frameRates, func(value string) {})
	frameRateSelect.SetSelected("8 fps")

	myWindow.SetContent(container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Resolution:"),
			resolutionSelect,
			widget.NewLabel("Frame Rate:"),
			frameRateSelect,
		),
		label,
		outputLabel,
	))
	myWindow.Resize(fyne.NewSize(500, 350))

	myWindow.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		for _, uri := range uris {
			filePath := uri.Path()
			log.Printf("File dropped: %v", filePath)
			if isVideoFile(filePath) {
				outputLabel.SetText("Converting: " + filePath)
				resolution := resolutionSelect.Selected
				frameRate := frameRateSelect.Selected
				gifPath := convertToGif(filePath, resolution, frameRate)
				if gifPath != "" {
					outputLabel.SetText("GIF saved at: " + gifPath)
				} else {
					outputLabel.SetText("Failed to convert video.")
				}
			} else {
				outputLabel.SetText("Unsupported file type.")
			}
		}
	})

	myWindow.ShowAndRun()
}

func isVideoFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".mp4" || ext == ".mov" || ext == ".mkv" || ext == ".avi"
}

func convertToGif(inputPath, resolution, frameRate string) string {
	frameRate = strings.TrimSuffix(frameRate, " fps")

	var ffmpegPath string
	if runtime.GOOS == "windows" {
		possiblePaths := []string{
			"ffmpeg.exe",
			"C:\\ffmpeg\\bin\\ffmpeg.exe",
			filepath.Join(os.Getenv("PROGRAMFILES"), "ffmpeg", "bin", "ffmpeg.exe"),
		}

		for _, path := range possiblePaths {
			if _, err := exec.LookPath(path); err == nil {
				ffmpegPath = path
				break
			}
		}

		if ffmpegPath == "" {
			fmt.Println("FFmpeg not found. Please install FFmpeg and add it to PATH")
			return ""
		}
	} else {
		ffmpegPath = "/usr/local/bin/ffmpeg"
	}

	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + "-" + resolution + "xAUTO-" + frameRate + "fps" + ".gif"
	scaleOption := fmt.Sprintf("scale=%s:-1", resolution)
	fpsOption := fmt.Sprintf("fps=%s", frameRate)

	outputPath = getNextAvailablePath(outputPath)

	cmd := exec.Command(ffmpegPath, "-i", inputPath, "-vf", fmt.Sprintf("%s,%s", scaleOption, fpsOption), outputPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error during ffmpeg execution:", err)
		fmt.Println("ffmpeg output:", string(output))
		return ""
	}
	return outputPath
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getNextAvailablePath(outputPath string) string {
	basePath := strings.TrimSuffix(outputPath, ".gif")
	i := 1

	if idx := strings.LastIndex(basePath, "-"); idx != -1 {
		if num, err := strconv.Atoi(basePath[idx+1:]); err == nil {
			i = num
			basePath = basePath[:idx]
		}
	}

	nextPath := outputPath
	for fileExists(nextPath) {
		nextPath = fmt.Sprintf("%s-%d.gif", basePath, i)
		i++
	}
	return nextPath
}
