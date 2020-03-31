package ted

import (
	"encoding/json"
)

func (b Bot) GetMe() (User, error) {
	req := GetMeRequest{}
	res, err := b.Do(req)
	if err != nil {
		return User{}, err
	}
	var me User
	err = json.Unmarshal(res.Result, &me)
	if err != nil {
		return User{}, err
	}
	return me, nil
}

// GetWebhookInfo returns the current webhook status. Requires no parameters.
// On success, returns a WebhookInfo object. If the bot is using getUpdates,
// will return an object with the url field empty.
func (b Bot) GetWebhookInfo() (WebhookInfo, error) {
	req := GetWebhookInfoRequest{}
	res, err := b.Do(req)
	if err != nil {
		return WebhookInfo{}, err
	}
	var info WebhookInfo
	err = json.Unmarshal(res.Result, &info)
	if err != nil {
		return WebhookInfo{}, err
	}
	return info, nil
}
