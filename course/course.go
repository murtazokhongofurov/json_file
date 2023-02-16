package course

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Courses struct {
	Id         int64  `json:"course_id"`
	CourseName string `json:"course_name"`
}

func Course() {
	info, err := ioutil.ReadFile("./course/courses.json")
	if err != nil {
		log.Println("error read json file", err)
	}
	var courses []Courses
	err = json.Unmarshal([]byte(info), &courses)
	if err != nil {
		fmt.Println("error while unmarshaling json", err)
	}
	for _, r := range courses {
		fmt.Println(r.Id, r.CourseName)
	}

}
