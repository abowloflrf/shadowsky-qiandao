package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

// ShadowskyConfig configuration for checkin
type ShadowskyConfig struct {
	URL       string
	UserAgent string
}

// Shadowsky shadowsky instance
type Shadowsky struct {
	email       string
	password    string
	config      *ShadowskyConfig
	restyClient *resty.Client
}

// CheckinResult 签到结果
// 成功： {Msg:获得了 259 MB流量. Ret:1}
// 已签到： {Msg:您似乎已经签到过了... Ret:1}
type CheckinResult struct {
	Msg string `json:"msg"`
	Ret int    `json:"ret"`
}

// NewShadowsky create new instance of shadowsky client
func NewShadowsky(email string, password string, c *ShadowskyConfig) (*Shadowsky, error) {
	s := &Shadowsky{
		email:       email,
		password:    password,
		config:      c,
		restyClient: resty.New().SetTimeout(5 * time.Second),
	}
	s.restyClient.SetHeaders(
		map[string]string{
			"User-Agent": s.config.UserAgent,
			"Origin":     s.config.URL,
		})
	if err := s.login(); err != nil {
		return nil, err
	}
	return s, nil
}

// login to shadowsky and save sessionid to http client
func (s *Shadowsky) login() error {
	resp, err := s.restyClient.R().
		SetHeader("Referer", s.config.URL+"/auth/login").
		SetFormData(map[string]string{
			"email":       s.email,
			"passwd":      s.password,
			"remember_me": "week",
		}).
		Post(s.config.URL + "/auth/login")
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
		Get(s.config.URL + "/user")
	if err != nil {
		return "", err
	}
	// parse html to user info
	return resp.String(), nil
}

// Checkin to get daily data
func (s *Shadowsky) Checkin() (*CheckinResult, error) {
	resp, err := s.restyClient.R().Post(s.config.URL + "/user/checkin")
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

// Parse parse the checkin response message and get the number of free data in MB just got
func (cr *CheckinResult) Parse() int {
	compRegex := regexp.MustCompile(`^获得了 (\d+) MB流量$`)
	res := compRegex.FindStringSubmatch(cr.Msg)
	if len(res) != 2 {
		return 0
	}
	n, err := strconv.Atoi(res[1])
	if err != nil {
		return 0
	}
	return n
}
