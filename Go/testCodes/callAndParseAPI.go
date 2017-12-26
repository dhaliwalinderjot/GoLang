package main

import (
	"json/encode"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() 
{
	response, err := http.Get("https://koinex.in/api/ticker")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var f interface{}
	
	err := json.Unmarshal(b, &f)
	

	fmt.Println(string(responseData))

}
