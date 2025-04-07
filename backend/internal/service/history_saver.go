package service

type HistorySaver interface {
	Save(hexPackage string) error
}

type historySaver struct{}

func NewHistorySaver() HistorySaver {
	return &historySaver{}
}

func (h *historySaver) Save(hexPackage string) error {
	return nil
}
