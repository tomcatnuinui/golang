package main
import(
	"fmt"
)

type Dog struct {
	name string
	age int
	weight float32
}

func (d *Dog) say() string {
	info := fmt.Sprintf("小狗的信息是 名字=%v 年龄=%v 重量=%v\n",d.name,d.age,d.weight)
	return info
}

func main() {
	d := Dog {
		name : "tom",
		age : 2,
		weight : 3.6,
	}

	fmt.Println(d.say())
}