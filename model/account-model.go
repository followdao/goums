package model

// Status account status like active / suspend / deleted
type Status int32

const (
	StatusNone               Status = 0
	StatusWaitForEmailVerify Status = 1
	StatusActive             Status = 2
	StatusSuspend            Status = 4
	StatusDeleted            Status = 8
)

// StatusName status name map
var StatusName = map[int32]string{
	0: "None",
	1: "WaitForEmailVerify",
	2: "Active",
	4: "Suspend",
	8: "Deleted",
}
// StatusValue  statue value map
var StatusValue = map[string]int32{
	"None":               0,
	"WaitForEmailVerify": 1,
	"Active":             2,
	"Suspend":            4,
	"Deleted":            8,
}

// Role role of account for access control
type Role int32

const (
	Role_Guest  Role = 0
	Role_Member Role = 1
	Role_VIP    Role = 2
)

// RoleName role name map
var RoleName = map[int32]string{
	0: "Guest",
	1: "Member",
	2: "VIP",
}
// RoleValue role value map
var RoleValue = map[string]int32{
	"Guest":  0,
	"Member": 1,
	"VIP":    2,
}

// Account  account  define
type Account struct {
	ID          int64  `json:"iD"`       // 全局唯一
	Email       string `json:"email"`    // email length >=5
	Password    string `json:"password"` // password length >=6
	Role        Role   `json:"role"`
	Status      Status `json:"status"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	AccessToken string `json:"accessToken"` // accessToken length >=32
}

// AccountOperation account operation interface define
type AccountOperation interface {
	Register(email, password string) (account *Account, err error)
	Exists(email string) (exists bool, err error)
	Login(email, password string) (token string, err error)
	Logout(token string) (err error)
	AuthToken(token string) (pass bool, err error)
	AuthUID(id int64) (pass bool, err error)
	Verify(token string) (pass bool, err error)
}
