package xmg

import(
    "os"
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
)


func GetRendererFromSurf(surf *sdl.Surface) (*sdl.Renderer, error) {
    rdr, err := sdl.CreateSoftwareRenderer(surf)
    if err != nil {
        os.Exit(2)
    }

    return rdr, err
}

func LoadImg(filename string) (*sdl.Surface, error) {
    image, err := img.Load(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to load image: %s\n", err)
        return nil, err
    }
    
    return image, err
}


func GetRendererFromFilename(filename string) (*sdl.Surface, *sdl.Renderer, error) {
    surfImg, err := LoadImg(filename)
    rdr, err := GetRendererFromSurf(surfImg)

    return surfImg, rdr, err
}