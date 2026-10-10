package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
GetAPIKeyFromRequest extracts the API key from the request headers.
It looks for the "Authorization: ApiKey {insert_api_key_here}" header and returns its value.
If the header is not present, it returns an error.
*/
func GetAPIKeyFromRequest(header http.Header) (string, error) {
	apiKey := header.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("api key not found in request headers")
	}
	
	if vals := strings.Split(apiKey, " "); len(vals) == 2 && vals[0] == "ApiKey" {
		return vals[1], nil
	}
	return "", errors.New("malformed API key in request headers")
}