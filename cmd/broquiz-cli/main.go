package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Cyfell/BroQuiz/pkg/client"
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	if len(flag.Args()) < 1 {
		help()
		return 255
	}

	client := client.New("http://localhost:8080")

	switch flag.Arg(0) {
	case "answer":
		return answer(client)
	case "clear":
		return clear(client)
	case "events":
		return events(client)
	case "help":
		help()
	default:
		help()
		return 255
	}

	return 0
}

func help() {
	fmt.Println("Commands:")
	fmt.Println(" events: will listen for events")
	fmt.Println(" answer <team>: will ask for answer")
}

func events(c *client.Client) int {
	ch, err := c.Events()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when establishing events connection: %s\n", err)
		return 255
	}

	for {
		evt, ok := <-ch
		if ok {
			fmt.Println(evt)
		} else {
			break
		}
	}

	return 0
}

func answer(c *client.Client) int {
	if len(flag.Args()) < 2 {
		fmt.Fprintf(os.Stderr, "this command needs a team nb as argument\n")
		return 255
	}

	team, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "an error occured on team number: %s\n", err)
		return 255
	}

	ok, err := c.Answer(team)
	if ok {
		fmt.Println("You got the hand")
	} else {
		fmt.Println("You don't have the hand")
	}

	return 0
}

func clear(c *client.Client) int {
	if err := c.Clear(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occured when clearing: %s\n", err)
		return 255
	}

	return 0
}
