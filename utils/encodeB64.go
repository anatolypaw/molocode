package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	a := fmt.Sprintf(`0104607009781495215E45z"%c93kxlE`, 29)
	fmt.Println(a)
	encoded := base64.StdEncoding.EncodeToString([]byte(a))
	fmt.Println(encoded)
}
