package delivery

type Response struct {
	Message  string `json:"message"`
	Data     any    `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	Code  string `json:"code,omitempty"`
	Error string `json:"error"`
	Trace string `json:"trace,omitempty"`
}

var SuccessMessage = "ok"
