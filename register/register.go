package register

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int64  `json:"age"`
	PhoneNumber string `json:"phone_number"`
}

func Register() {
	var f, l, p string
	var a int64
	print("Ismingizni kiriting: ")
	fmt.Scan(&f)
	print("Familiyangizni kiriting: ")
	fmt.Scan(&l)
	print("Yoshingizni kiriting: ")
	fmt.Scan(&a)
	print("Telefon raqamingizni kiriting: ")
	fmt.Scan(&p)

	info := User{
		FirstName:   f,
		LastName:    l,
		Age:         a,
		PhoneNumber: p,
	}

	t, err := json.Marshal(info)
	if err != nil {
		log.Println("Error marshaling json user info", err)
	}
	err = ioutil.WriteFile("./register/user_info.json", t, 0600)
	if err != nil {
		log.Println("Error writer json file", err)
	}
	fmt.Println(string(t))

}
