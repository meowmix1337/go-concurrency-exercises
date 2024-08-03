//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			// close the tweets channel when we are done producing
			close(tweets)
			return
		}

		// push tweet to channel for consumer to receive
		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet) {
	// we can loop over the tweets being received via channel
	// this basically is listening for tweets
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// tweets channel to push tweets onto
	tweets := make(chan *Tweet)

	// Producer, make concurrent
	go producer(stream, tweets)

	// Consumer
	// pass tweet channel so when we produce a tweet, the consumer will receive the tweet
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
