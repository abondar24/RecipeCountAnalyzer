package data

type PostcodeBusy struct {
	Code  string `json:"postcode"`
	Count int32  `json:"delivery_count"`
}

type PostcodeTime struct {
	Code  string `json:"postcode"`
	From  string `json:"from"`
	To    string `json:"to"`
	Count int32  `json:"delivery_count"`
}
