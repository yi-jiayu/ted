package ted

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

type SendMessageRequest struct {
	ChatID      int         `json:"chat_id"`
	Text        string      `json:"text"`
	ParseMode   string      `json:"parse_mode"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

func (r SendMessageRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("sendMessage", r)
}

type ReplyMarkup interface {
	replyMarkup()
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i InlineKeyboardMarkup) replyMarkup() {}

func (i InlineKeyboardMarkup) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(data))
}

type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       string `json:"cache_time,omitempty"`
}

func (r AnswerCallbackQueryRequest) doWith(bot Bot) (Response, error) {
	return bot.doJSON("answerCallbackQuery", r)
}
