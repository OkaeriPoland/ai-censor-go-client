package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Client struct {
	Client   *http.Client
	Url      string
	Username string
	Token    string
	Debug    bool
}

func CreateClient(token string, debug bool) (*Client, error) {
	httpClient := &http.Client{}
	return &Client{
		Client: httpClient,
		Url:    "https://ai-censor.okaeri.eu",
		Token:  token,
		Debug:  debug,
	}, nil
}

func (c *Client) Post(path string, body []byte) ([]byte, *http.Response, error) {
	reqUrl := c.Url + path
	r, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("fail when creating request for %v: %v", path, err)
	}
	r.Header.Set("Token", c.Token)
	r.Header.Set("Content-Type", "application/json")
	res, _ := c.Client.Do(r)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	if c.Debug {
		log.Printf("HTTP %v @ %v", res.StatusCode, reqUrl)
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err == nil && c.Debug {
		log.Printf("%s", responseData)
	}
	return responseData, res, nil
}

func (c *Client) Predict(content string) (PredictResponse, error) {
	queryByte, err := json.Marshal(PredictQuery{Phrase: content})
	if err != nil {
		return PredictResponse{}, errors.New("fail marshalling query for Predict: " + err.Error())
	}
	reqUrl := "/predict"
	body, resp, err := c.Post(reqUrl, queryByte)
	if err != nil {
		return PredictResponse{}, err
	}
	if resp.StatusCode != 200 {
		return PredictResponse{}, errors.New("wrong answer from API: " + strconv.Itoa(resp.StatusCode))
	}
	var result PredictResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return PredictResponse{}, errors.New("fail for unmarshalling Predict: " + err.Error())
	}
	if c.Debug {
		log.Printf("%v", result)
	}
	return result, err
}

type PredictQuery struct {
	Phrase string `json:"phrase"`
}

type PredictResponse struct {
	General General `json:"general"`
	Details Details `json:"details"`
	Elapsed Elapsed `json:"elapsed"`
}

type Details struct {
	BasicContainsHit bool    `json:"basic_contains_hit"`
	ExactMatchHit    bool    `json:"exact_match_hit"`
	AILabel          string  `json:"ai_label"`
	AIProbability    float64 `json:"ai_probability"`
}

type Elapsed struct {
	All        float64 `json:"all"`
	Processing float64 `json:"processing"`
}

type General struct {
	Swear     bool   `json:"swear"`
	Breakdown string `json:"breakdown"`
}
