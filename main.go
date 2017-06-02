package main

import (
    "github.com/billyninja/imgCom/xmg"
)

func Test(sman *xmg.SurfaceManager, fman *xmg.FontManager) {
    cmp := &xmg.Composition{
        MainImage: &xmg.GfxEl{
            GfxStr: "sample_media/images/strix-nebulosa.jpg",
        },
        Dimensions: &xmg.Scale{
            W: 800,
            H: 600,
        },
        BGColor: &xmg.Color{
            R: 255, G: 0, B: 0, A: 255,
        },
        Gfx: []*xmg.GfxEl{
            {
                GfxStr: "sample_media/images/trophy1600.png",
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
