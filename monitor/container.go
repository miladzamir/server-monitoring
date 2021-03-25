package monitor

import (
	DB "../database"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	SERVER_Addr = "5.90.90.25"
)

func StartMonitoring(ip string) {
	for {
		conn, err := connectToServer(ip)
		if err != nil {
			log.Fatal(err)
		}
		DB.Insert("INSERT INTO log (remote_addr, local_addr, ping_at)VALUES ('" + conn.RemoteAddr().String() + "', '" + conn.LocalAddr().String() + "', '" + time.Now().Format("2006-01-02 15:04:05") + "')")
		fmt.Println(time.Now(), "Ok", conn.RemoteAddr(), conn.LocalAddr())
		time.Sleep(time.Second)

		defer conn.Close()
	}
}
func connectToServer(ip string) (net.Conn, error) {
	for {
		conn, err := net.DialTimeout("tcp", ip, time.Second*2)
		if err != nil {
			fmt.Println("Connecting...")

			DB.Insert("INSERT INTO log (remote_addr, local_addr, ping_at, live)VALUES ('" + ip + "', '" + SERVER_Addr + "', '" + time.Now().Format("2006-01-02 15:04:05") + "', 0)")
			time.Sleep(time.Second)
			continue
		} else {
			return conn, err
		}
	}
}
