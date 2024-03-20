Что выведет программа? Объяснить вывод программы.

package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}


Ответ:
error

Когда test() возвращает nil, то она возвращает nil указатель типа *customError и когда этот nil указатель присваивается переменной типа error, и err хранит 2 значения: тип *customError и значение nil.
И поэтому err != nil будет истиной, так как для того чтобы интерфейс считался nil нужно чтобы его тип и значения были равын nil.
