package testutils

import "encoding/json"

func ObjectToJSON(body any) string {
	result, _ := json.Marshal(body)
	return string(result)
}
