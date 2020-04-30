package export

import "fmt"

func Write2Print(path, data string) {
	fmt.Println("########### Print Path ###########")
	fmt.Println(path)
	fmt.Println("########### Print Start ###########")
	fmt.Println(data)
	fmt.Println("########### Print End ###########")
}
