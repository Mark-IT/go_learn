package main

import "go_learn/day10/chat/chat_server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}
