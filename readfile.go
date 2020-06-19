package main

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFile()[] Serve{
	file, _ := ioutil.ReadFile("data.json")
	var data [] Serve 
	json.Unmarshal([]byte(file), &data)
	return data
}