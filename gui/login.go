package gui

import (
	"biliget/biliinfo/bihttp"
	"bytes"
	"image"
	"image/color"
	"image/png"
	"time"

	g "github.com/AllenDang/giu"
	qrcode "github.com/skip2/go-qrcode"
)

func requestQrcodeState() (*qrcodeState, error) {
	s := new(qrcodeState)
	s.ctime = time.Now()
	resp, err := bihttp.BiliGetQrcodeUrl()
	if err != nil {
		return nil, err
	}
	s.duration = 180
	s.key = resp.Data.QrcodeKey
	s.url = resp.Data.Url
	pngbytes, err := qrcode.Encode(s.url, qrcode.High, 180)
	if err != nil {
		return nil, err
	}
	image, err := png.Decode(bytes.NewReader(pngbytes))
	if err != nil {
		return nil, err
	}
	s.image = image
	return s, nil
}

func getQrcodeState() (*qrcodeState, error) {
	global := globalState()
	if global.qrstate == nil || global.qrstate.expired() {
		s, err := requestQrcodeState()
		if err != nil {
			return nil, err
		}
		global.qrstate = s
	}
	return global.qrstate, nil
}

func drawLogin() {
	radius := float32(12)
	center := image.Pt(773, 15)
	canvas := g.GetCanvas()
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
	if hovered {
		pMin := image.Point{537, 53}
		pMax := image.Point{787, 303}
		qrColor := color.RGBA{255, 255, 255, 255}
		canvas.AddRectFilled(pMin, pMax, qrColor, 12, g.DrawFlagsRoundCornersAll)
		canvas.AddText(image.Pt(633, 271), color.RGBA{0, 0, 0, 255}, "扫码登录")
		s, err := getQrcodeState()
		if err != nil {
			g.Msgbox("Qrcode Error", err.Error())
		} else {
			g.EnqueueNewTextureFromRgba(s.image, func(t *g.Texture) { s.texture = t })
			if s.texture != nil {
				texMin := image.Pt(571, 73)
				texMax := texMin.Add(image.Pt(180, 180))
				canvas.AddImage(s.texture, texMin, texMax)
			}
		}
	}
}
