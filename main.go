package main
import (
  "fmt"
  "os"
  "image"
  "flag"
  "path/filepath"
  _ "image/png"
  _ "image/jpeg"
	color "github.com/gookit/color"
	ico "github.com/mat/besticon/ico"
  color_extractor "github.com/marekm4/color-extractor"
)
func main() {
  var SkipCommon bool
  flag.BoolVar(&SkipCommon, "C", false, "Skip black, white and gray colors")
  flag.Parse()
  pic, where_is_file := os.Open(flag.Arg(0))
  if where_is_file != nil {
    fmt.Println("I did not find file. Run picdoc with filename.")
    os.Exit(1)
  }
  var dec image.Image
	ext := filepath.Ext(flag.Arg(0))
  if ext == ".jpg" || ext == ".png" || ext == ".jpeg"{
    dec,_,_ = image.Decode(pic)
  } else if ext == ".ico" {
    dec,_ = ico.Decode(pic)
	} else {
		fmt.Println("I can't recognize file. Is it picture (PNG, JPG, JPEG, ICO)?")
		os.Exit(1)
	}
  colors := color_extractor.ExtractColors(dec)
  nearest := colors[0]
  R, G, B, _ := nearest.RGBA()
	R >>= 8
	G >>= 8
	B >>= 8
  if R == G && R == B && G == B && SkipCommon {
    nearest = colors[1]
    R, G, B, _ = nearest.RGBA()
		R >>= 8
		G >>= 8
		B >>= 8
  }
	var SR, SG, SB string
SR = fmt.Sprintf("%X", R)
SG = fmt.Sprintf("%X", G)
SB = fmt.Sprintf("%X", B)
	if len(SR) < 2 {
		SR = "0" + SR
	}
	if len(SG) < 2 {
		SG = "0" + SG
	}
	if len(SB) < 2 {
		SB = "0" + SB
	}
 	color.Printf(`<bg=%s%s%s>    </> <fg=ffff00>(%d, %d, %d = #%s%s%s)</> %s`,SR,SG,SB,R,G,B,SR,SG,SB,flag.Arg(0))
	fmt.Println()
}