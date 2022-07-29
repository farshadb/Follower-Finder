# Follower-Finder
Random Twitter Follower Finder
This Little app helps you to find a follower in all of your twitter followers for draw

At first you should [Sign up ](https://developer.twitter.com/)for Twitter API service and answer a bunch of questions then Twitter team will activate your twitter developoer account in few days.

From Twitter [end-point API](https://developer.twitter.com/en/docs/twitter-api/migrate/twitter-api-endpoint-map) we use [Get followers/ids](https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids) and [Get users/show](https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/follow-search-get-users/api-reference/get-users-show) method.

You can find all tools in twitter [Tools and libraries](https://developer.twitter.com/en/docs/twitter-ads-api/tools-and-libraries) in this case used [Go-twitter](https://github.com/dghubble/go-twitter) library for Golang.

In twitter V2 API you should create a project to use API but I use API V1.1 and I create an stand-alone app [here](https://developer.twitter.com/en/portal/projects-and-apps) to get API credentionlas. And here we are done in server side tasks.
