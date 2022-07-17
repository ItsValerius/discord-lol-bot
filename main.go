package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Variables used for command line parameters
var (
	Token     string
	API_KEY   string
	ChannelID string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&API_KEY, "k", "", "API KEY")
	flag.StringVar(&ChannelID, "c", "", "ChannelID")

	flag.Parse()
}

func main() {
	dg := Start()

	var users Users

	getUsers(&users)
	for {
		for i := 0; i < len(users.Users); i++ {
			<-time.After(time.Second * 5)
			status := isInGame(users.Users[i])
			if status {
				dg.ChannelMessageSend(ChannelID, users.Users[i].Name+" is in game!")
				continue
			}
			dg.ChannelMessageSend(ChannelID, users.Users[i].Name+" is not in game!")

		}

		<-time.After(time.Second * 5)
	}
}

func getUsers(users *Users) {
	// Open our jsonFile
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)
}
