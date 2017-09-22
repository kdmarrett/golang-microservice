package main

import (
	"./twitch"
	"fmt"
	"net/http"
)

type UserData struct {
	Data []DataInfo
}

type FollowData struct {
	_Total  int
	Follows []MetaChannelFollow
}

type MetaChannelFollow struct {
	Channel ChannelFollow
}

type ChannelFollow struct {
	language string
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
	fmt.Println("Booting the server...")

	// Configure a sample route
	http.HandleFunc("/", handler)

	// Run your server
	http.ListenAndServe(":8080", nil)
}

// handler function for route of HTTP server
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Received the following request:", r.URL.Query())
	login := r.URL.Query().Get("login")
	fmt.Fprintf(w, "Welcome, %s!\n\n", login)

	user := DataInfo{}
	twitch.GetUserByLogin(login, &user)

	info := r.URL.Query().Get("info")

	switch info {
	case "views":
		fmt.Fprintf(w, "Views: %d\n", user.View_count)
	case "followers":
		// requires a secondary query
		follows := FollowData{}
		twitch.GetFollows(user.Id, follows)
		fmt.Fprintf(w, "Followers: %d", follows._Total)
	case "game":
		fmt.Fprintf(w, "Game:", 0)
	case "language":
		fmt.Fprintf(w, "Language:", 0)
	case "streaming":
		fmt.Fprintf(w, "Streaming:", 0)
	case "display_name":
		fmt.Fprintf(w, "Display Name: %s\n", user.Display_name)
	case "bio":
		fmt.Fprintf(w, "Bio: %s\n", user.Bio)
	case "creation":
		fmt.Fprintf(w, "Account Creation Date: %s", user.Created_at)
	}

}
