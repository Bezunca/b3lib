package history

import (
	"io/ioutil"
	"net/http"
)

func download(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, &Not200StatusCode{response.StatusCode, string(body[:])}
	}

	return body, nil
}
