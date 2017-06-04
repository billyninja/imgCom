package main

import (
    "fmt"
    "github.com/billyninja/imgCom/xmg"
    "github.com/veandco/go-sdl2/sdl"
    "time"
)

func Test(sman *xmg.SurfaceManager, fman *xmg.FontManager) {
    cmp := &xmg.Composition{
        MainImage: &xmg.GfxEl{
            GfxStr: "strix-nebulosa.jpg",
            Scale: &xmg.Scale{
                W: 200,
                H: 100,
            },
            Pos: &xmg.Pos{X: 100, Y: 100},
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
                GfxStr: "trophy1600.png",
                Pos: &xmg.Pos{
                    X: 0,
                    Y: 0,
                },
            },
            {
                GfxStr: "trophy1600.png",
                Scale: &xmg.Scale{
                    W: 24,
                    H: 24,
                },
                Pos: &xmg.Pos{
                    X: 0,
                    Y: 0,
                },
            },
        },
        Text: []*xmg.TextEl{
            {
                FontStr:  "Go-Regular.ttf",
                FontSize: 16,
                Message:  "Hello world!",
                Color:    &xmg.Color{255, 255, 255, 255},
                Pos: &xmg.Pos{
                    X: 0,
                    Y: 0,
                },
            },
        },
    }
    var surf *sdl.Surface
    for i := 0; i < 10; i++ {
        t1 := time.Now()
        surf = xmg.Render(cmp, sman, fman)
        fmt.Printf("lap %d: %v\n", i+1, time.Since(t1).Seconds())
    }
    filename := fmt.Sprintf("out/xmg-%s_%s", time.Now().Format("02012006T150405"), cmp.MainImage.GfxStr)

    xmg.Save(filename, surf)
}

func main() {
    t1 := time.Now()
    sman := xmg.NewSurfaceManager("sample_media/images", "")
    fman := xmg.NewFontManager("sample_media/fonts", "Go-Regular.ttf")
    Test(sman, fman)
    fmt.Printf("%v\n", time.Since(t1).Seconds())
}
