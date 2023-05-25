package uishell

import "github.com/c-bata/go-prompt"

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
		return prompt.FilterHasPrefix(showSubcommands, second, true)

	case "save":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(saveSubcommands, second, true)
		}
		if len(args) == 3 {
			third := args[2]
			switch second {
			case "text":
				return prompt.FilterHasPrefix(saveTextSubcommands, third, true)
			}
		}

	case "get":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(getSubcommands, second, true)
		}
		if len(args) == 3 {
			third := args[2]
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

	case "delete":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(deleteSubcommands, second, true)
		}
		if len(args) == 3 {
			third := args[2]
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
