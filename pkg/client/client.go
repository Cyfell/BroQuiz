package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Cyfell/BroQuiz/pkg/event"
	"github.com/gorilla/websocket"
)

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{
		url: url,
	}
}

func (c *Client) Answer(team int) (bool, error) {
	url := fmt.Sprintf("%s/answer/%d", c.url, team)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == 201, nil
}

func (c *Client) Clear() error {
	url := c.url + "/clear"
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Events() (chan event.Event, error) {
	ch := make(chan event.Event)

	url := "ws" + strings.TrimPrefix(c.url, "http") + "/events"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return ch, err
	}

	go transferEvents(ws, ch)

	return ch, nil
}

func transferEvents(ws *websocket.Conn, ch chan event.Event) {
	var evt event.Event
	for {
		if err := ws.ReadJSON(&evt); err != nil {
			fmt.Fprintf(os.Stderr, "error when receiving event: %s", err)
		}

		ch <- evt
	}
}
