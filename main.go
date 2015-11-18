package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const service = "SOLVER_SERVICE"

func main() {
	// http://kubernetes.io/v1.1/docs/user-guide/services.html#environment-variables
	var svcHost = os.Getenv(service + "_SERVICE_HOST")
	var svcPort = os.Getenv(service + "_SERVICE_PORT")
	var svcAddr = svcHost + ":" + svcPort
	fmt.Printf("address of target service: %s\n", svcAddr)

	for i := 0; i < 10; i++ {
		go generateLoad(svcAddr)
	}
	select {}
}

func generateLoad(svcAddr string) {
	for {
		_, err := http.Get("http://" + svcAddr + "/....7..2.8.......6.1.2.5...9.54....8.........3....85.1...3.2.8.4.......9.7..6....")
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
		} else {
			fmt.Printf(".")
		}
	}
}
