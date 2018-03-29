package models

//easyjson:json
type Error struct {
	Message string `json:"message"`
}

func (r *Error) ErrorMsgJSON(err string) []byte {
	r.Message = err
	errBuf, _ := r.MarshalJSON()
	return errBuf
}
