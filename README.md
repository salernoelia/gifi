![icon](icon.png)

# Gifi

App which allows you to directly convert a .mp4, .mov, .mkv or .avi to a gif and save it in same folder.

## Build for macOS

```sh
brew install ffmpeg
```

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
export PATH=$PATH:$HOME/go/bin
fyne package -os darwin -icon icon.png -appID com.eliasalerno.gifi -name gifi
```

## Build for Windows (on macOS)

> Hint: you must have docker installed & the daemon must be running

```sh
# Install MinGW-w64 & FFmpeg using Homebrew
brew install mingw-w64
brew install ffmpeg
```

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
go install github.com/fyne-io/fyne-cross@latest

export PATH=$PATH:$HOME/go/bin

go clean -cache -modcache

fyne-cross windows -arch=amd64 -icon icon.png -app-id com.eliasalerno.gifi -name gifi
```

## Dependencies

- FFmpeg
- Fyne CLI
- gcc

## for cross compilation

- Docker
- fyne-cross
