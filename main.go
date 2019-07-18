package main

import (
	"flag"
	"fmt"
	"github.com/jaakoob/ucp"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// read username, password, host and port from command line flags
	userPtr := flag.String("user", "none", "username for smsc")
	passwordPtr := flag.String("password", "none", "password for smsc")
	hostPtr := flag.String("host", "1.1.1.1", "hostname or ip address of smsc")
	portPtr := flag.Int("port", 5001, "port of smsc host")

	// read message from command line flags
	messagePtr := flag.String("message", "Test", "Message you want to send")
	fromNumberPtr := flag.String("from", "Hallo", "Number used to send messages")
	toNumberPtr := flag.String("to", "Hallo", "Number to send the message to")

	// parse all given flags
	flag.Parse()

	// create a random accescode
	rand.Seed(time.Now().UnixNano())
	accessCode := RandStringBytes(10)


	fmt.Println("Address: " + *hostPtr + ":" + strconv.Itoa(*portPtr))
	fmt.Println("User: ", *userPtr)
	fmt.Println("Password: ", *passwordPtr)
	fmt.Println("AccessCode: ", accessCode)

	opt := &ucp.Options{
		Addr: *hostPtr + ":" + string(*portPtr),
		User: *userPtr,
		Password: *passwordPtr,
		AccessCode: accessCode,
	}
	client := ucp.New(opt)
	if err := client.Connect(); err != nil {
		fmt.Println("Cant connect")
		os.Exit(1)
	}
	defer client.Close()
	fmt.Println("From: ", *fromNumberPtr)
	fmt.Println("To: ", *toNumberPtr)
	fmt.Println("Message: ", *messagePtr)
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
