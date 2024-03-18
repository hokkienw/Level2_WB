package main

import "fmt"

type Person interface{
	Accept(visitor Visitor)
	GetName() string
}

type Student struct{
	Name string
}
func (s *Student) GetName() string{
	return s.Name
}

func (s *Student) Accept(visitor Visitor){
	visitor.VisitStudent(s)
}

type Teacher struct{
	Name string
}

func (t *Teacher) GetName() string{
	return t.Name
}

func (t *Teacher) Accept(visitor Visitor){
	visitor.VisitTeacher(t)
}

type Visitor interface{
	VisitStudent(student *Student)
	VisitTeacher(teacher *Teacher)
}
type WbTech struct{}

func (s *WbTech) VisitStudent(student *Student){
	fmt.Println("Student: ", student.GetName())
}

func (s *WbTech) VisitTeacher(teacher *Teacher){
	fmt.Println("Teacher: ", teacher.GetName())
}

func main(){
	exam := WbTech{}
	
	student := Student{"Rashit"}
	teacher := Teacher{"Ivan"}

	student.Accept(&exam)
	teacher.Accept(&exam)
}







