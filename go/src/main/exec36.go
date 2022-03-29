package main
import(
	"fmt"
	"strconv"
)

type Movie struct {
	Name string
	Rating float64
}
func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating,'f',1,64)
	return m.Name + " "+ r
}

func main() {
	a := Movie {
		Name: "SKY",
		Rating: 172.23,
	}
	fmt.Println(a.summary())
}

