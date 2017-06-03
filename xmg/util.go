package xmg

import (
    "os"
    "path/filepath"
)

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, err
    }
    return true, err
}

func IsImgFile(path string) bool {
    switch filepath.Ext(path) {
    case ".jpg", ".png", ".gif", ".bmp", ".jpeg":
        return true
    }
    return false
}

func IsFontFile(path string) bool {
    return false
}
