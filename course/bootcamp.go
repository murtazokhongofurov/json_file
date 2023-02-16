package course

import (
	"fmt"
	"io/ioutil"
)



func Bootcamp() {
	b, err := ioutil.ReadFile("./course/bootcamp.txt")
	if err != nil {
		fmt.Println("error read bootcamp file", err)
	}
	fmt.Println(string(b))
}