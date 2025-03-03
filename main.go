package main

import (
	"image/color"
	"os"
	"time"

	g "github.com/AllenDang/giu"
)

var (
	peterTexture = &g.ReflectiveBoundTexture{}
)

func onClickMe() {
	os.Exit(0)
}

func loop() {
	g.PushStyleColor(g.StyleColorWindowBg, color.White)
	g.PushStyleColor(g.StyleColorBorder, color.Transparent)
	g.SingleWindow().Layout(
		g.Column(
			g.Align(g.AlignCenter).To(
				g.Dummy(0, 19),
				g.Style().SetColor(g.StyleColorBorder, color.Transparent).To(
					peterTexture.ToImageWidget().Scale(1.3, 1.3),
				),
				g.Dummy(0, 8),
				g.Style().SetColor(g.StyleColorButton, color.RGBA{222, 222, 222, 255}).SetColor(g.StyleColorText, color.Black).To(
					g.Button("OK").Size(98, 29).OnClick(onClickMe),
				),
			),
		),
	)
	g.PopStyleColor()
	g.PopStyleColor()
}

func main() {
	wnd := g.NewMasterWindow("Peter Alert", 232, 148, g.MasterWindowFlagsNotResizable)
	g.Context.FontAtlas.SetDefaultFont("Inter.ttc", 16)

	peterTexture.SetSurfaceFromURL("https://github.com/kokoscript/PeterAlert/blob/main/BeterAlert/peter.png?raw=true", 5*time.Second, false)

	wnd.Run(loop)
}
