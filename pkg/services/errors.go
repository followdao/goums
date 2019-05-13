package services

import (
	"golang.org/x/xerrors"
)

// ErrDuplicated  error define register email was duplicated
var ErrDuplicated = xerrors.New("account is existed, duplicated email address")

// ErrAccountMissing error define account profile can't found ( missing )
var ErrAccountMissing = xerrors.New("account is not existed")
