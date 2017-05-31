package xmg

import (
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "os"
)

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
