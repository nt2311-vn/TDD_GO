package select_mod

import (
	"errors"
	"net/http"
	"time"
)

var ErrServerNotResp = errors.New("server not respond")

const tenSecondTimeOut = 10

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeOut)
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrServerNotResp
	}
}
