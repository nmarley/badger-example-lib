package something

import (
	"expvar"
	"fmt"

	// import badger
	_ "github.com/dgraph-io/badger"
)

var NumReads *expvar.Int

func init() {
	NumReads = expvar.NewInt("badger_disk_reads_total")
}

// Hi says hi
func Hi() {
	fmt.Println("hi")
}
