package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	Node  *snowflake.Node
	mutex sync.Mutex
)

func GetSnowId() int64 {
	if Node == nil {
		mutex.Lock()
		// Create a new Node with a Node number of 1
		var err error
		Node, err = snowflake.NewNode(1)
		if err != nil {
			fmt.Println(err)
			mutex.Unlock()
			return 0
		}
		mutex.Unlock()
	}

	// Generate a snowflake ID.
	id := Node.Generate()

	// Print out the ID in a few different ways.
	//fmt.Printf("Int64  ID: %d\n", id)
	//fmt.Printf("String ID: %s\n", id)
	//fmt.Printf("Base2  ID: %s\n", id.Base2())
	//fmt.Printf("Base64 ID: %s\n", id.Base64())
	//
	//// Print out the ID's timestamp
	//fmt.Printf("ID Time  : %d\n", id.Time())
	//
	//// Print out the ID's node number
	//fmt.Printf("ID Node  : %d\n", id.Node())
	//
	//// Print out the ID's sequence number
	//fmt.Printf("ID Step  : %d\n", id.Step())
	//
	//// Generate and print, all in one.
	//fmt.Printf("ID       : %d\n", Node.Generate().Int64())

	return int64(id)
}
