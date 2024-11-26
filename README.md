Simple go app which allows you to directly convert a mp4 to a gif with resolution 320xAuto and 8 fps.

Build on macOS

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
export PATH=$PATH:$HOME/go/bin
yne package -os darwin -icon icon.png -appID com.eliasalerno.videotogif
```

Build on Windows

```sh
go mod tidy
go install fyne.io/fyne/v2/cmd/fyne@latest
```

> you need to add the go bin to PATH before building

```sh
fyne package -os windows -icon icon.png
```
