package ted

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Request interface {
	doWith(bot Bot) (Response, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Bot struct {
	Token      string
	HTTPClient HTTPClient
}

func (b Bot) Do(request Request) (Response, error) {
	return request.doWith(b)
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

func (b Bot) DoMulti(requests ...Request) ([]Response, error) {
	responses := make([]Response, len(requests))
	errs := make([]error, len(requests))
	var wg sync.WaitGroup
	wg.Add(len(requests))
	for i, request := range requests {
		go func(i int, req Request) {
			res, err := b.Do(req)
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
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// ResponseParameters contains information about why a request was unsuccessful.
type ResponseParameters struct {
	// Optional. The group has been migrated to a supergroup with the
	// specified identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id"`

	// Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
	RetryAfter int `json:"retry_after"`
}

func (r Response) Error() string {
	return r.Description
}

type temporaryError interface {
	Temporary() bool
}

// doReq makes the provided http.Request to the Telegram Bot API.
// Currently hardcoded to retry up to 3 times if the error is retryable (implements Temporary() bool).
func (b Bot) doReq(req *http.Request) (Response, error) {
	retries := 3
	var res *http.Response
	var err error
	for {
		res, err = b.HTTPClient.Do(req)
		if err != nil {
			if terr, ok := err.(temporaryError); ok && terr.Temporary() && retries > 0 {
				retries--
				time.Sleep(1 * time.Second)
				continue
			}
			return Response{}, err
		}
		break
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

// doQuery makes a GET request to the Telegram Bot API with URL query parameters.
func (b Bot) doQuery(method string, params map[string]interface{}) (Response, error) {
	form := url.Values{}
	for k, v := range params {
		form.Set(k, fmt.Sprintf("%v", v))
	}
	u := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.Token, method)
	req, err := http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return Response{}, err
	}
	req.URL.RawQuery = form.Encode()
	return b.doReq(req)
}

// doJSON makes a POST request to the Telegram Bot API with a JSON body.
func (b Bot) doJSON(method string, request interface{}) (Response, error) {
	u := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.Token, method)
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
	return b.doReq(req)
}

func IsMessageNotModified(err error) bool {
	res, ok := err.(Response)
	if !ok {
		return false
	}
	return strings.Contains(res.Description, "message is not modified")
}
