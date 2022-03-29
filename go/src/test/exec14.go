package main
import(
	"fmt"
)

type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func (s *Student) test() {
	fmt.Println(s.Name)
}

func main(){
	s := Student{
		Name : "tom",
		Age : 20,
	}

	a := 0
	fmt.Println(&s)
	fmt.Println(&a)
	s.test()
}