package biliinfo

type BvInfo struct {
	BVID       string
	Title      string
	Duration   float32 // in seconds
	PartID     int32   // 分P
	Up         upInfo
	Videos     []videoInfo
	Audios     []audioInfo
	Parts      []partInfo // 视频分P
	Statistics videoStat
}

type videoInfo struct {
	Url       string
	UrlBackup []string
	MimeType  string
	Width     uint32
	Height    uint32
	FrameRate float32
}

type audioInfo struct {
	Url       string
	UrlBackup []string
	MimeType  string
}

type upInfo struct {
	Name    string
	Profile string
	Sex     string
}

type videoStat struct {
	View      uint32 // 观看
	Like      uint32 // 点赞
	Coin      uint32 // 投币
	Favourite uint32 // 收藏
	Reply     uint32 // 评论
}

// 视频分P信息
type partInfo struct {
	PartId   int32
	PartName string
}
