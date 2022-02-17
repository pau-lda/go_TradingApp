package marketDataEndpoints

import "time"

type ServerTime struct {
	ServerTime int64 `json:"serverTime"`
}

func (s *ServerTime) GetServerTime() int64 {
	return s.ServerTime
}

func (s *ServerTime) GetServerTimeFormatted() string {
	return time.UnixMilli(s.ServerTime).String()
}
