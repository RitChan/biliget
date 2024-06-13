package biliinfo

type RawInitialState struct {
	BVID      string                   `json:"bvid"`
	P         int                      `json:"p"` // 分P
	VideoData rawInitialStateVideoData `json:"videoData"`
	UpData    rawInitialStateUpData    `json:"upData"`
	Tags      []rawInitialStateTags    `json:"tags"`
}

type rawInitialStateVideoData struct {
	NumVideos      int                             `json:"videos"`
	PictureURL     string                          `json:"pic"` // 视频封面URL
	Title          string                          `json:"title"`
	CreateTimeUnix int                             `json:"ctime"` // 视频创建时间(unix时间戳)
	Stat           rawInitialStateVideoDataStat    `json:"stat"`
	Pages          []rawInitialStateVideoDataPages `json:"pages"`
}

type rawInitialStateVideoDataStat struct {
	View      int `json:"view"`
	Like      int `json:"like"`
	Reply     int `json:"reply"`
	Favourite int `json:"favourite"`
	Coin      int `json:"coin"`
}

type rawInitialStateVideoDataPages struct {
	PageID    int    `json:"page"`
	PageTitle string `json:"part"`
}

type rawInitialStateUpData struct {
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	ProfileURL string `json:"face"`
}

type rawInitialStateTags struct {
	TagName string `json:"tag_name"`
}
