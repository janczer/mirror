# mirror

It's a very simple package with only one function for mirror image.

Function `MirrorImage` get 3 parameters:

```
func MirrorImage(in image.Image, ox bool, oy bool) image.Image
```

- `in` - object for mirror
- `ox` - if `true` when image mirrored horizontally
- `oy` - if `true` when image mirrored vertically

## Simple usage

First get the package:

```
$ go get github.com/janczer/mirror
```

After that create `main.go` and add this:

```
package main

import (
    "fmt"
    "github.com/janczer/mirror"
    "image"
    "image/jpeg"
    "log"
    "os"
)

func main() {
    fmt.Println("simple usage github.com/janczer/mirror")

    reader, err := os.Open("test.jpg") //open image file
    if err != nil {
        log.Fatal(err)
    }
    defer reader.Close()

    m, _, err := image.Decode(reader) //decode file
    if err != nil {
        log.Fatal(err)
    }

    mirroredImage := mirror.MirrorImage(m, true, true)

    //create new image file
    f, err := os.Create("mirror.jpg")
    jpeg.Encode(f, mirroredImage, nil)
    f.Close()
}

```

Save image that you want mirrored to the same directory that `main.go` and run it:

```
$ go run main.go
```

And you have two files `test.jpg` and `mirror.jpg`.

![Gopher](test.jpg)
![Gopher](mirror.jpg)


