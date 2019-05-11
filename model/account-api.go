package model


// Account  account  define
type AccountProfile struct {
	ID        int64  `json:"iD"`    // 全局唯一
	Email     string `json:"email"` // email length >=5
	Role      Role   `json:"role"`
	Status    Status `json:"status"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type Result struct {
	TransID   int64  `json:"transID"`
	Code      int    `json:"code"` // code = 200, success;  code=500, error`
	Msg       string `json:"msg"`  // return operation name when success,  return error message when error
	TimeStamp int64  `json:"timeStamp"`
}

type RegisterReq struct {
	TransID  int64  `json:"transID"`
	Email    string `json:"email"`    // email length >=5
	Password string `json:"password"` // password length >=6
}
type RegisterResp struct {
	Result Result  `json:"result"`
	Data   AccountProfile `json:"data"`
}

type AccountRequest struct {
	Email    string `json:"email"`    // email length >=5
	Password string `json:"password"` // password length >=6
}




