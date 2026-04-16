package service

type version struct {
	Clamav        string `json:"Clamav"`
	Signature     string `json:"Signature"`
	SignatureDate string `json:"Signature_date"`
}
