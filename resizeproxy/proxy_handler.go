package resizeproxy

import (
    "net/http"
    "image/png"
    "image"
    "strconv"
    "log"
)

type ProxyHandler struct {
    OriginUrl string
}

func NewProxyHandler(originUrl string) ProxyHandler {
    return ProxyHandler{ OriginUrl: originUrl}
}

func (handler ProxyHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
    url := handler.OriginUrl +  req.URL.Path
    w := req.URL.Query().Get("w")
    h := req.URL.Query().Get("h")

    if req.URL.RawQuery != "" {
        url = url + "?" + req.URL.Query().Encode()
    }

    log.Println("using url=", url)
    img, err := Download(url)
    if err != nil {
        http.Error(writer, err.Error(), 500)
        log.Println("error processing url=", url, err)
        return
    }

    resized := resizeFromParams(img, w, h)
    serveImage(writer, resized)
}

func resizeFromParams(img *image.Image, w string, h string) image.Image {
    width, height := 0,0
    if w != ""{
        width,_ = strconv.Atoi(w)
    }

    if h != ""{
        height,_ = strconv.Atoi(h)
    }
    return Resize(*img, width, height)
}

func serveImage(writer http.ResponseWriter, img image.Image) {
    png.Encode(writer, img)
}
