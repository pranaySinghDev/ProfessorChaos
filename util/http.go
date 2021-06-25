package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallHTTP(dependencyURL string) (string, error) {
	resp, err := http.Get(dependencyURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	}
	return "", fmt.Errorf("error with status %d", resp.StatusCode)
}
