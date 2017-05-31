package xmg

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "log"
)

type Pos struct {
    X          int
    Y          int
    RelativeTo uint8
}

type TextEl struct {
    FontStr  string
    FontSize int
    font     *ttf.Font
    Message  string
    Align    uint8
    Pos      *Pos
}

type GfxEl struct {
    GfxStr string
    gfx    *sdl.Surface
    Pos    *Pos
}

type Composition struct {
    ImageStr string
    img      *sdl.Surface
    Gfx      []*GfxEl
    Text     []*TextEl
    Loaded   bool
}

func (cmp *Composition) LoadResources(man *Manager) {
    println("LoadResources")
    surf, err := man.GetSurface(cmp.ImageStr)
    if err != nil {
        log.Printf("Coundnt load %s: \n\n %v", cmp.ImageStr, err)
        cmp.Loaded = false
        return
    }

    cmp.img = surf

    for _, gfx := range cmp.Gfx {
        println("Load gfx", gfx.GfxStr)
        surf, err := man.GetSurface(gfx.GfxStr)
        if err != nil {
            log.Printf("%v", err)
            cmp.Loaded = false

            return
        }
        gfx.gfx = surf
    }

    for _, txt := range cmp.Text {
        println("Load font", txt.FontStr)
        font, err := man.GetFont(txt.FontStr, txt.FontSize)
        if err != nil {
            log.Printf("%v", err)
            cmp.Loaded = false

            return
        }
        txt.font = font
    }

    cmp.Loaded = true
}

func (cmp *Composition) GetMainImage() *sdl.Surface {
    if cmp.Loaded {
        return nil
    }

    return cmp.img
}
