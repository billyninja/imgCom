package main

import (
    "github.com/billyninja/imgCom/xmg"
)

func Test(sman *xmg.SurfaceManager, fman *xmg.FontManager) {
    cmp := &xmg.Composition{
        ImageStr: "strix-nebulosa.jpg",
        Gfx: []*xmg.GfxEl{
            {
                GfxStr: "test.png",
                Pos: &xmg.Pos{
                    X: 0,
                    Y: 0,
                },
            },
        },
        Text: []*xmg.TextEl{
            {
                FontStr:  "test.ttf",
                FontSize: 10,
                Message:  "Hello world!",
                Pos: &xmg.Pos{
                    X: 0,
                    Y: 0,
                },
            },
        },
    }

    xmg.Render(cmp)
}

func main() {
    sman := xmg.NewSurfaceManager("sample_media/images", "")
    fman := xmg.NewFontManager("sample_media/fonts", "Go-Regular.ttf")

    Test(sman, fman)
}
