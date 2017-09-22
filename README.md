# Go HTTP Server

HTTP Server as an API to Twitch, to fetch stream data
and streamer's channel stats.


#### Prerequisites

- Installation of golang 1.9 or higher

## Installation

`$ git clone https://github.com/kdmarrett/golang-microservice`

From the project directory run:

`$ go build`
`$ ./server`

When the server has confirmed boot up, navigate to https://localhost:8080/

## Usage

This Server utilizes Twitch's v5 'kraken' API. Warning: this API version will
be deprecated in 2018 however the newest version 'helix' is not fully implemented yet
as of authoring this server.


