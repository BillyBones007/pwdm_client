package ui

import (
	"fmt"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	"github.com/charmbracelet/bubbles/table"
)

type TableModel struct {
	table table.Model
	// columns []table.Column
	// rows    []table.Row
}

func InitialTableModel(tabID int, stor *grpcclient.ClientGRPC) TableModel {
	columns := []table.Column{
		{Title: numRec, Width: 9},
		{Title: titleRec, Width: 15},
		{Title: commentRec, Width: 20},
	}

	var rows []table.Row
	switch tabID {
	case LogPwdTab:
		rows = setRowsTable(stor.Storage.LogPwdData)
	case CardTab:
		rows = setRowsTable(stor.Storage.CardData)
	case TextTab:
		rows = setRowsTable(stor.Storage.TextData)
	case BinaryTab:
		rows = setRowsTable(stor.Storage.BinaryData)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(1),
		table.WithWidth(100),
	)
	t.SetStyles(defaultTableStyle)
	tm := TableModel{table: t}
	return tm
}

func (t TableModel) View() string {
	return baseTableStyle.Render(t.table.View()) + "\n"
}

// helper function
func setRowsTable(storage map[int]models.InfoModel) []table.Row {
	if len(storage) == 0 {
		rows := []table.Row{{" ", " ", " "}}
		return rows
	}
	rows := make([]table.Row, len(storage))
	for i := 1; i <= len(storage); i++ {
		numrec := fmt.Sprint(i)
		rec := storage[i]
		rows = append(rows, table.Row{numrec, rec.Title, rec.Comment})
	}
	return rows
}
