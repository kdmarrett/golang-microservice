package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var client_id string = "8seombbxiqt9h8bcee78xdyhdi83zp"
var twitch_API string = "https://api.twitch.tv/kraken/"

func init() {
	fmt.Println("Initialized Twitch API...")
}

// Fill struct with Stream info, `Stream`:null if not streaming
func GetStreamStatus(login string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", twitch_API+"streams/"+login, nil)
	req.Header.Set("Client-ID", client_id)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(contents, target)
	return err
}

// Fill struct with general channel information
func GetChannel(login string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", twitch_API+"channels/"+login, nil)
	req.Header.Set("Client-ID", client_id)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(contents, target)
	return err
}

// Fill struct with general user information
func GetUserByLogin(Id string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", twitch_API+"users/"+Id, nil)
	req.Header.Set("Client-ID", client_id)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(contents, target)
	return err
}
