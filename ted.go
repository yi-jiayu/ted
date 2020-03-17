package ted

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Request interface {
	doWith(bot Bot) (Response, error)
}

type Bot struct {
	Token      string
	HTTPClient *http.Client
}

func (c Bot) Do(request Request) (Response, error) {
	return request.doWith(c)
}

type MultiError []error

func (m MultiError) Error() string {
	total := len(m)
	errored := 0
	for _, err := range m {
		if err != nil {
			errored++
		}
	}
	return fmt.Sprintf("%d out of %d requests were unsuccessful", errored, total)
}

func (c Bot) DoMulti(requests ...Request) ([]Response, error) {
	responses := make([]Response, len(requests))
	errs := make([]error, len(requests))
	var wg sync.WaitGroup
	wg.Add(len(requests))
	for i, request := range requests {
		go func(i int, req Request) {
			res, err := c.Do(req)
			if err != nil {
				errs[i] = err
			} else {
				responses[i] = res
			}
			wg.Done()
		}(i, request)
	}
	for _, err := range errs {
		if err != nil {
			return nil, MultiError(errs)
		}
	}
	return responses, nil
}

type Response struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func (r Response) Error() string {
	return r.Description
}

func (c Bot) doJSON(method string, request interface{}) (Response, error) {
	u := fmt.Sprintf("https://api.telegram.org/bot%s/%s", c.Token, method)
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(request)
	if err != nil {
		return Response{}, err
	}
	req, err := http.NewRequest(http.MethodPost, u, &body)
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return Response{}, err
	}
	if !response.OK {
		return Response{}, response
	}
	return response, nil
}

func IsMessageNotModified(err error) bool {
	res, ok := err.(Response)
	if !ok {
		return false
	}
	return strings.Contains(res.Description, "message is not modified")
}
