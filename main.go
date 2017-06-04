package main

import (
    "fmt"
    "github.com/billyninja/imgCom/xmg"
    "github.com/veandco/go-sdl2/sdl"
    "time"
)

func Test(man *xmg.Manager) {
    cmp := &xmg.Composition{
        MainImage: &xmg.GfxEl{
            GfxStr: "strix-nebulosa.jpg",
            Scale: &xmg.Scale{
                W: 200,
                H: 100,
            },
            Pos: &xmg.Pos{X: 100, Y: 100},
        },
        Thumbfy: &xmg.Thumbfy{
            Spec: "52x",
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
                FontStr:  "test.ttf",
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
    var (
        surf *sdl.Surface
        err  error
    )
    for i := 0; i < 1; i++ {
        t1 := time.Now()
        surf, err = man.Render(cmp)
        fmt.Printf("lap %d: %v err: %v\n", i+1, time.Since(t1).Seconds(), err)
    }
    filename := fmt.Sprintf("xmg-%s_%s", time.Now().Format("02012006T150405"), cmp.MainImage.GfxStr)

    man.Save(surf, filename)
}

func main() {
    t1 := time.Now()
    man := xmg.NewManager("sample_media/images", "sample_media/fonts", "out")
    Test(man)
    fmt.Printf("%v\n", time.Since(t1).Seconds())
}
