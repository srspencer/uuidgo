package main

import (
"fmt"
"io/ioutil"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	var path="/proc/sys/kernel/random/uuid"

	b, e := ioutil.ReadFile(path)
	check(e)	

	u := string(b)
	fmt.Printf(u)
	fmt.Printf("\n")
}
