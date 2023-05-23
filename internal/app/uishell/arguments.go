package uishell

import "github.com/c-bata/go-prompt"

var authUserCommands = []prompt.Suggest{
	{Text: "logout", Description: "Log out of your account"},
	{Text: "exit", Description: "Exit this program"},
	{Text: "show", Description: "Show user saved data"},
	{Text: "save", Description: "Save user data on the server"},
	{Text: "get", Description: "Get user data from server"},
	{Text: "delete", Description: "Delete user data from server"},
}

var guestUserCommands = []prompt.Suggest{
	{Text: "login", Description: "Log in to your account"},
	{Text: "register", Description: "Add a new user"},
	{Text: "exit", Description: "Exit this program"},
	{Text: "help", Description: "Help information"},
}

func (c *Completer) argumentsCompleter(args []string) []prompt.Suggest {
	if !c.client.AuthFlag {
		if len(args) <= 1 {
			return prompt.FilterHasPrefix(guestUserCommands, args[0], true)
		}
	}
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(authUserCommands, args[0], true)
	}
	first := args[0]
	switch first {
	case "show":
		second := args[1]
		subcommands := []prompt.Suggest{
			{Text: "lp", Description: "Show saved login/password data"},
			{Text: "card", Description: "Show saved bank cards data"},
			{Text: "text", Description: "Show saved text data"},
			{Text: "binary", Description: "Show saved binary data"},
		}
		return prompt.FilterHasPrefix(subcommands, second, true)

	case "delete":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "lp", Description: "Select login/password data to delete"},
				{Text: "card", Description: "Select bank cards data to delete"},
				{Text: "text", Description: "Select text data to delete"},
				{Text: "binary", Description: "Select binary data to delete"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}
		third := args[2]
		if len(args) == 3 {
			switch second {
			case "lp":
				return prompt.FilterContains(getLogPwdInfo(c.client), third, true)
			case "card":
				return prompt.FilterContains(getCardInfo(c.client), third, true)
			case "text":
				return prompt.FilterContains(getTextInfo(c.client), third, true)
			case "binary":
				return prompt.FilterContains(getBinaryInfo(c.client), third, true)
			}
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
