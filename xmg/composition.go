package xmg

type Pos struct {
    X          int32
    Y          int32
    RelativeTo uint8
}

type Scale struct {
    W int32
    H int32
}

type Color struct {
    R uint8
    G uint8
    B uint8
    A uint8
}

type TextEl struct {
    FontStr  string
    FontSize int
    Message  string
    Color    *Color
    Align    uint8
    Pos      *Pos
}

type GfxEl struct {
    GfxStr string
    Pos    *Pos
    Scale  *Scale
}

type Composition struct {
    MainImage  *GfxEl
    BGColor    *Color
    Dimensions *Scale
    Gfx        []*GfxEl
    Text       []*TextEl
    Loaded     bool
}
