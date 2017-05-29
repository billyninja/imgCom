package xmg

import(
    "log"
    "errors"
    "io/ioutil"
    "path/filepath"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_ttf"
)


type Manager struct {
    ImageDirs []*ImgDir
    FontDirs []*FontDir
}

type ImgDir struct {
    path string
    images map[string]*ImageResource
}

type FontDir struct {
    path string
    fonts map[string]*FontResource
}

type ImageResource struct {
    Filename string
    Handler *sdl.Surface
}

type FontResource struct {
    Filename string
    Handler *ttf.Font
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
        FontDirs: font_dirs,        
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
        path: path,
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
        ir.Handler = image
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
        path: path,
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
        fr.Handler = font

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
