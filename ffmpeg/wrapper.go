package ffmpeg

import (
	"errors"
	"os/exec"
)

var ffmpegPath = ""

func HasFFmpeg() bool {
	_, err := GetFFmpeg()
	return err == nil
}

func GetFFmpeg() (string, error) {
	if ffmpegPath == "" {
		return "", errors.New("cannot find ffmpeg")
	}
	return ffmpegPath, nil
}

func InitFFmpeg() error {
	p, err := exec.LookPath("ffmpeg")
	if err != nil {
		return err
	}
	ffmpegPath = p
	return nil
}

func MergeAudioVideo(audioPath string, videoPath string, outputPath string) error {
	return ExecFFmpeg("-i", videoPath, "-i", audioPath, "-c", "copy", outputPath)
}

func ConvertAudio(inputAudio string, outputAudio string) error {
	return ExecFFmpeg("-i", inputAudio, "-c:v", "copy", "-c:a", "libmp3lame", outputAudio)
}

func ExecFFmpeg(args ...string) error {
	ffmpeg, err := GetFFmpeg()
	if err != nil {
		return err
	}
	cmd := exec.Command(ffmpeg, args...)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
