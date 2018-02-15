package main

import (
	"log"
	"net/http"
)

func Start(res http.ResponseWriter, req *http.Request) {
	log.Print("START REQUEST")

	data, err := NewStartRequest(req)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(data)

	respond(res, StartResponse{
		Taunt:          "battlesnake-go!",
		Color:          "#75CEDD",
		Name:           "battlesnake-go",
		HeadType:       HEAD_PIXEL,
		TailType:       TAIL_ROUND_BUM,
		SecondaryColor: "#F7D3A2",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	log.Printf("MOVE REQUEST")

	data, err := NewMoveRequest(req)

	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	respond(res, MoveResponse{
		Move: handleMove(data),
	})
}
