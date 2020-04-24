package main


var b interface{}

func main() {

	a := "aaa"

	b = a
	print(b.(int))
}
