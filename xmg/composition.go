package xmg

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
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
