package main

import "fmt"

type deliveryNotFoundError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e deliveryNotFoundError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e deliveryNotFoundError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
