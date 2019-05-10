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
	ID          int64  // 全局唯一
	Email       string // email length >=5
	Password    string // password length >=6
	Role        Role
	Status      Status
	CreatedAt   int64
	UpdatedAt   int64
	AccessToken string // accessToken length >=32
}

type OperationResult struct {
	TransID   int64
	Code      int    // code = 200, success;  code=500, error
	Msg       string // return operation name when success,  return error message when error
	TimeStamp int64
}

type ReqRegister struct {
	TransID  int64
	Email    string // email length >=5
	Password string // password length >=6
}

type RespRegister struct {
	Result OperationResult
	Data   Account
}

// AccountOperation account operation interface define
type AccountOperation interface {
	Register(email, password string) (account *Account, result error)
	Exists(email string) (exists bool, result error)
	Login(email, password string) (token string, result error)
	Logout(token string) (result error)
	AuthToken(token string) (pass bool, result error)
	AuthUID(id int64) (pass bool, result error)
	Verify(token string) (pass bool, result error)
}
