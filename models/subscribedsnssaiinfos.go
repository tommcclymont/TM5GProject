package models

import (
	"encoding/json"
)

var SnssaiDataStr, err = json.Marshal(SnssaiData)

var SubscribedSnssaiInfosMap = map[string]DnnInfo{
	string(SnssaiDataStr): DNNInfo1,
}
