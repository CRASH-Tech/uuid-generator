package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/gofrs/uuid"
)

var (
	count int
)

func init() {
	flag.IntVar(&count, "n", 1, "number of UUID's for generate")
	flag.Parse()
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := 1; i <= count; i++ {
		fmt.Println(genUUIDv1(genRandomMac()))
	}

}

func genUUIDv1(mac string) string {
	fn := func() (net.HardwareAddr, error) {
		addr, err := net.ParseMAC(mac)
		if err != nil {
			fmt.Println("lol")
			panic(err)
		}
		return addr, nil
	}

	var g *uuid.Gen
	var err error
	var id uuid.UUID

	g = uuid.NewGenWithHWAF(fn)
	if g == nil {
		panic("unexpectedly nil")
	}

	id, err = g.NewV1()
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(id)
}

func genRandomMac() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	// Set the unicast bit
	buf[0] <<= 0

	mac := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	return mac
}
