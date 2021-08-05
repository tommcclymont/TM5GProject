package models

type BackupAmfInfo struct {
	BackupAmf string
	GuamiList []Guami
}

var BackupAmfInfoData = BackupAmfInfo{
	BackupAmf: "testbackupamf",
	GuamiList: []Guami{
		GuamiData,
	},
}
