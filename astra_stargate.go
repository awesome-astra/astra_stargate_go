package astra_stargate

import (
	"bytes"
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"io"
=======
	"io/ioutil"
>>>>>>> c6e6fc19bb6a510d9f0314b1619dbea5d8eb678d
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

func (s *Client) APICall(path, payload string) error {
	url := fmt.Sprintf(s.baseURL+"%s", path)
	fmt.Println(url)
	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	method := "GET"
	body := bytes.NewBuffer(j)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	_, err = s.doRequest(req)
	return err
}

func (s *Client) doRequest(req *http.Request) (string, error) {
	req.Header.Set("Authorization", "Bearer: "+s.Token)
	req.Header.Set("x-cassandra-token", s.Token)
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	bodystring := string(body)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", bodystring)
	}
	return bodystring, nil
}
