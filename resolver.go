package lambda

import (
	"os"
)

func resolver(h func(payload map[string]interface{}) ([]interface{}, error)) Response {
	response := Response{
		Payload: Payload{},
	}

	var event map[string]interface{}
	err := decode(os.Getenv("LAMBDA_EVENT"), &event)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	res, err := h(event)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	response.Payload.Success = res
	return response
}
