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
    FontStr string
    font    *ttf.Font
    Message string
    Align   uint8
    Pos     *Pos
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
}

func (cmp *Composition) LoadResources(man *Manager) {
    println("LoadResources")
    surf, err := man.GetSurface(cmp.ImageStr)
    if err != nil {
        log.Printf("Coundnt load %s: \n\n %v", cmp.ImageStr, err)
    }

    cmp.img = surf

    for _, gfx := range cmp.Gfx {
        println("Load gfx", gfx.GfxStr)
        surf, err := man.GetSurface(gfx.GfxStr)
        if err != nil {
            log.Printf("%v", err)
        }
        gfx.gfx = surf
    }

    for _, txt := range cmp.Text {
        println("Load font", txt.FontStr)
        font, err := man.GetFont(txt.FontStr)
        if err != nil {
            log.Printf("%v", err)
        }
        txt.font = font
    }
}
