package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/srlehn/termimg/internal/assets"
	"github.com/srlehn/termimg/tui/termuiimg"
)

var (
	imgTermui  = widgets.NewImage
	imgTermImg = termuiimg.NewImage
	imfn       = imgTermImg
)

func main() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	m, _, _ := image.Decode(bytes.NewReader(assets.SnakePic))

	var p termui.Drawable = imfn(m)
	p.SetRect(10, 8, 50, 30)

	termui.Render(p)

	if pc, ok := p.(interface{ Close() error }); ok {
		_ = pc.Close()
	}

	for e := range termui.PollEvents() {
		if e.Type == termui.KeyboardEvent {
			break
		}
	}
}
