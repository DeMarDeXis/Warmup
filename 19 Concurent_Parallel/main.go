package main

import (
	"fmt"
)

func main() {
	//defer block
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(sum(2, 3))
	//deferValues()

	// goroutines block
	//runtime.GOMAXPROCS(1) //Кол во выполняемых горутин выполняемых одновременно => прога будет выполнятся конкурентно
	//fmt.Println(runtime.NumCPU()) //Кол -во логических ядер = кол-во горутин

	//#1
	//go showNumbers(100)

	//runtime.Gosched() // Переключалка Горутин вручнуб
	//time.Sleep(time.Second) //Здесь уже сам планировщик считает что нужно переключится
	fmt.Println("exit")

	makePanic()
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return
}

func deferValues() {
	for i := 0; i < 10; i++ { // 9 8 7 6 5 4 3 2 1 0
		defer fmt.Println("first", i)
	}
	for i := 0; i < 10; i++ { // 10 10 10 10 10 10 10
		defer func() {
			fmt.Println("second", i)
		}()
	}

	for i := 0; i < 10; i++ { //9-0
		k := i
		defer func() {
			fmt.Println("third", k)
		}()
	}
	for i := 0; i < 10; i++ { //9-0
		defer func(k int) {
			fmt.Println("fourth", k)
		}(i)
	}
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("some panic")
	fmt.Println("Unreachable code")
}
