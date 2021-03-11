package main

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	Fofa_email string `fofa_email`
	Fofa_key   string `fofa_key`
}{}

func init() {
	configor.Load(&Config, "config.yml")
}
