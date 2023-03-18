package IPC

import (
	"context"
)

// Service struct
type Service struct {
	ctx context.Context
}

// NewService creates a new IPC Service struct.
func NewService() *Service {
	return &Service{}
}

// Represents IPC shared data.
type data map[string]interface{}

// Expected form of Service request parameters.
type Request struct {
	Module string `json:"module"`
	Method string `json:"method"`
	Data   data   `json:"data"`
}

// Expected form of Service response data.
type response struct {
	Module string `json:"module"`
	Method string `json:"method"`
	Data   data   `json:"data"`
}

// Request – method through which all IPC service requests are received.
func (a *Service) Request(req Request) response {
	resData := data{}

	if req.Module == "ergo" {
		resData["content"] = processErgo(&req)
	} else {
		resData["content"] = "Module not recognised"
	}

	res := response{}
	res.Module = req.Module
	res.Method = req.Method
	res.Data = resData

	return res
}

// process request for Ergo.
func processErgo(req *Request) string {
	var content string

	if req.Method != "get" {
		content = "Method not allowed"
	} else {
		content = "I am thinking, therefore I exist"
	}

	return content
}
