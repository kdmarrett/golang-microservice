# Golang Microservice

A microservice API to Twitch, the video game streaming platform. 
This service handles requests relating to a single user's basic information
and channel statistics. 


## Installation

#### Prerequisites

- Installation of golang 1.9 or higher


`$ git clone https://github.com/kdmarrett/golang-microservice`

From the project directory run:

`$ go build`

`$ ./server`
```
Initialized Twitch API...
Server booted successfully, listening...
```

Once the server has confirmed proper boot up, navigate to https://localhost:8080/

## Usage

This Server utilizes Twitch's v5 'kraken' API. Note that although this API versioning will
be deprecated in 2018, the newest version ('helix') is not fully implemented yet
at the time of authoring this project.

The request response cycle between a user's query and the response through the golang
server is mediated through the standard query parameters of the URL.
 All requests query some field of a specific user, this user is specified by
setting the query parameter `?login=<login-name>` and the field in question
is set by the parameter `?info=<info-type`. For example, to find the short
bio description of the Twitch account user 'kdmarrett' we would use the following query:

`http://localhost:8080/?login=kdmarrett&info=bio`

which would print the following to the screen:

```
User: kdmarrett

Bio: kdmarrett's short test bio
```

This is done in the same way for all of the servers supported requests to the following fields.

#### User channel # of views

`http://localhost:8080/?login=<login-name>&info=views`

#### User channel # of followers

`http://localhost:8080/?login=<login-name>&info=followers`

#### User channel game name

`http://localhost:8080/?login=<login-name>&info=game`

#### User channel language

`http://localhost:8080/?login=<login-name>&info=language`

#### User currently streaming status

`http://localhost:8080/?login=<login-name>&info=streaming`

#### User display name

`http://localhost:8080/?login=<login-name>&info=display_name`

#### User bio

`http://localhost:8080/?login=<login-name>&info=bio`

#### User account creation date

`http://localhost:8080/?login=<login-name>&info=creation`

## Resources
- [Twitch API v5 documentation](https://dev.twitch.tv/docs/v5/reference)
