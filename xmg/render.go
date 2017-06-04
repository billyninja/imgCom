package xmg

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "os"
    "regexp"
    "strconv"
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
        s, _ = createBlankSurface(w, h)
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
    s, _ = createBlankSurface(rct.W, rct.H)
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

    RenderThumb(cmp, s)

    return s
}

func createBlankSurface(w, h int32) (*sdl.Surface, error) {
    return sdl.CreateRGBSurface(0, w, h, 32, 0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
}

func RenderThumb(cmp *Composition, s *sdl.Surface) *sdl.Surface {
    var w, h int32
    spc := cmp.Thumbfy.Spec

    // Todo, send to a global init
    var validSpec = regexp.MustCompile(`([0-9]+)?x([0-9]+)?`)

    res := validSpec.FindAllStringSubmatch(spc, 2)
    if len(res) == 0 {
        println("Invalid thumb format: ", spc)
        os.Exit(2)
    }

    w64, err := strconv.ParseInt(string(res[0][1]), 10, 32)
    if err == nil {
        w = int32(w64)
    }

    h64, err := strconv.ParseInt(string(res[0][2]), 10, 32)
    if err == nil {
        h = int32(h64)
    }

    findAR := func(sw, sh int32) float32 {
        if s.W >= s.H {
            return float32(s.H) / float32(s.W)
        } else {
            return float32(s.W) / float32(s.H)
        }

        return 1
    }

    if h == 0 {
        h = int32(findAR(s.W, s.H) * float32(w))
        println(h)
    }

    if w == 0 {
        w = int32(findAR(s.W, s.H) * float32(h))
    }

    println(w, h)

    st, _ := createBlankSurface(w, h)
    s.BlitScaled(nil, st, &sdl.Rect{0, 0, w, h})

    Save("out/thumb.png", st)
    return st
}

func Save(filename string, s *sdl.Surface) error {
    return img.SavePNG(s, filename)
}
