package entity

type HistoryRecord struct {
	ID             int
	EgtsPackageHex string
}

func (hr *HistoryRecord) TableName() string {
	return "egts_history_records"
}
