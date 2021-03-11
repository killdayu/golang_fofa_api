package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://fofa.so"

type Client struct {
	fofa_email string
	fofa_key   string
	fofa_query string
}

func New(fofa_email string, fofa_key string, fofa_query string) *Client {
	return &Client{fofa_email: fofa_email, fofa_key: fofa_key, fofa_query: fofa_query}
}

type APIInfo struct {
	Results [][]string `results`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api/v1/search/all?email=%s&key=%s&qbase64=%s", BaseURL, s.fofa_email, s.fofa_key, s.fofa_query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
