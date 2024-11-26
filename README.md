![icon](icon.png)

# Gifi

Simple go app which allows you to directly convert a mp4 to a gif and save it in same folder.

Build for macOS

```sh
brew install ffmpeg
```

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
export PATH=$PATH:$HOME/go/bin
fyne package -os darwin -icon icon.png -appID com.eliasalerno.videotogif
```

Build for Windows (on macOS)

```sh
# Install MinGW-w64 & FFmpeg using Homebrew
brew install mingw-w64
brew install ffmpeg

# Set the environment variables for cross-compilation
export CC=x86_64-w64-mingw32-gcc
export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64
```

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest

export PATH=$PATH:$HOME/go/bin

# Run the fyne package command
fyne package -os windows -icon icon.png
```

## Dependencies

- FFmpeg
- Fyne CLI
- gcc
