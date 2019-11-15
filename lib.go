package something

import (
	"fmt"

	// import badger
	_ "github.com/dgraph-io/badger"
)

// Hi says hi
func Hi() {
	fmt.Println("hi")
}
