// Package httpchecker provides a checker implementation specifically tailored for HTTP endpoints.
package httpchecker

import (
	"net/http"
	"time"

	"github.com/gabrielfioravante/bahc-go/internal/checker"
)

const defaultTimeout = 10 * time.Second

type HttpChecker struct {
	config *HttpCheckerConfig
	client *http.Client
	req    *http.Request
}

func NewHttpChecker(config *HttpCheckerConfig) checker.Checker {
	req, err := http.NewRequest(string(config.Method), config.Url, nil)

	if err != nil {
		panic(err)
	}

	if config.Authorization != "" {
		req.Header.Set("Authorization", config.Authorization)
	}

	timeout := defaultTimeout

	if config.Timeout > 0 {
		timeout = config.Timeout
	}

	client := &http.Client{
		Timeout: timeout,
	}

	return &HttpChecker{
		config: config,
		req:    req,
		client: client,
	}
}

func (hc *HttpChecker) Run() (checker.Result, error) {
	res, err := hc.client.Do(hc.req)
	var checkerResult checker.Result

	if err != nil {
		checkerResult.SetUnavailable(hc.config.ID, err.Error())
		return checkerResult, nil
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		checkerResult.SetUnavailable(hc.config.ID, res.Status)
		return checkerResult, nil
	}

	checkerResult.SetAvailable(hc.config.ID)

	return checkerResult, nil
}

func (hc *HttpChecker) GetID() string {
	return hc.config.ID
}

func (hc *HttpChecker) GetInterval() time.Duration {
	return hc.config.Interval
}
