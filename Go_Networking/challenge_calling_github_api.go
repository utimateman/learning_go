// challenge: calling github api
package main

import (

	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

// user is a github user information
type User struct {
	Name	 string `json:"name"`
	PublicRepos int `json:"public_repos"`
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	// http call
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode JSON
	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	user, err := userInfo("utimateman")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%+v\n", user)

}