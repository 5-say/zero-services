package jwtx

type JWTXConfig struct {
	CheckIP         bool  // 是否开启 IP 一致性校验，当前 IP 必须与登录时 IP 一致
	RefreshInterval int64 // 刷新时间间隔（秒）
	FaultTolerance  int64 // 并发容错时间（秒）
}
