package resizeproxy

import (
    "image"
    "github.com/nfnt/resize"
)

func Resize(img image.Image, width int, height int) image.Image{
    if width > 0 && height > 0 {
        // preserve the aspect-ration
        return resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)
    } else if ( width > 0 || height > 0 ) && (width < img.Bounds().Dx() && height < img.Bounds().Dy()) {
        // resize to fit either width or height, but only if the new width & height are smaller
        return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
    } else {
        // return the original image otherwise
        return img
    }
}

