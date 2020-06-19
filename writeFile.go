package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)
func writeFile(arrServe [] Serve) {
	emptyFile, err := os.Create("serve.json")
	if err != nil {
		log.Println("Tạo file thất bại")
		return
	}
	file , err := json.MarshalIndent(arrServe, "","  ")
	ioutil.WriteFile("serve.json", file, 0644)
	if err != nil {
		log.Println("Ghi file thất bại")
		defer emptyFile.Close()
		return
	}

	//log.Println(file)
	defer emptyFile.Close()
}