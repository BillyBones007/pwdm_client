package tablesgen

import (
	"fmt"

	"github.com/BillyBones007/pwdm_client/internal/storage/models"
	"github.com/alexeyco/simpletable"
)

// NewInfoTable - returns a table with data of the specified type as a string.
func NewInfoTable(dataType int32, data map[int]models.InfoModel) string {
	dataSize := len(data)
	if dataSize < 1 {
		return ""
	}
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "TITLE"},
			{Align: simpletable.AlignCenter, Text: "COMMENT"},
			{Align: simpletable.AlignCenter, Text: "TAG"},
		},
	}

	for k, v := range data {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", k)},
			{Align: simpletable.AlignLeft, Text: v.Title},
			{Align: simpletable.AlignLeft, Text: v.Comment},
			{Align: simpletable.AlignLeft, Text: v.Tag},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	return table.String()
}

// NewLogPwdTable - returns a table with login/password data as a string.
func NewLogPwdTable(data models.LogPwdModel) string {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "LOGIN"},
			{Align: simpletable.AlignCenter, Text: "PASSWORD"},
			{Align: simpletable.AlignCenter, Text: "TITLE"},
			{Align: simpletable.AlignCenter, Text: "COMMENT"},
			{Align: simpletable.AlignCenter, Text: "TAG"},
		},
	}
	row := []*simpletable.Cell{
		{Align: simpletable.AlignRight, Text: data.Login},
		{Align: simpletable.AlignRight, Text: data.Password},
		{Align: simpletable.AlignRight, Text: data.TechData.Title},
		{Align: simpletable.AlignRight, Text: data.TechData.Comment},
		{Align: simpletable.AlignRight, Text: data.TechData.Tag},
	}
	table.Body.Cells = append(table.Body.Cells, row)

	table.SetStyle(simpletable.StyleCompactLite)
	return table.String()
}

// NewCardTable - returns a table with bank card data as a string.
func NewCardTable(data models.CardModel) string {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "NUMBER"},
			{Align: simpletable.AlignCenter, Text: "CVC"},
			{Align: simpletable.AlignCenter, Text: "EXP DATE"},
			{Align: simpletable.AlignCenter, Text: "FIRST NAME"},
			{Align: simpletable.AlignCenter, Text: "LAST NAME"},
			{Align: simpletable.AlignCenter, Text: "TITLE"},
			{Align: simpletable.AlignCenter, Text: "COMMENT"},
			{Align: simpletable.AlignCenter, Text: "TAG"},
		},
	}
	row := []*simpletable.Cell{
		{Align: simpletable.AlignRight, Text: data.Num},
		{Align: simpletable.AlignRight, Text: data.CVC},
		{Align: simpletable.AlignRight, Text: data.Date},
		{Align: simpletable.AlignRight, Text: data.FirstName},
		{Align: simpletable.AlignRight, Text: data.LastName},
		{Align: simpletable.AlignRight, Text: data.TechData.Title},
		{Align: simpletable.AlignRight, Text: data.TechData.Comment},
		{Align: simpletable.AlignRight, Text: data.TechData.Tag},
	}
	table.Body.Cells = append(table.Body.Cells, row)

	table.SetStyle(simpletable.StyleCompactLite)
	return table.String()
}
