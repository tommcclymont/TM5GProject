package logger

import "time"

type LogInfo struct {
	//request info
	Method  string
	Uri     string
	IpAddr  string
	ReqSize int
	ReqBody interface{}
	// response info
	RspCode  int
	RspSize  int64
	Duration time.Duration
}
