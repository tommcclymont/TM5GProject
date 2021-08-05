package models

type KeyAmf struct {
	KeyType string
	KeyVal  string
}

var KeyAmfData = KeyAmf{
	KeyType: "KAMF",
	KeyVal:  "C",
}
