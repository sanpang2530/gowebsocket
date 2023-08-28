package conn

// Login 用户登录
type Login struct {
	AppId  uint32
	UserId string
	Client *Client
}

// GetKey 获取 key
func (l *Login) GetKey() (key string) {
	key = GetUserKey(l.AppId, l.UserId)
	return
}
