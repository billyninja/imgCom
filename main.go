package main 

import(
    "github.com/billyninja/imgCom/xmg"
    "github.com/veandco/go-sdl2/sdl_image"
)


func main() {
    surf, rdr, err := xmg.GetRendererFromFilename("sample_media/images/strix-nebulosa.jpg")
    err = img.SavePNG(surf, "out/strix-nebulosa.png")
    man, err := xmg.NewManager(
        []string{"sample_media/images"},
        []string{"sample_media/fonts"},
    )

    println(man, rdr, err)
}
