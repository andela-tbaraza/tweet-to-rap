package main

import (
	"log"
	"net/http"
)

/*
Create a function that consumes a twiiter api
Try to get a pattern of converting tweet to rap
*/

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.twitter.com/1.1/search/tweets.json?q=@andela", nil)
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}
	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAALjFzQAAAAAA2imCmXZVfkZn4fuF%2FOK7nFqPgK4%3DQdDKESuRFFggk1SlpT2vOM0Fz953EXTMzXjshB8Y3aQnn6RaG0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}
	log.Printf("The response is %v", resp.Body)

}
