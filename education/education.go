package education

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Education() {
	edu, err := ioutil.ReadFile("./education/inforamtion.txt")
	if err != nil {
		log.Println("Error read edu info: ", err)
	}
	fmt.Println(string(edu))
}
