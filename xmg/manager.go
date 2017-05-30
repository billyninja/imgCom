package xmg

import (
    "errors"
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "io/ioutil"
    "log"
    "path/filepath"
)

type Manager struct {
    ImageDirs []*ImgDir
    FontDirs  []*FontDir
}

type ImgDir struct {
    path   string
    images map[string]*ImageResource
}

type FontDir struct {
    path  string
    fonts map[string]*FontResource
}

type ImageResource struct {
    Filename string
    Surface  *sdl.Surface
}

type FontResource struct {
    Filename string
    Font     *ttf.Font
}

func NewManager(imgDirs []string, fontDirs []string) (*Manager, error) {
    var img_dirs []*ImgDir
    var font_dirs []*FontDir

    for _, path := range imgDirs {
        dir, err := NewImgDir(path)
        if err == nil {
            img_dirs = append(img_dirs, dir)
        } else {
            log.Panicf("Warning: %s not loaded!", path)
        }
    }

    for _, path := range fontDirs {
        dir, err := NewFontDir(path)
        if err == nil {
            font_dirs = append(font_dirs, dir)
        } else {
            log.Panicf("Warning: %s not loaded!", path)
        }
    }

    return &Manager{
        ImageDirs: img_dirs,
        FontDirs:  font_dirs,
    }, nil
}

func NewImgDir(path string) (*ImgDir, error) {
    log.Printf("Initializing new image dir at path: %s", path)
    images := make(map[string]*ImageResource)

    exists, err := exists(path)
    if !exists || err != nil {
        return nil, err
    }

    fileinfo, err := ioutil.ReadDir(path)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    dir := &ImgDir{
        path:   path,
        images: images,
    }

    for _, fl := range fileinfo {
        ext := filepath.Ext(fl.Name())
        if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
            dir.Add(fl.Name())
        }
    }

    return dir, err
}

func (id *ImgDir) List() {
    for _, image := range id.images {
        println(image.Filename)
    }
}

func (id *ImgDir) LoadAndGet(filename string) (*ImageResource, error) {

    if ir, ok := id.images[filename]; ok {
        path := filepath.Join(id.path, filename)
        image, err := img.Load(path)
        ir.Surface = image
        return ir, err
    } else {
        return nil, errors.New("Entry dont exist")
    }
}

func (id *ImgDir) Add(filename string) error {
    fullpath := filepath.Join(id.path, filename)
    println(fullpath)
    if ok, err := exists(fullpath); ok == false {
        return err
    }

    id.images[filename] = &ImageResource{
        Filename: filename,
    }

    log.Printf("Adding new image file entry at: %s", fullpath)

    return nil
}

func NewFontDir(path string) (*FontDir, error) {
    log.Printf("Initializing new font dir at path: %s", path)
    fonts := make(map[string]*FontResource)

    exists, err := exists(path)
    if !exists || err != nil {
        return nil, err
    }

    fileinfo, err := ioutil.ReadDir(path)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    for _, fl := range fileinfo {
        log.Println(">", fl.Name())
    }

    return &FontDir{
        path:  path,
        fonts: fonts,
    }, err
}

func (fd *FontDir) List() {
    println(">", fd.path)
    for _, font := range fd.fonts {
        println(">>>", font.Filename)
    }
}

func (fd *FontDir) LoadAndGet(filename string, size int) (*FontResource, error) {

    if fr, ok := fd.fonts[filename]; ok {
        path := filepath.Join(fd.path, filename)
        font, err := ttf.OpenFont(path, size)
        fr.Font = font

        return fr, err
    } else {
        return nil, errors.New("Font entry dont exist")
    }
}

func (fd *FontDir) Add(filename string) error {
    fullpath := filepath.Join(fd.path, filename)
    if ok, err := exists(fullpath); ok == false {
        return err
    }

    if _, ok := fd.fonts[filename]; ok {
        return errors.New("Font Already exists")
    }

    fd.fonts[filename] = &FontResource{
        Filename: filename,
    }

    return nil
}

func (man *Manager) GetSurface(str string) (*sdl.Surface, error) {
    println("Entered get Surface")

    for _, iDir := range man.ImageDirs {
        println("searching at ", iDir.path)
        r, err := iDir.LoadAndGet(str)

        if r != nil && r.Surface != nil {
            return r.Surface, err
        }
    }

    msg := fmt.Sprintf("Couldnt find Surface %s, looked %d dirs.", str, len(man.ImageDirs))
    log.Print(msg)
    return nil, errors.New(msg)
}

func (man *Manager) GetFont(str string) (*ttf.Font, error) {
    println("Entered get Font")

    for _, fDir := range man.FontDirs {
        println("searching at ", fDir.path)
        r, err := fDir.LoadAndGet(str)

        if r != nil && r.Surface != nil {
            return r.Surface, err
        }
    }

    msg := fmt.Sprintf("Couldnt find Font %s, looked %d dirs.", str, len(man.ImageDirs))
    log.Print(msg)
    return nil, errors.New(msg)
}
