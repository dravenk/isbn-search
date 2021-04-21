package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func OpacDotNlcDotCn(tok, isbcode string) []byte {

	URL := "http://opac.nlc.cn/F/" + tok
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, URL, nil)

	// if you appending to existing query this works fine
	q := req.URL.Query()
	q.Add("find_base", "NLC01")
	q.Add("find_base", "NLC09")
	q.Add("func", "find-m")
	q.Add("find_code", "ISB")

	q.Add("request", isbcode)

	req.URL.RawQuery = q.Encode()

	// fmt.Println(req.URL.String())

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// fmt.Println(string(body))

	return body
}
