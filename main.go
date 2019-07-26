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
	"log/syslog"
)

type Configuration struct {
	Port			int
	Address			string
	Username		string
	Password		string
}

func main() {
	logwriter, err := syslog.New(syslog.LOG_ERR, "ucp-cli")
	if err != nil{
		fmt.Println(err)
	}


	// read message from command line flags
	messagePtr := flag.String("message", "Test", "Message you want to send")
	fromNumberPtr := flag.String("from", "Hallo", "Number used to send messages")
	toNumberPtr := flag.String("to", "Hallo", "Number to send the message to")

	configPathPtr := flag.String("config", "config.json", "Absoloute location of config file")

	// parse all given flags
	flag.Parse()

	// read username, password, host and port from config file
	configuration := Configuration{}
	err = gonfig.GetConf(*configPathPtr, &configuration)
	if err != nil {
		fmt.Println(err)
	}

	// create a random accescode
	rand.Seed(time.Now().UnixNano())
	accessCode := RandStringBytes(10)

	address := configuration.Address + ":" + strconv.Itoa(configuration.Port)

	message := "Sending SM from " + *fromNumberPtr + " to " + *toNumberPtr
	logwriter.Info(message)
	message = "Message is: " + *messagePtr
	logwriter.Info(message)

	opt := &ucp.Options{
		Addr: address,
		User: configuration.Username,
		Password: configuration.Password,
		AccessCode: accessCode,
	}
	client := ucp.New(opt)
	if err := client.Connect(); err != nil {
		logwriter.Err("Cant connect to SMSC")
		os.Exit(1)
	}

	defer client.Close()

	ids, err := client.Send(*fromNumberPtr, *toNumberPtr, *messagePtr)

	message = "Returning from sending:" + ids[1]
	logwriter.Info(message)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
