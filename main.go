package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Video to GIF Converter")

    label := widget.NewLabel("Drag and drop a video file to convert it to a GIF.")
    outputLabel := widget.NewLabel("")

    myWindow.SetContent(container.NewVBox(
        label,
        outputLabel,
    ))
    myWindow.Resize(fyne.NewSize(400, 200))

    myWindow.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
        for _, uri := range uris {
            filePath := uri.Path()
            log.Printf("File dropped: %v", filePath)
            if isVideoFile(filePath) {
                outputLabel.SetText("Converting: " + filePath)
                gifPath := convertToGif(filePath)
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

func convertToGif(inputPath string) string {
    outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".gif"
    ffmpegPath := "/usr/local/bin/ffmpeg" // Update this path as needed
    cmd := exec.Command(ffmpegPath, "-i", inputPath, "-vf", "scale=320:-1,fps=8", outputPath)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error during ffmpeg execution:", err)
        fmt.Println("ffmpeg output:", string(output))
        return ""
    }
    return outputPath
}