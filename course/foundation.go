package course

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Foundation() {
	f, err := ioutil.ReadFile("./course/foundation.txt")

	if err != nil {
		log.Println("error read foundation.txt", err)
	}
	fmt.Println(string(f))
}
