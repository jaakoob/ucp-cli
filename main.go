package main

import (
	"flag"
	"fmt"
	"github.com/jaakoob/ucp"
	"math/rand"
	"os"
	"strconv"
	"time"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port			int
	Address			string
	Username		string
	Password		string
}

func main() {
	// read message from command line flags
	messagePtr := flag.String("message", "Test", "Message you want to send")
	fromNumberPtr := flag.String("from", "Hallo", "Number used to send messages")
	toNumberPtr := flag.String("to", "Hallo", "Number to send the message to")

	configPathPtr := flag.String("config", "config.json", "Absoloute location of config file")

	// parse all given flags
	flag.Parse()

	// read username, password, host and port from config file
	configuration := Configuration{}
	err := gonfig.GetConf(*configPathPtr, &configuration)

	// create a random accescode
	rand.Seed(time.Now().UnixNano())
	accessCode := RandStringBytes(10)

	fmt.Println("User:", configuration.Username)
	fmt.Println("Password:", configuration.Password)
	fmt.Println("AccessCode:", accessCode)
	address := configuration.Address + ":" + strconv.Itoa(configuration.Port)
	fmt.Println("Address:", address)

	fmt.Println("From:", *fromNumberPtr)
	fmt.Println("To:", *toNumberPtr)
	fmt.Println("Message:", *messagePtr)

	opt := &ucp.Options{
		Addr: address,
		User: configuration.Username,
		Password: configuration.Password,
		AccessCode: accessCode,
	}
	client := ucp.New(opt)
	if err := client.Connect(); err != nil {
		fmt.Println("Cant connect")
		os.Exit(1)
	}

	defer client.Close()

	ids, err := client.Send(*fromNumberPtr, *toNumberPtr, *messagePtr)

	fmt.Println("Error sending:", err)
	fmt.Println("Return from sending:", ids)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
