package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
)

func main() {

	client := NewTwitterClient()
	ids, err := GetFollowersIDs(client)
	handleError(err)
	i := GetWinnerIndex(len(ids))
	user, err := GetWinnerUser(client, ids[i])
	handleError(err)

	fmt.Printf("Random Index: %v\nUserID: %v\nUser: %v\n\n%s(@%s) \nURL: https:twitter.com/@%s\n",
		i, ids[i], user.Name, user.Description, user.ScreenName, user.ScreenName)
}

func handleError(err error) { 
	if err != nil {
		log.Fatal(err)
	}
}

func GetWinnerUser(client *twitter.Client, id int64) (*twitter.User, error) {
	user, _, err := client.Users.Show(&twitter.UserShowParams{UserID: id})
	if 	err != nil {
		return nil, err
	}
	return user, nil
}

func GetWinnerIndex(count int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(count)
}

func NewTwitterClient() *twitter.Client {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

func GetFollowersIDs(client  *twitter.Client) ([]int64, error) {

	var cursor int64 = -1
	var ids[]int64

	for cursor != 0 {
		followers, _, err := client.Followers.IDs(&twitter.FollowerIDParams{
			// you can replace this with your own user ID or any other user ID
			ScreenName: "elonmusk",
			Cursor: cursor,
		})
		if err != nil {
			return nil, err
		}

		cursor = followers.NextCursor
		ids = append(ids, followers.IDs...)
	}

	return ids, nil
}