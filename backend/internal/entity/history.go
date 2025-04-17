package entity

type HistoryRecord struct {
	ID  int    `json:"-"`
	Hex string `json:"hex"`
}
