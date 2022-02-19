package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	helloMessege := emoji.Sprint("hello :world_map:")
	fmt.Println(helloMessege)
}
