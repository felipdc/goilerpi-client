package goilerpi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetUsers() ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	users := []User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) GetUser(id string) ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	users := []User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
