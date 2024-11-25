Simple go app which allows you to directly convert a mp4 to a gif with resolution 320xAuto and 8 fps.

Build on macOS

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
export PATH=$PATH:$HOME/go/bin
fyne package -os darwin -icon icon.png --tags="vcs=false"
```

Build on Windows

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os windows -icon icon.png
```
