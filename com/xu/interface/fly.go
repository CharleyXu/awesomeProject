package main

type Bird struct {
}

func (b *Bird) fly() {
	println("flying")
}

type IFly interface {
	fly()
}

func main()  {
	var fly IFly = new(Bird)
	fly.fly()
}
