package model

type Status int32

const (
	Status_None               Status = 0
	Status_WaitForEmailVerify Status = 1
	Status_Active             Status = 2
	Status_Suspend            Status = 4
	Status_Deleted            Status = 8
)

var Status_name = map[int32]string{
	0: "None",
	1: "WaitForEmailVerify",
	2: "Active",
	4: "Suspend",
	8: "Deleted",
}

var Status_value = map[string]int32{
	"None":               0,
	"WaitForEmailVerify": 1,
	"Active":             2,
	"Suspend":            4,
	"Deleted":            8,
}

type Role int32

const (
	Role_Guest  Role = 0
	Role_Member Role = 1
	Role_VIP    Role = 2
)

var Role_name = map[int32]string{
	0: "Guest",
	1: "Member",
	2: "VIP",
}

var Role_value = map[string]int32{
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
