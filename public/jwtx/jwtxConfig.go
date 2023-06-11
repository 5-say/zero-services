package jwtx

type JWTXConfig struct {
	AccessSecret    string // jwt 密钥
	AccessExpire    int64  // token 过期时间（秒）
	RefreshInterval int64  // 刷新时间间隔（秒）
	FaultTolerance  int64  // 并发容错时间（秒）
	CheckIP         bool   // 是否开启 IP 一致性校验，当前 IP 必须与登录时 IP 一致
}
