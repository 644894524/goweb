package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, err := http.Get("http://edu.ln139.com/apps/resource/")
	if err != nil {

	}

	byte,_ := ioutil.ReadAll(resp.Body)
	con := string(byte)
	fmt.Println(con)
}