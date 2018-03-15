package main

//import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	l := 0
	r := x
	for {
		if l+1 >= r {
			break
		}

		m := l + (r-l)/2
		//fmt.Println("m changed: ", m)
		if m*m == x {
			return m
		}
		if m*m < x {
			//fmt.Println("l changed: ", l)
			l = m
		} else {
			r = m
			//fmt.Println("r changed: ", r)
		}

	}
	return l
}

/*
func main() {
	ret := mySqrt(65)
	fmt.Println(ret)
}
*/