package resizeproxy

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
    "io/ioutil"
    "image"
    "fmt"
)

func getImage(t *testing.T, query string) image.Image {
    originServer := httptest.NewServer( http.FileServer(http.Dir("images")))
    defer originServer.Close()

    proxyHandler := NewProxyHandler(originServer.URL)
    proxyServer := httptest.NewServer(proxyHandler)
    defer proxyServer.Close()

    url := proxyServer.URL + "/" + query
    fmt.Println("Serving from url:", url)

    resp, err := http.Get(url)
    if err != nil {
        t.Fatalf("Could not download image from proxy! Error=%s", err)
    }

    if resp.StatusCode != 200 {
        t.Fatalf("Could download image from proxy! Statuscode=%s", resp.StatusCode, resp.Body)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Could not download image from proxy! Statuscode=%s", err)
    }

    reader := bytes.NewReader(body)
    img,_, err := image.Decode(reader)
    if err != nil{
        t.Fatalf("Could not decode image from proxy! Statuscode=%s", err)
    }
    return img
}

func TestResizePngByWidth(t *testing.T){
    i := getImage(t, "gopher.png?w=100")
    if i.Bounds().Dx() != 100 {
        t.Fatalf("Wrong image width!", i.Bounds().Dx())
    }
}

func TestResizePngByHeigh(t *testing.T){
    i := getImage(t, "gopher.png?h=100")
    if i.Bounds().Dy() != 100 {
        t.Fatalf("Wrong image height!", i.Bounds().Dy())
    }
}

func TestResizeWithoutChange(t *testing.T){
    i := getImage(t, "gopher.png")
    if i.Bounds().Dx() != 250 {
        t.Fatalf("Wrong image width!", i.Bounds().Dx())
    }
}
