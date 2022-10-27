package utils

import (
	"context"
	"encoding/json"
	"log"
)

type Userinfo struct {
	Id       string
	Level    string
	LptkId   string
	Name     string
	Username string
	Email    string
}

func GetUserinfo(ctx context.Context) Userinfo {
	userinfo := Userinfo{}
	dataJson := ctx.Value("userinfo").(string)
	_ = json.Unmarshal([]byte(dataJson), &userinfo)
	log.Printf("Unmarshaled: %v", userinfo)

	return userinfo
}
