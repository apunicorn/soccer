package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("hello soccer")
	resp, err := http.Get("https://system.gotsport.com/org_event/events/54541/schedules?team=4218655")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("ayo", string(body))
}
