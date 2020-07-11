package main

import "fmt"

type notFoundError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e notFoundError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e notFoundError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

type internalError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e internalError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e internalError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
