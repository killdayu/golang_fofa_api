package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fofa_query_input := bufio.NewReader(os.Stdin)
	str, err := fofa_query_input.ReadString('\n')
	if err != nil {
		fmt.Println("input error")
	}
	fofa_query := base64.StdEncoding.EncodeToString([]byte(str))
	//fmt.Println(fofa_query)
	s := New(Config.Fofa_email, Config.Fofa_key, fofa_query)
	info, err := s.APIInfo()
	if err != nil {
		fmt.Println("network err")
	}
	fmt.Printf("%v", info.Results)
}
