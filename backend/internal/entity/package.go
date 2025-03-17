package entity

import (
	"strconv"

	"github.com/kuznetsovin/egts-protocol/libs/egts"
)

type egtsField struct {
	Description string `json:"description"`
	Value       string `json:"value"`
	Pos         [2]int `json:"pos"`
}

type DebugPackage struct {
	EGTSPackage     *egts.Package `json:"-"`
	ProtocolVersion egtsField     `json:"PRV"`
	SecurityKeyID   egtsField     `json:"SKID"`
}

func NewDebugPackage(egtsPackage *egts.Package) *DebugPackage {
	debugPkg := &DebugPackage{
		EGTSPackage: egtsPackage,
	}

	return debugPkg.mapFields()
}

func (p *DebugPackage) mapFields() *DebugPackage {
	p.ProtocolVersion = egtsField{
		Description: "Protocol version",
		Value:       strconv.Itoa(int(p.EGTSPackage.ProtocolVersion)),
		Pos:         [2]int{0, 1},
	}
	p.SecurityKeyID = egtsField{
		Description: "Security Key ID | Параметр SKID определяет идентификатор ключа, используемого при шифровании.",
		Value:       strconv.Itoa(int(p.EGTSPackage.SecurityKeyID)),
		Pos:         [2]int{1, 2},
	}

	return p
}
