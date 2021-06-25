package main

func main() {
	//Initial value is cero by default
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			//break
			continue
		}
		println("after continue")
	}
}
