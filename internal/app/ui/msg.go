package ui

type startMsg struct {
	signIn bool
	reg    bool
}

type userMsg struct {
	login    string
	password string
}

type errorMsg struct {
	err error
}
