package entity

type AuthenticationInfo struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	ClientCode string `json:"client_code"`
}
