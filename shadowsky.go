package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

const addr = "https://www.shadowsky.icu"

type Shadowsky struct {
	restyClient *resty.Client
}

func NewShadowsky() (*Shadowsky, error) {
	s := &Shadowsky{
		restyClient: resty.New(),
	}
	s.restyClient.SetHeaders(
		map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.34 Safari/537.36 Edg/81.0.416.20",
			"Origin":     addr,
		})
	if err := s.login(); err != nil {
		return nil, err
	}
	return s, nil
}

// login to shadowsky and save sessionid to http client
func (s *Shadowsky) login() error {
	resp, err := s.restyClient.R().
		SetHeader("Referer", addr+"/auth/login").
		SetFormData(map[string]string{
			"email":       os.Getenv("SHADOWSKY_EMAIL"),
			"passwd":      os.Getenv("SHADOWSKY_PASSWORD"),
			"remember_me": "week",
		}).
		Post(addr + "/auth/login")
	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusOK {
		return nil
	}
	return fmt.Errorf("login: %d", resp.StatusCode())
}

// User get shadowsky user info
func (s *Shadowsky) User() (string, error) {
	resp, err := s.restyClient.R().
		Get(addr + "/user")
	if err != nil {
		return "", err
	}
	// parse html to user info
	return resp.String(), nil
}

// Checkin to get daily data
func (s *Shadowsky) Checkin() (*CheckinResult, error) {
	resp, err := s.restyClient.R().Post(addr + "/user/checkin")
	if err != nil {
		return nil, err
	}
	r := CheckinResult{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, err
	}
	log.Printf("checkin resp: code %d ,  body %s", resp.StatusCode(), r.Msg)
	return &r, nil
}
