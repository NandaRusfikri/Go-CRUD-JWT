package structs

import "time"

type Component struct {
}

func (c *Component) GetTimeNow() string {

	t := time.Now()
	return string(t.Format("2006-01-02 15:04:05.999999"))
}

type Response struct {
	Count   int         `json:"count,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
