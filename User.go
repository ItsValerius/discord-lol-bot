package main

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Tagline  string `json:"tagline"`
	Id       string `json:"id"`
	Puuid    string `json:"puuid"`
}
