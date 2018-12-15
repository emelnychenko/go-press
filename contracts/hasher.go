package contracts

import "github.com/emelnychenko/go-press/common"

type (
	Hasher interface {
		Make(password string) (string, common.Error)
		Check(hashedPassword, password string) common.Error
	}
)
