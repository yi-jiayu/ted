package ted

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMessageNotModified(t *testing.T) {
	var err error = Response{
		OK:          false,
		ErrorCode:   400,
		Description: "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message",
	}
	assert.True(t, IsMessageNotModified(err))
}

type tempError struct {
}

func (r tempError) Error() string {
	return ""
}

func (r tempError) Temporary() bool {
	return true
}

type result struct {
	res *http.Response
	err error
}

type httpClient struct {
	results []result
}

func (h *httpClient) Do(*http.Request) (*http.Response, error) {
	var result result
	result, h.results = h.results[0], h.results[1:]
	return result.res, result.err
}

func TestBot_doReq(t *testing.T) {
	t.Run("retries up to 3 times", func(t *testing.T) {
		client := &httpClient{
			results: []result{
				{
					err: tempError{},
				},
				{
					err: tempError{},
				},
				{
					err: tempError{},
				},
				{
					res: &http.Response{Body: ioutil.NopCloser(strings.NewReader(`{"ok":true}`))},
				},
			},
		}
		bot := Bot{HTTPClient: client}
		response, err := bot.doReq(nil)
		assert.NoError(t, err)
		assert.True(t, response.OK)
	})
	t.Run("returns error after running out of retries", func(t *testing.T) {
		client := &httpClient{
			results: []result{
				{
					err: tempError{},
				},
				{
					err: tempError{},
				},
				{
					err: tempError{},
				},
				{
					err: tempError{},
				},
			},
		}
		bot := Bot{HTTPClient: client}
		_, err := bot.doReq(nil)
		assert.Error(t, err)
	})
}
