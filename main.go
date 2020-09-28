package main

import "fmt"

func main() {
	var infinito int
	var e float64
	var auxMultiplicar int
	infinito = 10
	for n := 0; n <= infinito; n++ {
		if n == 0 {
			e = 1.0
		} else {
			auxMultiplicar = 1
			for aux := n; aux > 0; aux-- {
				auxMultiplicar = auxMultiplicar * aux
			}
			e = e + float64(1/float64(auxMultiplicar))
		}
	}

	fmt.Println(e)
}
