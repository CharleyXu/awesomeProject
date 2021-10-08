package main

func main() {
	_, _, _, nickName := getName()
	println(nickName)
	exec()
}

func getName() (firstName, middleName, lastName, nickName string) {
	firstName = "A"
	middleName = "B"
	lastName = "C"
	nickName = "D"
	return
}

func exec() {
	f := func(x, y int) int {
		return x + y
	}
	a := f(1, 1)
	println(a)
	println(f(1, 1))
}
