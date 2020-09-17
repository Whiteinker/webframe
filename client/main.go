// client/main.go
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Result struct {
	Num, Ans int
}

func main() {
	fmt.Print("1")
	time.Sleep(9000)
	fmt.Println("2")
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans)

}