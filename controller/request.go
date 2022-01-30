package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
)

type RequestType string

const (
	RequestJoin           RequestType = "join"
	RequestCreateGameRoom RequestType = "createGameRoom"
)

type Request struct {
	Type   RequestType `json:"type"`
	Id     string      `json:"id"`
	UserId string      `json:"userId,omitempty"`
}

func parseRequest(requestRaw string) *Request {
	if r, ok := util.JsonToObject(&Request{}, requestRaw); ok {
		return r.(*Request)
	}
	return nil
}
