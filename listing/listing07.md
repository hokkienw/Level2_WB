Что выведет программа? Объяснить вывод программы.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}


Ответ:

Порядок вывода значений в данной программе будет варьироваться при каждом запуске программы, но в конце приведёт к бесконечному выводу 0.
Сначала мы создём 2 канала, в которые записываются целочисленные значения.
Затем в функции merge происходит их объединение в один канал (select будет ждать из какого канала данные придут быстрее) и затем происходит чтение из этого канала.
Когда каналы буду закрыты, то в канал 'с' будут постоянно поступать данные нулевого значения для данного типа int (это 0). И таким образом программа не выйдет из цикла for.
