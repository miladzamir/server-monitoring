package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	SERVER_Addr = "5.90.90.25"
)

var wg sync.WaitGroup
var remoteAddr = ""

func startMonitoring(ip string) {
	for {
		conn, err := net.DialTimeout("tcp", ip, time.Second*2)
		if err != nil {
			if remoteAddr != "" {
				fmt.Println(ip, "STOP!")
				remoteAddr = ""
				break
			}
			fmt.Println("Connecting...")

			insert("INSERT INTO log (remote_addr, local_addr, ping_at, live)VALUES ('" + ip + "', '" + SERVER_Addr + "', '" + time.Now().Format("2006-01-02 15:04:05") + "', 0)")
			time.Sleep(time.Second)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		if remoteAddr != "" {
			if conn.RemoteAddr().String() == remoteAddr {
				fmt.Println(remoteAddr, "STOP!")
				conn.Close()
				remoteAddr = ""
				break
			}
		}
		wg.Wait()
		insert("INSERT INTO log (remote_addr, local_addr, ping_at)VALUES ('" + conn.RemoteAddr().String() + "', '" + conn.LocalAddr().String() + "', '" + time.Now().Format("2006-01-02 15:04:05") + "')")
		fmt.Println(time.Now(), "Ok", conn.RemoteAddr(), conn.LocalAddr())
		time.Sleep(time.Second)

		defer conn.Close()
	}
}

func stopMonitoring(ip string) {
	wg.Add(1)
	remoteAddr = ip
	wg.Done()
}
