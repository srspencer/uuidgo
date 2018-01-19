package uuid

import (
"io/ioutil"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func Gen() string {
	var path="/proc/sys/kernel/random/uuid"

	b, e := ioutil.ReadFile(path)
	check(e)	

	return string(b)
}
