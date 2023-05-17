package common

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewResult() *Result {
	return &Result{}
}
func (r *Result) Success(data any) {
	r.Code = 200
	r.Message = "success"
	r.Data = data
}
func (r *Result) Fail(err error) {
	r.Code = 500
	r.Message = "Fail"
}
func (r *Result) Deal(data any, err error) *Result {
	if err != nil {
		r.Fail(err)
		return r
	}
	r.Success(data)
	return r
}
