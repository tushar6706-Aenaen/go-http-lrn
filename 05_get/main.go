package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/posts/1"

	res, er := http.Get(url)

	if er != nil {
		fmt.Println("error fetching data:", er)
		return
	}

 
	defer res.Body.Close()
	fmt.Println("status code:", res.StatusCode)
	fmt.Println("status:", res.Status)




}
