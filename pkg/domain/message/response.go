package message

import "time"

// custom code:
// 80 -> BAD_REQUEST
// 99 -> INTERNAL_SERVER_ERROR
// 0 -> SUCCESS

// annotation -> OMITEMPTY -> menghapus json property yang
// value null, atau kosong
type Response struct {
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	Error     string     `json:"error,omitempty"`      // nullable
	Data      any        `json:"data,omitempty"`       // nullable
	StartTime *time.Time `json:"start_time,omitempty"` // nullable
}
