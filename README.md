![ted](https://user-images.githubusercontent.com/11734309/76626866-e329ca80-6574-11ea-87da-b158d6fff768.png)

# ted
Go bindings for the Telegram Bot API

## Usage

Ted's API is inspired by the Go standard library's `net/http` package.

Create a `ted.Bot` with your bot token and HTTP client, then build your request and send it:

```go
bot := ted.Bot{
    Token:      "BOT_TOKEN",
    HTTPClient: http.DefaultClient,
}
req := ted.SendMessageRequest{
    ChatID:    123,
    Text:      "*Hello, World*",
    ParseMode: "Markdown",
}
res, err := bot.Do(req)
if err != nil {
    if response, ok := err.(ted.Response); ok {
        // handle Telegram error
    }
    // handle HTTP error
}
```

The bot also handles making multiple requests concurrently:

```go
answerCallbackQueryRequest := ted.AnswerCallbackQueryRequest{
    CallbackQueryID: "abc",
}
sendMessageRequest := ted.SendMessageRequest{
    ChatID:    123,
    Text:      "*Hello, World*",
    ParseMode: "Markdown",
}
// err will be nil if all requests were successful
responses, err := bot.DoMulti(answerCallbackQueryRequest, sendMessageRequest)
```

To use the result, unmarshal it just as you would a HTTP response body:

```go
getMe := GetMeRequest{}
res, err := bot.Do(getMe)
if err != nil {
    panic(err)
}
var me User
err = json.Unmarshal(res.Result, &me)
if err != nil {
    panic(err)
}
```
