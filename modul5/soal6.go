package main
import "fmt"
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	hasil := HitungPangkat(x, y)
	fmt.Printf("%d ", hasil)
}

func HitungPangkat (x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * HitungPangkat(x, y-1)
	}
}
