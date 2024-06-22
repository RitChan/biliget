package gui

import (
	"image"
	"time"

	"github.com/AllenDang/giu"
)

type guiState struct {
	qrstate *qrcodeState
}

type qrcodeState struct {
	texture  *giu.Texture
	image    image.Image
	key      string
	url      string
	duration int       // seconds
	ctime    time.Time // unix timestamp
}

var globalGuiState *guiState

func globalState() *guiState {
	if globalGuiState == nil {
		globalGuiState = new(guiState)
	}
	return globalGuiState
}

func (s *qrcodeState) expired() bool {
	return time.Now().After(s.ctime.Add(time.Duration(s.duration * 1e9)))
}
