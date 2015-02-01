/* 
* @Author: dreamzml
* @Date:   2015-02-01 05:44:56
* @Last Modified by:   zoro
* @Last Modified time: 2015-02-01 06:16:02
*/

package lib

import (
	//"strconv"
	"image"
    "image/color"
    "image/draw"
    "image/png"
    "bytes"
    //"math/rand"
    "encoding/base64"

	"flag"

	"io/ioutil"
	"log"

	//"fmt"
	//"reflect"
	 "math/rand"

	"github.com/scpayson/freetype-go/freetype"
)


func TexToImg(text string) string {
	
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
    //blue := color.RGBA{0, 255, 0, 255}

    n := rand.Intn(len(clolrs))
	blue := clolrs[n]

    draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)


	flag.Parse()
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return ""
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return ""
	}


	fg := image.Black

	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(font)
	c.SetFontSize(*size)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(fg)


	pt := freetype.Pt(10, 0+c.FUnitToPixelRU(font.UnitsPerEm()))
	//转化为数组处理
	texts := []string{text}
	for _, s := range texts {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return ""
		}
		pt.Y += c.PointToFix32(*size * *spacing)
	}


   	var buf bytes.Buffer
    png.Encode(&buf, m)
    return base64.StdEncoding.EncodeToString(buf.Bytes())
}


var (
	dpi      = flag.Int("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./static/fonts/STXINGKA.TTF", "filename of the ttf font")
	size     = flag.Float64("size", 75, "font size in points")
	spacing  = flag.Float64("spacing", 1.0, "line spacing (e.g. 2 means double spaced)")
)

var clolrs = []color.RGBA{
 	//color.RGBA{88,  87,  86,  255},		//象牙黑
 	color.RGBA{128, 138, 135, 255},		//冷灰
 	//color.RGBA{128, 118, 105, 255},		//暖灰
 	//color.RGBA{118, 128, 105, 255},		//石板灰
 	color.RGBA{252, 230, 202, 255},		//蛋壳灰

 	//color.RGBA{227, 23,  13,  255},		//镉红
 	//color.RGBA{156, 102, 31,  255},		//砖红
 	color.RGBA{255, 127, 80,  255},		//珊瑚红
 	color.RGBA{255, 99,  71,  255},		//番茄红
 	color.RGBA{255, 192, 203, 255},		//粉红

 	//color.RGBA{176, 23,  31,  255},		//印度红
 	//color.RGBA{255, 0,   255, 255},		//深红
 	//color.RGBA{116, 0,   0,   255},		//黑红
 	color.RGBA{0,   255, 255, 255},		//青色
 	//color.RGBA{8,   46,  84,  255},		//靛青色

 	color.RGBA{127, 255, 0,   255},		//黄绿色
 	color.RGBA{64,  224, 205, 255},		//青绿色
 	color.RGBA{3,   168, 158, 255},		//锰蓝
 	color.RGBA{0,   199, 140, 255},		//土耳其蓝
 	color.RGBA{255, 153, 18,  255},		//镉黄

 	color.RGBA{227, 207, 87,  255},		//香蕉黄
 	color.RGBA{255, 128, 0,   255},		//橘黄
 	color.RGBA{237, 145, 33,  255},		//萝卜黄
 	color.RGBA{85,  102, 0,   255},		//黑黄
 	color.RGBA{199, 97,  20,  255},		//土色

 	color.RGBA{244, 164, 95,  255},		//沙棕色
 	color.RGBA{210, 180, 140, 255},		//棕褐色
 	color.RGBA{188, 143, 143, 255},		//玫瑰红
 	color.RGBA{160, 82,  45,  255},		//赫色
 	color.RGBA{199, 97,  20,  255},		//肖贡土色

 	color.RGBA{160,32,240, 255},		//肖贡土色
 	color.RGBA{218,112,214, 255},		//淡紫色
 	color.RGBA{138,43,226, 255},		//紫罗兰
 	color.RGBA{153,51,250, 255},		//胡紫色
}