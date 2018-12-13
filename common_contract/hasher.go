package common_contract

import "../common"

type (
	Hasher interface {
		Make(password string) (string, common.Error)
		Check(hashedPassword, password string) common.Error
	}
)
