package resizeproxy

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "fmt"
)

var (
    formats = []string{
        "png",
        "jpg",
        "gif",
    }
)

func TestImageFormats(t *testing.T) {
    s := httptest.NewServer( http.FileServer(http.Dir("images")))
    defer s.Close()

    for _,format := range formats {
        url := s.URL + "/gopher." + format

        fmt.Println("Downloading from url: %s", url)
        img, err := Download(url)
        if err != nil {
            t.Fatalf("Error on Download image=" + url, err)
        }

        i := *img
        if i.Bounds().Dx() != 250 {
            t.Fatalf("Wrong image width!", i.Bounds().Dx())
        }
        if i.Bounds().Dy() != 340 {
            t.Fatalf("Wrong image height!", i.Bounds().Dy())
        }
    }

}

func TestNotFound(t *testing.T){
    img, err := Download("http://not-existing.com/img.png")
    if err == nil {
        t.Fatalf("Should have returned an error!")
    }
    if img != nil {
        t.Fatalf("Should have returned a nil obj!")
    }
}

func TestBadUrl(t *testing.T){
    img, err := Download("asdf asfd")
    if err == nil {
        t.Fatalf("Should have returned an error!")
    }
    if img != nil {
        t.Fatalf("Should have returned a nil obj!")
    }
}
