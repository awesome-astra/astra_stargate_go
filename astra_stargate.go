package astra_stargate

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Token   string
	DBID    string
	Region  string
	baseURL string
}

func NewBasicAuthClient(token, dbid, region string) *Client {
	baseURL := fmt.Sprintf("https://%s-%s.apps.astra.datastax.com", dbid, region)
	return &Client{
		Token:   token,
		DBID:    dbid,
		Region:  region,
		baseURL: baseURL,
	}
}

func (s *Client) GetURL() string {
	return s.baseURL
}

func (s *Client) APIPost(path string, payload *bytes.Buffer) (string, error) {
	url := fmt.Sprintf(s.baseURL+"%s", path)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	responsebody, err := s.doRequest(req)
	return responsebody, err
}

func (s *Client) APIPut(path string, payload *bytes.Buffer) (string, error) {
	url := fmt.Sprintf(s.baseURL+"%s", path)
	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return "", err
	}

	responsebody, err := s.doRequest(req)
	return responsebody, err
}

func (s *Client) APIDelete(path string) (string, error) {
	url := fmt.Sprintf(s.baseURL+"%s", path)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "", err
	}

	responsebody, err := s.doRequest(req)
	return responsebody, err
}

func (s *Client) APIGet(path string) (string, error) {
	url := fmt.Sprintf(s.baseURL+"%s", path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	responsebody, err := s.doRequest(req)
	return responsebody, err
}

func (s *Client) doRequest(req *http.Request) (string, error) {
	req.Header.Set("Authorization", "Bearer: "+s.Token)
	req.Header.Set("x-cassandra-token", s.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	bodystring := string(body)

	if err != nil {
		return "", err
	}

	if bodystring == "" {
		bodystring = "Response code: " + fmt.Sprintf("%d", resp.StatusCode)
	}

	return bodystring, nil
}
