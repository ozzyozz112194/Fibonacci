package main

import "fmt"

func fib(x int) int {
	if x <= 3 {
		return 1
	} else {

		return fib(x-1) + fib(x-2)
	}
}

func Burbuja(args ...int) []int {

	for i := len(args); i > 0; i-- {
		for j := 1; j < i; j++ {
			if args[j-1] > args[j] {
				inter := args[j]
				args[j] = args[j-1]
				args[j-1] = inter

			}
		}
	}

	return args
}

func generadorPares() func() uint {
	i := uint(0) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint {
		var par = i
		var dif = i / 2
		if dif != 0 {
			i += 2
			return par
		} else {
			i += 3
			return 1
		}
	}
}

func incrementa(x, y *int) { // se recibe un puntero
	*x = *x + 1 // *x dereferenciación/acceso a la dirección de memoria
	*y = *x
}

func main() {
	a := []int{10, 15, 5, 2}
	l := len(a)
	Burbuja(a...)
	println("Numero mayor es :", a[l-1])
	println("Fibonacci es :", fib(15))
	nextPar := generadorPares()
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	fmt.Println(nextPar())
	x := 10

}
