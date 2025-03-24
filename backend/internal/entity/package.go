package entity

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kuznetsovin/egts-protocol/libs/egts"
)

type egtsField struct {
	Description string `json:"description"`
	Value       string `json:"value"`
	BytesSize   int    `json:"bytes_size"`
}

type DebugPackage struct {
	EGTSPackage      *egts.Package `json:"-"`
	ProtocolVersion  egtsField     `json:"PRV"`
	SecurityKeyID    egtsField     `json:"SKID"`
	HeaderFlags      egtsField     `json:"PRF RTE ENA CMP PR"`
	HeaderLength     egtsField     `json:"HL"`
	FrameDataLength  egtsField     `json:"FDL"`
	PacketIdentifier egtsField     `json:"PID"`
	PacketType       egtsField     `json:"PT"`

	HeaderCheckSum            egtsField `json:"HCS"`
	ServicesFrameData         egtsField `json:"SFRD"`
	ServicesFrameDataCheckSum egtsField `json:"SFRCS"`
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
		BytesSize:   1,
	}
	p.SecurityKeyID = egtsField{
		Description: "Security Key ID | Параметр SKID определяет идентификатор ключа, используемого при шифровании.",
		Value:       strconv.Itoa(int(p.EGTSPackage.SecurityKeyID)),
		BytesSize:   1,
	}
	p.HeaderFlags = egtsField{
		Description: `Параметр PRV содержит значение 0x01. Значение данного параметра инкрементируется каждый раз при внесении изменений в структуру заголовка.
Поле PR (Priority) определяет приоритет маршрутизации данного пакета`,
		Value: fmt.Sprintf(
			`PRF = %s RTE = %s ENA = %s CMP = %s PR = %s`,
			p.EGTSPackage.Prefix,
			p.EGTSPackage.Route,
			p.EGTSPackage.EncryptionAlg,
			p.EGTSPackage.Compression,
			p.EGTSPackage.Priority,
		),
		BytesSize: 1,
	}
	p.HeaderLength = egtsField{
		Description: `Поле HL — длина заголовка транспортного уровня в байтах с учетом байта контрольной суммы (поля HCS).`,
		Value:       strconv.Itoa(int(p.EGTSPackage.HeaderLength)),
		BytesSize:   1,
	}

	p.FrameDataLength = egtsField{
		Description: `Поле FDL определяет размер в байтах поля данных SFRD, содержащего информацию протокола уровня поддержки услуг.`,
		Value:       strconv.Itoa(int(p.EGTSPackage.FrameDataLength)),
		BytesSize:   2,
	}

	p.PacketIdentifier = egtsField{
		Description: `Поле PID содержит номер пакета транспортного уровня, увеличивающийся на 1 при отправке каждого
нового пакета на стороне отправителя. Значения в данном поле изменяются по правилам циклического счетчика в
диапазоне от 0 до 65535, т. е. при достижении значения 65535 следующее значение 0.`,
		Value:     strconv.Itoa(int(p.EGTSPackage.PacketIdentifier)),
		BytesSize: 2,
	}

	p.PacketType = egtsField{
		Description: `Поле РТ — тип пакета транспортного уровня. Поле РТ может принимать следующие значения:
- 0 — EGTS_PT_RESPONSE (подтверждение на пакет транспортного уровня);
- 1 — EGTS_PT_APPDATA (пакет, содержащий данные протокола уровня поддержки услуг);
- 2 — EGTS_PT_SIGNED_APPDATA (пакет, содержащий данные протокола уровня поддержки услуг с цифровой подписью).`,
		Value:     strconv.Itoa(int(p.EGTSPackage.PacketType)),
		BytesSize: 1,
	}

	p.HeaderCheckSum = egtsField{
		Description: `Поле HCS — контрольная сумма заголовка Транспортного уровня (начиная с поля «PRV» до поля
«HCS», не включая поле «HCS»). Для подсчета значения поля HCS ко всем байтам указанной последовательности
применяется алгоритм CRC-8.`,
		Value:     strconv.Itoa(int(p.EGTSPackage.HeaderCheckSum)),
		BytesSize: 1,
	}

	serviceFrameData, _ := json.Marshal(p.EGTSPackage.ServicesFrameData)

	p.ServicesFrameData = egtsField{
		Description: `Поле SFRD — структура данных, зависящая от типа пакета и содержащая информацию протокола
уровня поддержки услуг.`,
		Value:     string(serviceFrameData),
		BytesSize: int(p.EGTSPackage.FrameDataLength),
	}

	p.ServicesFrameDataCheckSum = egtsField{
		Description: `Поле SFRCS — контрольная сумма поля протокола уровня поддержки услуг. Для подсчета контрольной
суммы по данным из поля SFRD используется алгоритм CRC-16. Данное поле присутствует только в том случае,
если есть поле SFRD.`,
		Value:     strconv.Itoa(int(p.EGTSPackage.ServicesFrameDataCheckSum)),
		BytesSize: 2,
	}

	return p
}
