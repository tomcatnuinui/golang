package main
import(
	"fmt"
	"reflect"
)
type Drink struct {
	Name  string
	Ice   bool
}

func main() {
	a := Drink {
		Name:  "skr",
		Ice:   true,
	}

	b := Drink {
		Name:   "sky",
		Ice:    true,
	}
	fmt.Printf("%+v\n",a)
	fmt.Printf("%+v\n",b)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	if a == b {
		fmt.Println("a and b are the same")
	} else {
		fmt.Println("a and b no same")
	}
}