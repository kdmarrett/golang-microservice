package main

import (
	"./twitch"
	"fmt"
	"log"
	"net/http"
)

type StreamData struct {
	Stream StreamInfo
}

type StreamInfo struct {
	_Id  int
	Game string
}

type ChannelInfo struct {
	Id           string
	Created_at   string
	Login        string
	Display_name string
	Type         string
	Language     string
	Game         string
	Bio          string
	Views        int
	Followers    int
	Status       string
}

type DataInfo struct {
	Id               string
	Created_at       string
	Login            string
	Display_name     string
	Type             string
	Broadcaster_type string
	Bio              string
	View_count       int
}

func main() {
	fmt.Println("Server booted successfully, listening...")

	// Configure a sample route
	http.HandleFunc("/", handler)

	// Run your server
	http.ListenAndServe(":8080", nil)
}

// handler function for route of HTTP server
func handler(w http.ResponseWriter, r *http.Request) {
	// get user inputted query information
	login := r.URL.Query().Get("login")
	info := r.URL.Query().Get("info")
	fmt.Fprintf(w, "User: %s\n", login)

	//get standard user data json
	user := DataInfo{}
	err := twitch.GetUserByLogin(login, &user)
	if err != nil {
		log.Fatal(err)
	}

	//get standard channel data json
	channel := ChannelInfo{}
	err = twitch.GetChannel(login, &channel)
	if err != nil {
		log.Fatal(err)
	}

	switch info {
	case "views":
		fmt.Fprintf(w, "Views: %d\n", channel.Views)
	case "followers":
		fmt.Fprintf(w, "Followers: %d\n", channel.Followers)
	case "game":
		fmt.Fprintf(w, "Game: %s\n", channel.Game)
	case "language":
		fmt.Fprintf(w, "Language: %s\n", channel.Language)
	case "streaming":
		stream := StreamData{}
		twitch.GetStreamStatus(login, &stream)
		status := "Currently streaming"
		if stream.Stream.Game == "" {
			status = "Not currently streaming"
		}
		fmt.Fprintf(w, "Streaming: %s\n", status)
	case "display_name":
		fmt.Fprintf(w, "Display Name: %s\n", user.Display_name)
	case "bio":
		fmt.Fprintf(w, "Bio: %s\n", user.Bio)
	case "creation":
		fmt.Fprintf(w, "Account Creation Date: %s", user.Created_at)
	}

}
