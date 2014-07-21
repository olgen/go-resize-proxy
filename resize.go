package resizeproxy

import (
    "image"
    "github.com/nfnt/resize"
)

func Resize(img image.Image, width int, height int) image.Image{
    return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

