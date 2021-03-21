package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	HOST = "88.99.104.53"
	PORT = "80"
)

func main() {
	i := check()

	for {
		func() {
			var d net.Dialer
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			conn, err := d.DialContext(ctx, "tcp", HOST+":"+PORT)
			if err != nil {
				log.Fatalf("Failed to dial: %v", err)
			}
			connected(conn, i)
		}()
	}
}
func check() int {
	var i int
	fmt.Println("Time between each check per second: ")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return i
}

func connected(conn net.Conn, i int) {

	fmt.Println(conn.LocalAddr())
	fmt.Println(time.Now().Format(time.ANSIC), "Alive...", HOST+":"+PORT)

	time.Sleep(time.Second * time.Duration(i))
	defer conn.Close()
}
