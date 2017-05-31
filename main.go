package main

import (
    "github.com/billyninja/imgCom/xmg"
)

func Test(man *xmg.Manager) {
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

    cmp.LoadResources(man)
    xmg.Render(cmp)
}

func main() {
    man, _ := xmg.NewManager(
        []string{"sample_media/images"},
        []string{"sample_media/fonts"},
    )

    Test(man)
}
