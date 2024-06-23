package goroutines

import (
	"biliget/biliinfo/bihttp"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type (
	QrcodeState struct {
		Message       string
		LoginState    LoginState
		QrcodeExpired bool // bilibili server side status && local status
		LoginExpired  bool // bilibili server side status
		QrKey         string
		QrUrl         string

		qctime     time.Time // qrcode creation time
		qrDuration int       // qrcode duration seconds
		potime     time.Time // qrcode poll time
		poInterval int       // qrcode poll interval
		lctime     time.Time // login check time
		lcInterval int       // login check interval, seconds
	}

	pollResult struct {
		key     string
		resp    *bihttp.QrCodePollResp
		cookies []*http.Cookie
		err     error
	}

	LoginState int32
)

const (
	ToScan    = LoginState(0)
	ToVerify  = LoginState(1)
	Succeeded = LoginState(2)
)

var (
	qrstate *QrcodeState
	// test
	pollLock    sync.Locker
	pollResults []*pollResult = make([]*pollResult, 0, 1)
)

func GetQrcodeState() (*QrcodeState, error) {
	s := getQrcodeState()

	if s.isTimeToPoll() {
		s.potime = time.Now()
		go pollLogin(s.QrKey)
	}

	switch s.LoginState {
	case Succeeded:
		if s.lcExpired() && loginExpired() {
			s.reset()
			s.LoginExpired = true
		}
	case ToScan, ToVerify:
		if s.qrExpired() {
			s.reset()
			s.QrcodeExpired = true
		} else if poResult := fetchLatestPollResult(s.QrKey); poResult != nil {
			if poResult.err != nil {
				return nil, poResult.err
			}
			switch poResult.resp.Data.Code {
			case 0:
				s.Message = "登录成功"
				s.LoginState = Succeeded
				s.QrcodeExpired = false
				s.LoginExpired = false
				bihttp.SetCookies(poResult.cookies)
			case 86038:
				s.reset()
				s.QrcodeExpired = true
			case 86090:
				s.LoginState = ToVerify
				s.Message = "请在手机确认"
			case 86101:
				// 未扫码
			default:
				return nil, fmt.Errorf("polling error, code = %d", poResult.resp.Data.Code)
			}
		}
	}
	return s, nil
}

func getQrcodeState() *QrcodeState {
	if qrstate == nil {
		qrstate = new(QrcodeState)
		qrstate.reset()
	}
	return qrstate
}

func pollLogin(qrkey string) {
	resp, cookies, err := bihttp.BiliQrcodePoll(qrkey)
	result := pollResult{qrkey, resp, cookies, err}
	writePollResult(&result)
}

func writePollResult(poResult *pollResult) {
	pollLock.Lock()
	defer pollLock.Unlock()
	pollResults = append(pollResults, poResult)
}

func fetchLatestPollResult(qrkey string) *pollResult {
	pollLock.Lock()
	defer pollLock.Unlock()
	if len(pollResults) == 0 {
		return nil
	}
	var result *pollResult
	for _, r := range pollResults {
		if r.key == qrkey {
			result = r
		}
	}
	pollResults = make([]*pollResult, 0, 1)
	return result
}

func loginExpired() bool {
	resp, err := bihttp.GetUserInfoResp()
	if err != nil {
		return false
	}
	return resp.IsLogin()
}

func (s *QrcodeState) reset() error {
	resp, err := bihttp.BiliGetQrcodeUrl()
	if err != nil {
		return err
	}
	s.Message = "扫码登录"
	s.LoginState = ToScan
	s.QrcodeExpired = false
	s.LoginExpired = false
	s.QrKey = resp.Data.QrcodeKey
	s.QrUrl = resp.Data.Url

	s.qctime = time.Now()
	s.qrDuration = 180
	s.lctime = time.Now()
	s.lcInterval = 600
	s.potime = time.Now()
	s.poInterval = 3
	return nil
}

func (s *QrcodeState) qrExpired() bool {
	return time.Now().After(s.qctime.Add(time.Duration(s.qrDuration * 1e9)))
}

func (s *QrcodeState) lcExpired() bool {
	return time.Now().After(s.lctime.Add(time.Duration(s.lcInterval * 1e9)))
}

func (s *QrcodeState) isTimeToPoll() bool {
	return time.Now().After(s.potime.Add(time.Duration(s.poInterval * 1e9)))
}
