package main

import (
	"fmt"
	_ "log"
	"time"
)

type Serve struct {
	Name string  `json:"name"`
	Class string `json:"class"`
}

func main() {
	// 1. Reading file
	data := ReadFile()
	//log.Println(data)	

	//2. Write file
	writeFile(data)

	//3. Fliter word "admin"
	// dataFilter := filterAdmin(data)
	// if len(dataFilter) != 0 {
	// 	log.Println(dataFilter)
	// } else {
	// 	log.Println("Không có class nào chứa từ admin")
	// }

	//4. Add new Item to slice
	var newItem = Serve{"fileCustome","org.cofax.cds.FileServlet.Custome"}
	data = append(data, newItem)
	//log.Println(data)

	// 5. In ra màn hình các địa chỉ của các item trong slice trên
	for i := 0; i < len(data) ; i++ {
		fmt.Printf("Address of class %d : %p\n", i, &data[i].Class)
		fmt.Printf("Address of name %d : %p\n", i, &data[i].Name)
	}
 }
