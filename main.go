package main

import (
    "github.com/billyninja/imgCom/xmg"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
)

func Test(man *xmg.Manager, rdr *sdl.Renderer) {
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
                FontStr: "test.ttf",
                Message: "Hello world!",
                Pos: &xmg.Pos{
                    X: 0,
                    Y: 0,
                },
            },
        },
    }

    cmp.LoadResources(man)
}

func main() {
    surf, rdr, _ := xmg.GetRendererFromFilename("sample_media/images/strix-nebulosa.jpg")
    _ = img.SavePNG(surf, "out/strix-nebulosa.png")
    man, _ := xmg.NewManager(
        []string{"sample_media/images"},
        []string{"sample_media/fonts"},
    )

    Test(man, rdr)
}
