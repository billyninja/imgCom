package xmg

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "os"
)

func RenderGFX(r *sdl.Renderer, g *GfxEl, sman *SurfaceManager) {
    var (
        w, h, x, y   int32
        should_scale bool
        s            *sdl.Surface
    )

    gs, err := sman.Load(g.GfxStr)
    if err != nil {
        os.Exit(2)
    }
    if g.Scale != nil {
        should_scale = true
        w = g.Scale.W
        h = g.Scale.H
    } else {
        should_scale = false
        w = gs.W
        h = gs.H
    }

    if should_scale {
        s, _ = sdl.CreateRGBSurface(0, w, h, 32, 0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
        defer s.Free()
        gs.BlitScaled(nil, s, &sdl.Rect{0, 0, w, h})
    } else {
        s = gs
    }
    gt, _ := r.CreateTextureFromSurface(s)
    defer gt.Destroy()

    if g.Pos != nil {
        x = g.Pos.X
        y = g.Pos.Y
    }

    r.Copy(
        gt,
        &sdl.Rect{0, 0, w, h},
        &sdl.Rect{x, y, w, h},
    )
}

func (t *TextEl) Bake(r *sdl.Renderer, f *ttf.Font) *sdl.Surface {
    color := sdl.Color{}
    if t.Color != nil {
        color = sdl.Color{t.Color.R, t.Color.G, t.Color.B, t.Color.A}
    }
    ts, _ := f.RenderUTF8_Blended(t.Message, color)

    return ts
}

func RenderText(r *sdl.Renderer, t *TextEl, fman *FontManager) {
    f, err := fman.Load(t.FontStr, t.FontSize)
    if err != nil {
        os.Exit(2)
    }
    ts := t.Bake(r, f)
    tt, _ := r.CreateTextureFromSurface(ts)
    r.Copy(
        tt,
        &sdl.Rect{0, 0, ts.W, ts.H},
        &sdl.Rect{t.Pos.X, t.Pos.Y, ts.W, ts.H},
    )
}

func Render(cmp *Composition, sman *SurfaceManager, fman *FontManager) *sdl.Surface {
    var (
        r *sdl.Renderer
        s *sdl.Surface
    )

    rct := &sdl.Rect{0, 0, cmp.Dimensions.W, cmp.Dimensions.H}
    s, _ = sdl.CreateRGBSurface(0, rct.W, rct.H, 24, 0, 0, 0, 0)
    r, _ = sdl.CreateSoftwareRenderer(s)
    if cmp.BGColor != nil {
        println("bgcolor", cmp.BGColor.R, cmp.BGColor.G, cmp.BGColor.B, cmp.BGColor.A)
        _ = r.SetDrawColor(cmp.BGColor.R, cmp.BGColor.G, cmp.BGColor.B, cmp.BGColor.A)
        r.FillRect(rct)
    }
    if cmp.MainImage != nil {
        if cmp.MainImage.Scale == nil {
            cmp.MainImage.Scale = &Scale{rct.W, rct.H}
        }
        RenderGFX(r, cmp.MainImage, sman)
    }

    for _, g := range cmp.Gfx {
        RenderGFX(r, g, sman)
    }

    for _, t := range cmp.Text {
        RenderText(r, t, fman)
    }

    return s
}

func Save(filename string, s *sdl.Surface) error {
    return img.SavePNG(s, filename)
}
