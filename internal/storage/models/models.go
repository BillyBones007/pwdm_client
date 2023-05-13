package models

// InfoModel - contains the information by user data.
type InfoModel struct {
	Id      int32  // record id
	Type    int32  // type data
	Title   string // record title
	Tag     string // record tag
	Comment string // record comment
}

// UserModel - model for authentication/registration users.
type UserModel struct {
	Login    string
	Password string
}

// ReqTechDataModel - model with general information for requests.
type TechDataModel struct {
	Title   string
	Tag     string
	Comment string
	Type    int32
}

// LogPwdModel - model login/password pair.
type LogPwdModel struct {
	Login    string
	Password string
	TechData TechDataModel
}

// CardModel - model card data.
type CardModel struct {
	Num       string // card number
	Date      string // validity period
	CVC       string // cvc card code
	FirstName string
	LastName  string
	TechData  TechDataModel
}

// TextDataModel - model for text data.
type TextDataModel struct {
	Data     string // some text data
	TechData TechDataModel
}

// BinaryDataModel - model for binary data.
type BinaryDataModel struct {
	Data     []byte // some binary data
	TechData TechDataModel
}
