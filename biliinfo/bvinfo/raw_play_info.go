package bvinfo

type RawPlayInfo struct {
	Data rawPlayInfoData `json:"data"`
}

type rawPlayInfoData struct {
	Dash rawPlayInfoDataDash `json:"dash"`
}

type rawPlayInfoDataDash struct {
	Duration int32                      `json:"duration"` // seconds
	Video    []rawPlayInfoDataDashVideo `json:"video"`
	Audio    []rawPlayInfoDataDashAudio `json:"audio"`
}

type rawPlayInfoDataDashVideo struct {
	BaseUrl   string   `json:"baseUrl"`
	BackUpUrl []string `json:"backupUrl"`
	MimeType  string   `json:"mimeType"`
	Width     int32    `json:"width"`
	Heihgt    int32    `json:"height"`
	FrameRate string   `json:"frameRate"`
}

type rawPlayInfoDataDashAudio struct {
	BaseUrl   string   `json:"baseUrl"`
	BackUpUrl []string `json:"backupUrl"`
	MimeType  string   `json:"mimeType"`
}
