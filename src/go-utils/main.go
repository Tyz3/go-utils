package main

import (
	"fmt"
	"go-utils/eventbroker"
)

// Custom Event

type PlayerJoinGameEvent struct {
	eventbroker.Event

	playerName string
}

func NewPlayerJoinGameEvent(playerName string) *PlayerJoinGameEvent {
	return &PlayerJoinGameEvent{
		playerName: playerName,
		Event:      eventbroker.NewEvent("PlayerJoinGameEvent"),
	}
}

func (e *PlayerJoinGameEvent) GetPlayerName() string {
	return e.playerName
}

// Custom Listener

func run1(e *PlayerJoinGameEvent) {
	fmt.Println("Handle event run1", e.GetPlayerName())
}

func main() {
	e := NewPlayerJoinGameEvent("Kronos")
	fmt.Println(e.GetEventName())

	eventbroker.GetEventManager().RegisterListener(e, run1, eventbroker.LOWEST, false)
	eventbroker.GetEventManager().RegisterListener(e, run1, eventbroker.LOWEST, false)

	e.Call()

	//Log.INFO.Printf("Hello World... ")
	//Log.Failed()
	//
	//file, _ := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND, 0666)
	//w := io.MultiWriter(os.Stdout, writers.WrapFreeColorsWriter(file))
	//Memo.SetWriter(w)
	//memo := Memo.New("TEST").Info("Haha %d", 1)
	//Memo.New("TEST").Info("Haha %d", 1).Debug("debug").Ok()
	//Memo.New("TEST").Info("Haha %d", 1).Debug("debug").Failed()
	//memo.Timestamp().Ok()
	//
	//Log.SetWriter(w)
	//Log.ERROR.Printf("Error1\n")
	//
	//list := make([]int, 0, 1)
	//
	//list2 := append(list, 1)
	//fmt.Println(list)
	//fmt.Println(list2)

}
