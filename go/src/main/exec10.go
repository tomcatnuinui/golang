package main
import(
	"fmt"
)

type Movie struct {
	Name string
	Rating float32
}
func main() {
	// m := Movie {
	// 	Name: "skr",
	// 	Rating: 10,
	// }
	// fmt.Println(m.Name,m.Rating)
	var m Movie
	m.Name = "skr"
	m.Rating = 10
	fmt.Println(m.Name,m.Rating)
}