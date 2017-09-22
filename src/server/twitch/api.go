package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
)

var client_id string = "8seombbxiqt9h8bcee78xdyhdi83zp"
var twitch_API string = "https://api.twitch.tv/kraken/"

func init() {
	fmt.Println("Initializing Twitch API...")
}

func GetFollows(Id string, target interface{}) error {
	client := &http.Client{}
	//req, err := http.NewRequest("GET", twitch_API+"channels/"+login+"/follows", nil)
	//Get list up to 25 chans followed by user
	req, err := http.NewRequest("GET", twitch_API+"users/"+Id+"/follows/channels", nil)
	req.Header.Set("Client-ID", client_id)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	fmt.Println("GetFollows response object:")
	fmt.Println(string(contents))
	json.Unmarshal(contents, target)
	return err
}

func GetUserByLogin(Id string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", twitch_API+"users/"+Id, nil)
	req.Header.Set("Client-ID", client_id)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		fmt.Println("%s", err)
		return err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		fmt.Println("%s", err)
		return err
	}
	fmt.Println(string(contents))
	json.Unmarshal(contents, target)
	return err
}
