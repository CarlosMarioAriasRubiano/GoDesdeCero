package variables

import (
	"fmt"
)

func ShowEntires() {
	var intComun int
	intOf32 := int32(10)
	intOf64 := int64(100)

	fmt.Println("intComun = ", intComun)
	fmt.Println("intOf32 = ", intOf32)
	fmt.Println("intOf64 = ", intOf64)
}
