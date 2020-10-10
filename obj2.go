package main

import "fmt"

type Per struct {
	Name string
	Age int
	Sex string
}

func (p *Per) Show() string {
	return fmt.Sprintf("my name is %s, Im %d years old", p.Name, p.Age)
}

func (p *Per) SetSex(sex string) {
	p.Sex = sex
}

func main(){
	teacher := Per{"sunyang", 19, "M"}
	fmt.Println(teacher.Show())

	student := Per{Name: "sunyang_one", Age: 19}
	student.SetSex("Male")
	fmt.Println(student.Show())
}