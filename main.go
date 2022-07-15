package main

import (
	"flag"
	"fmt"
	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"strings"
)

const (
	numberOfIterations = 20
	numberOfTweets     = 200
	envPrefix          = "TWITTER"
)

func deleteTweets(client *twitter.Client) error {
	includeRetweets := true
	userTimelineParams := &twitter.UserTimelineParams{Count: numberOfTweets, IncludeRetweets: &includeRetweets}
	statusUnretweetParams := &twitter.StatusUnretweetParams{}
	statusDestroyParams := &twitter.StatusDestroyParams{}
	totalDeleted := 0

	fmt.Println("===>>> Deleting tweets")

	for i := 1; i <= numberOfIterations; i++ {
		tweets, _, err := client.Timelines.UserTimeline(userTimelineParams)
		if len(tweets) == 0 {
			break
		} else if err != nil {
			return err
		}

		fmt.Printf("%d tweet(s) will be deleted in this iteration.\n\n", len(tweets))

		totalDeleted += len(tweets)

		for _, tweet := range tweets {
			if tweet.Retweeted {
				_, _, err := client.Statuses.Unretweet(tweet.ID, statusUnretweetParams)
				if err == nil {
					fmt.Printf("Deleted: %d - %s\n", tweet.ID, tweet.Text)
				} else {
					return err
				}
			} else if !tweet.Favorited {
				_, _, err := client.Statuses.Destroy(tweet.ID, statusDestroyParams)
				if err == nil {
					fmt.Printf("Deleted: %d - %s - %s\n", tweet.ID, tweet.CreatedAt, tweet.Text)
				} else {
					return err
				}
			}
		}
	}

	fmt.Printf("===>>> Deleted %d tweets\n\n", totalDeleted)
	return nil
}

func unfavoriteTweets(client *twitter.Client) error {
	favoriteListParams := &twitter.FavoriteListParams{Count: numberOfTweets}
	totalUnfavorited := 0

	fmt.Println("===>>> Unfavoriting tweets")

	for i := 1; i <= numberOfIterations; i++ {
		tweets, _, err := client.Favorites.List(favoriteListParams)
		if len(tweets) == 0 {
			break
		} else if err != nil {
			return err
		}

		fmt.Printf("%d tweet(s) will be unfavorited in this iteration.\n\n", len(tweets))

		totalUnfavorited += len(tweets)

		for _, tweet := range tweets {
			_, _, err := client.Favorites.Destroy(&twitter.FavoriteDestroyParams{ID: tweet.ID})
			if err == nil {
				fmt.Printf("Unfavorited: %d - %s - %s\n", tweet.ID, tweet.CreatedAt, tweet.Text)
			} else {
				return err
			}
		}
	}

	fmt.Printf("===>>> Unfavorited %d tweets\n\n", totalUnfavorited)
	return nil
}

func main() {
	credentials := struct {
		consumerKey       string
		consumerSecret    string
		accessToken       string
		accessTokenSecret string
	}{}

	flag.StringVar(&credentials.consumerKey, "consumer-key", "", "Twitter Consumer Key")
	flag.StringVar(&credentials.consumerSecret, "consumer-secret", "", "Twitter Consumer Secret")
	flag.StringVar(&credentials.accessToken, "access-token", "", "Twitter Access Token")
	flag.StringVar(&credentials.accessTokenSecret, "access-token-secret", "", "Twitter Access Token Secret")
	flag.Parse()

	if err := flagutil.SetFlagsFromEnv(flag.CommandLine, envPrefix); err != nil {
		log.Fatalf("Error loading flags: %v\n", err)
	}

	flags := map[string]string{
		"consumer-key":        credentials.consumerKey,
		"consumer-secret":     credentials.consumerSecret,
		"access-token":        credentials.accessToken,
		"access-token-secret": credentials.accessTokenSecret,
	}

	for name, value := range flags {
		if value == "" {
			envName := envPrefix + "_" + strings.ToUpper(strings.ReplaceAll(name, "-", "_"))
			flagName := "-" + name
			log.Fatalf("Missing required field \"%s\". Either set it through env var (%s) or flag (%s).\n", name, envName, flagName)
		}
	}

	client := twitter.NewClient(
		oauth1.
			NewConfig(credentials.consumerKey, credentials.consumerSecret).
			Client(
				oauth1.NoContext,
				oauth1.NewToken(credentials.accessToken, credentials.accessTokenSecret),
			),
	)

	if err := deleteTweets(client); err != nil {
		log.Fatalf("Error deleting tweets: %v\n", err)
	}

	if err := unfavoriteTweets(client); err != nil {
		log.Fatalf("Error unfavoriting tweets: %v\n", err)
	}
}
