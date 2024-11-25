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

    // Handle file drag-and-drop
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
    return ext == ".mp4" || ext == ".mkv" || ext == ".avi" // Add more formats if needed
}

func convertToGif(inputPath string) string {
    outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".gif"
    cmd := exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=320:-1,fps=8", outputPath)
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error:", err)
        return ""
    }
    return outputPath
}