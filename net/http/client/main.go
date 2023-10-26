package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := new(http.Client)
	req, _ := http.NewRequest("POST", "http://localhost:8080/test", bytes.NewBuffer([]byte("{\"test\":\"i am client\"}")))
	req.Header["test"] = []string{"t1", "t2"}
	do, err := client.Do(req)
	if err != nil {
		return
	}
	b, _ := io.ReadAll(do.Body)
	fmt.Println(string(b))
}
