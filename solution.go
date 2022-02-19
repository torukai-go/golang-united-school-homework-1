package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	helloMessege := emoji.Sprint("hello :world_map:")
	return helloMessege
}
