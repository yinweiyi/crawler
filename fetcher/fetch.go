package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

//Get contents by http
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	get, err := http.Get(url)
	defer get.Body.Close()
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, fmt.Errorf("wrong status code: %d", get.StatusCode)
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}
