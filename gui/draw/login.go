package draw

import (
	"biliget/gui/goroutines"
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"

	"github.com/AllenDang/giu"
	"github.com/skip2/go-qrcode"
)

var qrkey string
var qrtexture *giu.Texture

func DrawLogin() {
	radius := float32(12)
	center := image.Pt(773, 15)
	canvas := giu.GetCanvas()
	hovered := mouseInCirclePt(center, radius)

	// Profile
	var profileColor color.Color
	if hovered {
		profileColor = color.RGBA{255, 0, 0, 255}
	} else {
		profileColor = color.RGBA{255, 255, 255, 255}
	}
	canvas.AddCircle(center, radius, profileColor, 32, 2)

	// QR Code
	s, err := goroutines.GetQrcodeState()
	if err != nil {
		giu.Msgbox("Qrcode Error", err.Error()).Buttons(giu.MsgboxButtonsOk)
	} else {
		if s.LoginExpired {
			giu.Msgbox("登录过期", "登录过期, 请重新登录").Buttons(giu.MsgboxButtonsOk)
		}
		if s.LoginState != goroutines.Succeeded && hovered {
			pMin := image.Point{537, 53}
			pMax := image.Point{787, 303}
			qrColor := color.RGBA{255, 255, 255, 255}
			canvas.AddRectFilled(pMin, pMax, qrColor, 12, giu.DrawFlagsRoundCornersAll)
			canvas.AddText(image.Pt(633, 271), color.RGBA{0, 0, 0, 255}, s.Message)
			if s.QrKey != qrkey {
				pngbytes, err := qrcode.Encode(s.QrUrl, qrcode.High, 180)
				if err != nil {
					log.Println(err.Error())
					qrkey = ""
					qrtexture = nil
				} else {
					img, err := png.Decode(bytes.NewReader(pngbytes))
					if err != nil {
						log.Println(err.Error())
						qrkey = ""
						qrtexture = nil
					}
					giu.EnqueueNewTextureFromRgba(img, func(t *giu.Texture) { qrtexture = t })
				}
			}
			if qrtexture != nil {
				pMin := image.Pt(571, 73)
				pMax := pMin.Add(image.Pt(180, 180))
				canvas.AddImage(qrtexture, pMin, pMax)
			}
		} else if s.LoginState == goroutines.Succeeded && hovered {
			pMin := image.Point{537, 53}
			pMax := image.Point{787, 303}
			qrColor := color.RGBA{255, 255, 255, 255}
			canvas.AddRectFilled(pMin, pMax, qrColor, 12, giu.DrawFlagsRoundCornersAll)
			canvas.AddText(image.Pt(633, 271), color.RGBA{0, 0, 0, 255}, "登录成功")
		}
	}

}
