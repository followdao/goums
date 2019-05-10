package services

import (
	"golang.org/x/xerrors"
)

var ErrDuplicated = xerrors.New("account is existed, duplicated email address")
var ErrAccountMissing = xerrors.New("account is not existed")
