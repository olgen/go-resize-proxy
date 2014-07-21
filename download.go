package resizeproxy

import (
    "image"
    _ "image/png"
    _ "image/jpeg"
    _ "image/gif"
    "net/http"
    "bytes"
    "io/ioutil"
    "errors"
    "fmt"
)

func Download(url string) ( *image.Image, error ) {
    img, err := loadHttp(url)
    if err != nil {
        return nil, err
    }
    return img, nil
}

func loadHttp(url string) ( *image.Image, error ) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != 200 {
        return nil, errors.New(fmt.Sprintf("Could not download image! Statuscode=%s", resp.StatusCode))
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    reader := bytes.NewReader(body)
    img,_, err := image.Decode(reader)
    if err != nil{
        return nil, err
    }
    return &img, nil
}
