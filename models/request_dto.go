package models

type UppercaseRequest struct {
	S string `json:"s"`
}

type UpperceResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}
