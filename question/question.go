package question

import (
	"fmt"

	"gitlab.com/home_task/course"
	"gitlab.com/home_task/register"
)

func Question() {
	var str string
	print("Davom etishni xohlaysizmi? -> (ha yoki yo'q)")
	fmt.Scan(&str)
	if str == "Ha" || str == "ha" {
		course.Course()
	} else if str == "yo'q" || str == "Yo'q" || str == "yoq" {
		fmt.Println("Tashrifingiz uchun rahmat!!!")
		return
	} else {
		Question()
	}

}

func Question1() {
	var n int64
	print("Kurslardan birini tanlang -> (1, yoki 2 ...): ")
	fmt.Scan(&n)
	if n == 1 {
		course.Bootcamp()
	} else if n == 2 {
		course.Foundation()
	} else {
		fmt.Println("Nimadir xato, Qaytadan urinib ko'ring!")
		Question1()
	}
}

func Question2() {
	var str string
	print("Registratsiydan o'tasizmi -> (ha yoki yo'q): ")
	fmt.Scan(&str)
	if str == "Ha" || str == "ha" {
		register.Register()
	} else if str == "yoq" || str == "Yoq" || str == "Yo'q" || str == "yo'q" {
		fmt.Println("Bizning platformamizga tashirif buyurganingiz uchun tashakjur!!!")
		return
	} else {
		fmt.Println("Nimadir xato qaytadan urinib ko'ring!")
		Question2()
	}
}
