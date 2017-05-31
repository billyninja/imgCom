package xmg

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "log"
    "os"
)

func GetRendererFromSurf(surf *sdl.Surface) (*sdl.Renderer, error) {
    r, err := sdl.CreateSoftwareRenderer(surf)
    if err != nil {
        os.Exit(2)
    }

    return r, err
}

func RenderGFX(r *sdl.Renderer, g *GfxEl) {
    gs, err := img.Load("sample_media/images/trophy1600.png")
    if err != nil {
        os.Exit(2)
    }
    gt, _ := r.CreateTextureFromSurface(gs)

    r.Copy(
        gt,
        &sdl.Rect{0, 0, gs.W, gs.H},
        &sdl.Rect{int32(g.Pos.X), int32(g.Pos.Y), gs.W, gs.H},
    )
}

func (t *TextEl) Bake(r *sdl.Renderer) *sdl.Surface {
    bkl := sdl.Color{0, 0, 0, 0}
    font, err := ttf.OpenFont("sample_media/fonts/Go-Regular.ttf", t.FontSize)
    if err != nil {
        log.Printf("%v", err)
        os.Exit(2)
    }

    ts, _ := font.RenderUTF8_Blended(t.Message, bkl)
    return ts
}

func RenderText(r *sdl.Renderer, t *TextEl) {
    ts := t.Bake(r)
    tt, _ := r.CreateTextureFromSurface(ts)
    r.Copy(
        tt,
        &sdl.Rect{0, 0, ts.W, ts.H},
        &sdl.Rect{int32(t.Pos.X), int32(t.Pos.Y), ts.W, ts.H},
    )
}

func Render(cmp *Composition) *sdl.Renderer {
    bg_s, err := img.Load("sample_media/images/strix-nebulosa.jpg")
    if err != nil {
        os.Exit(2)
    }

    r, _ := sdl.CreateSoftwareRenderer(bg_s)

    for _, g := range cmp.Gfx {
        RenderGFX(r, g)
    }

    ttf.Init()
    for _, t := range cmp.Text {
        RenderText(r, t)
    }

    img.SavePNG(bg_s, "out/strix-nebulosa+trophy1600.png")

    return r
}
