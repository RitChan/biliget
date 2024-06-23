package bihttp

type UserInfoResp struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    userInfo `json:"data"`
}

type userInfo struct {
	IsLogin bool   `json:"isLogin"`
	Face    string `json:"face"`
	Mid     int    `json:"mid"`
}

func GetUserInfoResp() (*UserInfoResp, error) {
	info := new(UserInfoResp)
	err := BiliGetUrlJson(urlUserInfo, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *UserInfoResp) IsLogin() bool {
	return s.Code == 0
}
