package xmg

import (
    "errors"
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "regexp"
    "strconv"
)

func renderGFX(r *sdl.Renderer, g *GfxEl, sman *ImageManager) error {
    var (
        w, h, x, y   int32
        should_scale bool
        s            *sdl.Surface
    )

    gs, err := sman.Load(g.GfxStr)
    if err != nil {
        return err
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
        s, err = createBlankSurface(w, h)
        defer s.Free()
        gs.BlitScaled(nil, s, &sdl.Rect{0, 0, w, h})
    } else {
        s = gs
    }
    gt, err := r.CreateTextureFromSurface(s)
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

    return err
}

func (t *TextEl) bake(r *sdl.Renderer, f *ttf.Font) (*sdl.Surface, error) {
    color := sdl.Color{}
    if t.Color != nil {
        color = sdl.Color{t.Color.R, t.Color.G, t.Color.B, t.Color.A}
    }
    ts, err := f.RenderUTF8_Blended(t.Message, color)

    return ts, err
}

func renderText(r *sdl.Renderer, t *TextEl, fman *FontManager) error {
    f, err := fman.Load(t.FontStr, t.FontSize)
    if err != nil {
        return err
    }

    ts, err := t.bake(r, f)
    tt, err := r.CreateTextureFromSurface(ts)
    if err != nil {
        return err
    }
    r.Copy(
        tt,
        &sdl.Rect{0, 0, ts.W, ts.H},
        &sdl.Rect{t.Pos.X, t.Pos.Y, ts.W, ts.H},
    )

    return err
}

func render(cmp *Composition, sman *ImageManager, fman *FontManager) (*sdl.Surface, error) {
    var (
        r *sdl.Renderer
        s *sdl.Surface
    )

    rct := &sdl.Rect{0, 0, cmp.Dimensions.W, cmp.Dimensions.H}
    s, err := createBlankSurface(rct.W, rct.H)
    r, err = sdl.CreateSoftwareRenderer(s)
    if cmp.BGColor != nil {
        _ = r.SetDrawColor(cmp.BGColor.R, cmp.BGColor.G, cmp.BGColor.B, cmp.BGColor.A)
        r.FillRect(rct)
    }
    if cmp.MainImage != nil {
        if cmp.MainImage.Scale == nil {
            cmp.MainImage.Scale = &Scale{rct.W, rct.H}
        }
        err = renderGFX(r, cmp.MainImage, sman)
    }

    for _, g := range cmp.Gfx {
        err = renderGFX(r, g, sman)
    }

    for _, t := range cmp.Text {
        err = renderText(r, t, fman)
    }

    return s, err
}

func createBlankSurface(w, h int32) (*sdl.Surface, error) {
    return sdl.CreateRGBSurface(0, w, h, 32, 0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
}

func renderThumb(cmp *Composition, s *sdl.Surface) (*sdl.Surface, error) {
    var w, h int32
    spc := cmp.Thumbfy.Spec

    var validSpec = regexp.MustCompile(`([0-9]+)?x([0-9]+)?`)
    res := validSpec.FindAllStringSubmatch(spc, 2)
    if len(res) == 0 {
        return nil, errors.New(fmt.Sprintf("Invalid thumb format: %s", spc))
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
    }

    if w == 0 {
        w = int32(findAR(s.W, s.H) * float32(h))
    }

    st, err := createBlankSurface(w, h)
    err = s.BlitScaled(nil, st, &sdl.Rect{0, 0, w, h})

    return st, err
}
