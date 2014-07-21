package resizeproxy

import (
    "testing"
    "os"
    "log"

    "image"
    _ "image/png"
)

func testResize(t *testing.T, width int, height int) image.Image {
    file, err := os.Open("images/gopher.png")
    if err != nil {
        log.Fatal(err)
    }

    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    return Resize(img, width, height)
}

func TestResizeToWidth(t *testing.T){
    resized := testResize(t, 100, 0)

    if resized.Bounds().Dx() > 100 {
        t.Fatalf("Wrong image width!", resized.Bounds().Dx())
    }
}

func TestResizeToHeight(t *testing.T){
    resized := testResize(t, 0, 100)

    if resized.Bounds().Dy() > 100 {
        t.Fatalf("Wrong image height!", resized.Bounds().Dy())
    }
}


func TestUpsize(t *testing.T){
    resized := testResize(t, 500, 0)

    if resized.Bounds().Dx() > 250 {
        t.Fatalf("Wrong image width!", resized.Bounds().Dx())
    }
}

