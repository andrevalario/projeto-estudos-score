package domain

import "encoding/json"

type ApiError struct {
	Id       uint64 `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Status   int    `json:"status"`
	Debug    *Debug `json:"debug,omitempty"`
	Original error  `json:"-"`
}

func (a *ApiError) Error() string {
	return a.Detail
}

type ErrorConfig struct {
	SkipCaller int
	WrapMsg    string
	Wrap       bool
}

type Debug struct {
	File      string   `json:"file,omitempty"`
	Line      int      `json:"line,omitempty"`
	RootCause string   `json:"root_cause,omitempty"`
	Stack     ApiStack `json:"stack,omitempty"`
}

func (a *ApiError) String() string {
	jsonToReturn, err := json.Marshal(a)
	if err != nil {
		return ""
	}

	return string(jsonToReturn)
}

func (a *ApiError) Unwrap() error {
	return a.Original
}

func (a *ApiError) ErrorAttributes() map[string]interface{} {
	attrs := make(map[string]interface{})

	attrs["error.file"] = a.Debug.File
	attrs["error.line"] = a.Debug.Line
	attrs["error.cause"] = a.Debug.RootCause
	attrs["error.stack"] = a.Debug.Stack.Error()
	attrs["error.title"] = a.Title
	attrs["error.detail"] = a.Detail

	return attrs
}

func (a *ApiError) Cause() error {
	return a.Debug.Stack
}

type ApiStack struct {
	error
}

func (a ApiStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Error())
}
