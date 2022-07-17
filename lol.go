package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AddUser(username string, tagline string) {

	url := "https://europe.api.riotgames.com/riot/account/v1/accounts/by-riot-id/" + username + "/" + tagline
	bearer := API_KEY
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// add authorization header to the req
	req.Header.Add("X-Riot-Token", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)
}

func isInGame(user User) bool {
	fmt.Println("checking status for: ", user.Name)
	url := "https://euw1.api.riotgames.com/lol/spectator/v4/active-games/by-summoner/" + user.Id
	bearer := API_KEY
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// add authorization header to the req
	req.Header.Add("X-Riot-Token", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println(user.Name, " is in game")
		return true
	} else {
		fmt.Println(user.Name, " is not in game")
		return false
	}

}
