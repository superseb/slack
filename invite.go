package slack

import (
	"errors"
	"net/url"
)

// SendInvite sends an invite
func (api *Client) SendInvite(email string) error {
	values := url.Values{
		"email":      {email},
		"set_active": {"false"},
		"token":      {api.config.token},
	}
	response := &SlackResponse{}
	if err := post("users.admin.invite", values, response, api.debug); err != nil {
		return err
	}
	if !response.Ok {
		return errors.New(response.Error)
	}
	return nil
}
