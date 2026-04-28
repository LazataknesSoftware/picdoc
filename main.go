package main
import (
  "fmt"
  "os"
  "image"
  "flag"
  "path/filepath"
  "encoding/json"
  _ "image/png"
  _ "image/jpeg"
	color "github.com/gookit/color"
	ico "github.com/mat/besticon/ico"
  color_extractor "github.com/marekm4/color-extractor"
  locallib "github.com/jeandeaual/go-locale"
)
type ErrorDescription struct {
  Error_filenotfound string
  Error_unrecognizablefiletype string
  Help_blackgraywhite string
}
func main() {
  Lang := ErrorDescription{}
  UserLocale, _ := locallib.GetLanguage()
  Localefile, not_ru_or_en := os.ReadFile(fmt.Sprintf("picdoc_locale_%s.json",UserLocale));
  if not_ru_or_en != nil {
    UserLocale = "en"
  }
  _ = json.Unmarshal(Localefile, &Lang)

  var SkipCommon bool
  flag.BoolVar(&SkipCommon, "C", false, Lang.Help_blackgraywhite)
  flag.Parse()
  pic, where_is_file := os.Open(flag.Arg(0))
  if where_is_file != nil {
    fmt.Println(Lang.Error_filenotfound)
    os.Exit(1)
  }
  var dec image.Image
	ext := filepath.Ext(flag.Arg(0))
  if ext == ".jpg" || ext == ".png" || ext == ".jpeg"{
    dec,_,_ = image.Decode(pic)
  } else if ext == ".ico" {
    dec,_ = ico.Decode(pic)
	} else {
		fmt.Println(Lang.Error_unrecognizablefiletype)
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