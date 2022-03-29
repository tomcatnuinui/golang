package main
import(
	"fmt"
)

type Superhero struct {
	Name   string
	Age    int
	Address Address
}

type Address struct {
	Number   int
	Street   string
	City     string
}

func main () {
	e := Superhero {
		Name: "skr",
		Age:   28,
		Address: Address{
			Number:  99,
			Street:  "jinghua",
			City:    "china",
		},
	}
	fmt.Printf("%+v\n",e)
	fmt.Println(e.Address.Street)
}