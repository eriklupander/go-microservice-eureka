package model

/**
 * From http://thenewstack.io/make-a-restful-json-api-go/
 */

type JsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
