package xmg

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "log"
    "os"
)

func RenderGFX(r *sdl.Renderer, g *GfxEl) {
    var w, h, x, y int32

    gs, err := img.Load(g.GfxStr)
    if err != nil {
        os.Exit(2)
    }
    gt, _ := r.CreateTextureFromSurface(gs)

    if g.Scale != nil {
        w = g.Scale.W
        h = g.Scale.H
    } else {
        w = gs.W
        h = gs.H
    }

    if g.Pos != nil {
        x = g.Pos.X
        y = g.Pos.Y
    }

    r.Copy(
        gt,
        &sdl.Rect{0, 0, gs.W, gs.H},
        &sdl.Rect{x, y, w, h},
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
        &sdl.Rect{t.Pos.X, t.Pos.Y, ts.W, ts.H},
    )
}

func Render(cmp *Composition) *sdl.Renderer {
    var (
        r *sdl.Renderer
        s *sdl.Surface
    )

    rct := &sdl.Rect{0, 0, cmp.Dimensions.W, cmp.Dimensions.H}
    s, _ = sdl.CreateRGBSurface(0, rct.W, rct.H, 24, 0, 0, 0, 0)
    r, _ = sdl.CreateSoftwareRenderer(s)
    if cmp.BGColor != nil {
        _ = r.SetDrawColor(cmp.BGColor.R, cmp.BGColor.G, cmp.BGColor.B, cmp.BGColor.A)
        r.FillRect(rct)
    }
    if cmp.MainImage != nil {
        if cmp.MainImage.Scale == nil {
            cmp.MainImage.Scale = &Scale{rct.W, rct.H}
        }
        RenderGFX(r, cmp.MainImage)
    }

    for _, g := range cmp.Gfx {
        RenderGFX(r, g)
    }

    ttf.Init()
    for _, t := range cmp.Text {
        RenderText(r, t)
    }

    img.SavePNG(s, "out/strix-nebulosa+trophy1600.png")

    return r
}
