package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go generateLoad()
	}
	select {}
}

func generateLoad() {
	for {
		_, err := http.Get("http://localhost:8080/....7..2.8.......6.1.2.5...9.54....8.........3....85.1...3.2.8.4.......9.7..6....")
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
		} else {
			fmt.Printf(".")
		}
	}
}
