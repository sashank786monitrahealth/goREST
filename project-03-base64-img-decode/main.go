
// This example demonstrates decoding a JPEG image and examining its pixels.
package main;

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"
        "os"
        "io"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// "image/gif"
	 "image/png"
	 "image/jpeg"
)



//Given a base64 string of a JPEG, encodes it into an JPEG image test.jpg
func base64toJpg(m image.Image, formatString string) {


    bounds := m.Bounds()
    fmt.Println("base64toJpg", bounds, formatString)
    var vFileName string = "test.jpg";
    /*begin*/



    //Encode from image format to writer
    f, err := os.OpenFile(vFileName, os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil {
        log.Fatal(err)
        return
    }

    err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
    if err != nil {
        log.Fatal(err)
        return
    }
    fmt.Println("Jpg file", vFileName, "created")


    /*end*/



fmt.Printf("vFileName = %v\n",vFileName);

}



//Given a base64 string of a PNG, encodes it into an PNG image test.png
func base64toPng(m image.Image, formatString string) {


    bounds := m.Bounds()
    fmt.Println("base64toPng", bounds, formatString)
    var vFileName string = "test.png";
    /*begin*/



    //Encode from image format to writer
    f, err := os.OpenFile(vFileName, os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil {
        log.Fatal(err)
        return
    }

    err = png.Encode(f, m)
    if err != nil {
        log.Fatal(err)
        return
    }
    fmt.Println("Png file", vFileName, "created")


    /*end*/



fmt.Printf("vFileName = %v\n",vFileName);

}




const data = `
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
`

func main(){


    var m image.Image;
    var formatString string;
    var err error;

    var reader io.Reader = base64.NewDecoder(base64.StdEncoding, strings.NewReader(data));
    m, formatString, err = image.Decode(reader);

    if err != nil {
        log.Fatal(err)
    }


    if formatString == "jpeg"{
       base64toJpg(m, formatString)
       
    } else if formatString == "png"{
       base64toPng(m, formatString)
       
    }

}
