package main

import (
	"net/http"
	"time"
	"fmt"
	"ui"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
	ui.Run()
	time.Sleep(time.Second * 10)
}
