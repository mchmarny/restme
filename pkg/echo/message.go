package echo

import "encoding/json"

type Message struct {
	On      int64  `json:"on"`
	Message string `json:"msg"`
}

// String returns the JSON serialized representation of the object
func (m *Message) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
