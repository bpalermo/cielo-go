package utils

import "net/http"

// ReturnJsonResponse function for returning movies data in JSON format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) (int, error) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	return res.Write(resMessage)
}
