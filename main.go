package main

import (
	"flag"
	"fmt"
	"github.com/jaakoob/ucp"
	"math/rand"
	"time"
)

func main() {
	// read username, password, host and port from command line flags
	userPtr := flag.String("user", "none", "username for smsc")
	passwordPtr := flag.String("password", "none", "password for smsc")
	hostPtr := flag.String("host", "none", "hostname or ip address of smsc")
	portPtr := flag.Int("port", 5001, "port of smsc host")

	// read message from command line flags
	messagePtr := flag.String("message", "Test", "Message you want to send")
	fromNumberPtr := flag.String("from", "123", "Number used to send messages")
	toNumberPtr := flag.String("to", "456", "Number to send the message to")

	// create a random accescode
	rand.Seed(time.Now().UnixNano())
	accessCode := RandStringBytes(10)

	opt := &ucp.Options{
		Addr: *hostPtr + ":" + string(*portPtr),
		User: *userPtr,
		Password: *passwordPtr,
		AccessCode: accessCode,
	}
	client := ucp.New(opt)
	client.Connect()
	defer client.Close()
	ids, err := client.Send(*fromNumberPtr, *toNumberPtr, *messagePtr)

	fmt.Println(err)
	fmt.Println(ids)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
