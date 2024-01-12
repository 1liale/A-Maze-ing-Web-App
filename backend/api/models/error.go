package models

type InternalError struct {
	Code   int    `json:"code"`
	Msg    string `json:"message"`
	Reason string `json:"reason"`
}

func (e *InternalError) Error() string {
	return e.Reason
}
