package ui

import "time"

type startMsg struct {
	signIn bool
	reg    bool
}

type userMsg struct {
	login    string
	password string
}

type tickMsg time.Time
