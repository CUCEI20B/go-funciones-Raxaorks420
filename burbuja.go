package main

import "fmt"

func Burbuja(s []int64) {
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j-1] > s[j] {
				aux := s[j]
				s[j] = s[j-1]
				s[j-1] = aux
			}
		}
	}
	fmt.Println(s)
}

func main() {
	a := []int64{5, 10, 74, 55, 1, 66, 32, 19, 3, 87} // un slice de 10 elementos
	Burbuja(a)
}
