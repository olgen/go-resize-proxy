package resizeproxy

import (
    "testing"
    "os"
    "log"

    "image"
    _ "image/png"
    _ "image/jpeg"
    _ "image/gif"
)

func TestResize(t *testing.T) {
    for _,format := range formats {
        file, err := os.Open("images/gopher."+format)
        if err != nil {
            log.Fatal(err)
        }

        img, _, err := image.Decode(file)
        if err != nil {
            log.Fatal(err)
        }
        file.Close()

        // resize to width 1000 using Lanczos resampling
        // and preserve aspect ratio
        resized := Resize(img, 100, 100)
        i := resized

        if i.Bounds().Dx() > 100 {
            t.Fatalf("Wrong image width!", i.Bounds().Dx())
        }
        if i.Bounds().Dy() > 100 {
            t.Fatalf("Wrong image height!", i.Bounds().Dy())
        }
    }
}
