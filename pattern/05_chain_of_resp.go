 package main

 import "fmt"

 type Handler interface{
	 Handle(int)
	 SetNext(Handler)
	 SetCanHandele(int)
 }

 type ConcreteHandler struct{
	 canHandele  int
	 next Handler
 }

 func (h *ConcreteHandler) Handle(i int){
	if i < h.canHandele{
		fmt.Println("Request Handeled")
	} else if h.next != nil {
		fmt.Println("Passing request to next handler")
		h.next.Handle(i)
	} else {
		fmt.Println("No Handlers availibale")
	}
}

func (h *ConcreteHandler) SetNext(next Handler){
	h.next = next
}

func (h *ConcreteHandler) SetCanHandele(i int){
	h.canHandele = i
}


 func main() {
	handlerA := &ConcreteHandler{}
	handlerB := &ConcreteHandler{}
	handlerA.SetNext(handlerB)
	handlerA.SetCanHandele(10)
	handlerB.SetCanHandele(20)
	handlerA.Handle(15)
 }
