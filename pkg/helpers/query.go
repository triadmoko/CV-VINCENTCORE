package helpers

import (
	"strings"
)

func URLQueryToParamsGin(queries map[string][]string) map[string]string {
	params := make(map[string]string)

	for k, v := range queries {
		params[k] = ""
		if len(v) > 0 {
			params[k] = v[0]
		}
	}

	return params
}
func URLQueryToParamsFiber(query string) map[string]string {
	params := make(map[string]string)
	if query == "" {
		return map[string]string{}
	}
	splitParams := strings.Split(query, "&")

	for _, param := range splitParams {
		ok := strings.Contains(param, "=")
		if !ok {
			continue
		}
		keyAndValues := strings.Split(param, "=")
		if len(keyAndValues) > 0 {
			params[keyAndValues[0]] = keyAndValues[1]
		}
	}

	return params

}
