package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"net"

	"github.com/google/uuid"
)

var (
	count int
)

func init() {
	flag.IntVar(&count, "n", 1, "number of UUID's for generate")
	flag.Parse()
}

func main() {
	for i := 1; i <= count; i++ {
		fmt.Println(generateUUID1())
	}

}

func generateUUID1() string {
	mac := generateMac()
	uuid.SetNodeID([]byte(mac))

	result, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	return result.String()
}

func generateMac() net.HardwareAddr {
	buf := make([]byte, 6)
	var mac net.HardwareAddr

	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	// // Set the local bit
	// buf[0] |= 2

	mac = append(mac, buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	return mac
}
