package draw

import (
	"biliget/gui/data"
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
	s, err := data.GetQrcodeState()
	if err != nil {
		giu.Msgbox("Qrcode Error", err.Error()).Buttons(giu.MsgboxButtonsOk)
	} else {
		if s.LoginExpired {
			giu.Msgbox("登录过期", "登录过期, 请重新登录").Buttons(giu.MsgboxButtonsOk)
		}
		if hovered {
			rectMin := image.Point{537, 53}
			rectMax := image.Point{787, 303}
			rectColor := color.RGBA{255, 255, 255, 255}
			rectRound := float32(12)
			textPos := image.Pt(633, 271)
			textCol := color.RGBA{0, 0, 0, 255}
			canvas.AddRectFilled(rectMin, rectMax, rectColor, rectRound, giu.DrawFlagsRoundCornersAll)
			canvas.AddText(textPos, textCol, s.Message)
			if s.LoginState != data.Succeeded {
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
						qrkey = s.QrKey
						giu.EnqueueNewTextureFromRgba(img, func(t *giu.Texture) { qrtexture = t })
					}
				}
				if qrtexture != nil {
					pMin := image.Pt(571, 73)
					pMax := pMin.Add(image.Pt(180, 180))
					canvas.AddImage(qrtexture, pMin, pMax)
				}
			}
		}
	}

}
