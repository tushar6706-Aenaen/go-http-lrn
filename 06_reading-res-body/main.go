package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseStructure struct{
	Fact string `json:"fact"`
	Length int `json:"length"`
}


func main() {
	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		fmt.Println(res.Status)
		return
	}

	bodyBytes , err := io.ReadAll(res.Body) 
	if err != nil{
		fmt.Println("read body failed",err)
		return
	}


	var data ResponseStructure

	err = json.Unmarshal(bodyBytes,&data) 
	if  err != nil {
		fmt.Println("json unMarshal failed")
	}

	fmt.Println(data.Fact)
}