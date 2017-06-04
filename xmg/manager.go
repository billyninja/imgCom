package xmg

import (
    "errors"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "os"
    "path/filepath"
)

var (
    image_not_found error = errors.New("image not found!")
    font_not_found  error = errors.New("font not found!")
)

type SurfaceManager struct {
    BasePath  string
    Resources map[string]*sdl.Surface
}

type FontManager struct {
    BasePath  string
    Resources map[string]map[int]*ttf.Font
}

func NewSurfaceManager(image_dir, fallback string) *SurfaceManager {
    m := &SurfaceManager{
        BasePath:  image_dir,
        Resources: make(map[string]*sdl.Surface),
    }

    filepath.Walk(image_dir, func(p string, i os.FileInfo, e error) error {
        if IsImgFile(p) {
            m.Resources[p] = nil
        }
        return nil
    })

    return m
}

func (m *SurfaceManager) Load(resource string) (*sdl.Surface, error) {
    resource = filepath.Join(m.BasePath, resource)
    surf, ok := m.Resources[resource]
    if !ok {
        return nil, image_not_found
    }

    if ok && surf != nil {
        println("Hit")
        return surf, nil
    }

    println("miss")
    s, err := img.Load(resource)
    if err != nil {
        delete(m.Resources, resource)
    } else {
        println("stored ", resource)
        m.Resources[resource] = s
    }

    return s, err
}

func (m *SurfaceManager) List() []string {
    out := make([]string, len(m.Resources))
    i := 0
    for k := range m.Resources {
        out[i] = k
        i++
    }

    return out
}

func NewFontManager(font_dir, fallback string) *FontManager {
    ttf.Init()
    m := &FontManager{
        Resources: make(map[string]map[int]*ttf.Font),
        BasePath:  font_dir,
    }

    filepath.Walk(font_dir, func(p string, i os.FileInfo, e error) error {
        if IsFontFile(p) {
            m.Resources[p] = make(map[int]*ttf.Font)
        }
        return nil
    })

    return m
}

func (m *FontManager) Load(resource string, size int) (*ttf.Font, error) {
    resource = filepath.Join(m.BasePath, resource)
    size_map, ok := m.Resources[resource]
    if !ok {
        println("Font resource miss!", resource)
        return nil, font_not_found
    }

    if ok {
        println("Font resource hit!")

        font, ok2 := size_map[size]
        if !ok2 {
            println("Font size miss!")
            f, err := ttf.OpenFont(resource, size)
            if err != nil {
                delete(m.Resources, resource)
            } else {
                size_map[size] = f
            }
            return f, err
        } else {
            println("Font size hit!")
            return font, nil
        }
    }

    return nil, font_not_found
}

func (m *FontManager) List() []string {
    out := make([]string, len(m.Resources))
    i := 0
    for k := range m.Resources {
        out[i] = k
        i++
    }

    return out
}
