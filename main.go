package main

import (
	"fmt"
	"os"
)

func main() {
	err, resp := os.ReadFile("schedule.html")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	// const gotSportURL = "https://system.gotsport.com/org_event/events/54541/schedules?team=4218655"
	fmt.Println("hello soccer")
	// resp, err := http.Get(gotSportURL)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// err = os.WriteFile("schedule.html", body, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("ayo", string(body))
}
