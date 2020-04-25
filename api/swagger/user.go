package swagger


type UserInfo struct {
	// 用户ID
	UserID int32 `json:"user_id"`
	// 用户名
	Username string `json:"username"`
	// 头像
	Avatar string `json:"avatar"`
}

type AuthToken struct {
	// JWT 认证
	Token string `json:"token"`
}


// 返回的用户信息
// swagger:response userInfoResp
type swaggerUserInfoResp struct {
	Body struct{
		// err code
		Code int32 `json:"code"`
		Data UserInfo `json:"data"`
	}
}

// 用户登陆
// swagger:response userLoginResp
type swaggerUserLogin struct {
	Body struct{
		Code int32 `json:"code"`
		Data AuthToken `json:"data"`
	}
}

