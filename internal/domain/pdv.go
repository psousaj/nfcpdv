package domain

import (
	"github.com/uptrace/bun"
)

type PDV struct {
	bun.BaseModel `bun:"table:pdvs"`

	ID               int64  `bun:",pk,autoincrement" json:"id"`
	Ip               string `bun:"notnull,unique" json:"ip"`
	Loja             string `bun:"notnull" json:"loja"`
	Numero           string `bun:"notnull" json:"numero"`
	SshUser          string `bun:"notnull, default:root" json:"ssh_user"`
	SshPort          int16  `bun:"notnull,default:22" json:"ssh_port"`
	SshPass          string `bun:"notnull,default:consinco" json:"_"`
	DriverMFEPresent bool   `bun:"default:true" json:"driver_mfe_present"`
	LastScanAt       string `bun:"nullzero" json:"last_scan_at"`
}
