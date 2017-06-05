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

type Manager struct {
    Images *ImageManager
    Fonts  *FontManager
    Out    string
}

type ImageManager struct {
    BasePath  string
    Resources map[string]*sdl.Surface
}

type FontManager struct {
    BasePath  string
    Fallback  string
    Resources map[string]map[int]*ttf.Font
}

func NewImageManager(image_dir, fallback string) *ImageManager {
    m := &ImageManager{
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

func (m *ImageManager) Load(resource string) (*sdl.Surface, error) {
    resource = filepath.Join(m.BasePath, resource)
    surf, ok := m.Resources[resource]
    if !ok {
        return nil, image_not_found
    }

    if ok && surf != nil {
        return surf, nil
    }

    s, err := img.Load(resource)
    if err != nil {
        delete(m.Resources, resource)
    } else {
        println("stored ", resource)
        m.Resources[resource] = s
    }

    return s, err
}

func (m *ImageManager) List() []string {
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
        Fallback:  fallback,
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
        if len(m.Fallback) > 0 {
            println("Going for the callback!", m.Fallback)
            return m.Load(m.Fallback, size)
        }
        return nil, font_not_found
    }

    if ok {
        font, ok2 := size_map[size]
        if !ok2 {
            println("Font size miss!")
            f, err := ttf.OpenFont(resource, size)
            if err != nil {
                println("ERR deleting entry: ", resource)
                delete(m.Resources, resource)
            } else {
                println("new size entry for: ", resource, size)
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

func NewManager(image_dir, font_dir, output_dir string) *Manager {
    return &Manager{
        Images: NewImageManager(image_dir, ""),
        Fonts:  NewFontManager(font_dir, "Go-Regular.ttf"),
        Out:    output_dir,
    }
}

func (m *Manager) Render(cmp *Composition) (*sdl.Surface, error) {
    return render(cmp, m.Images, m.Fonts)
}

func (m *Manager) RenderAndSave(cmp *Composition, filename string) (*sdl.Surface, error) {
    s, err := render(cmp, m.Images, m.Fonts)
    if err != nil {
        return nil, err
    }
    err = m.Save(s, filename)

    return s, err
}

func (m *Manager) RenderThumb(cmp *Composition, s *sdl.Surface) (*sdl.Surface, error) {
    return renderThumb(cmp, s)
}

func (m *Manager) Save(surf *sdl.Surface, filename string) error {
    filename = filepath.Join(m.Out, filename)
    return img.SavePNG(surf, filename)
}
