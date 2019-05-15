package round

import "strconv"

// Server 服务地址信息
type Server struct {
	// IP 地址
	Host string
	// 端口号
	Port int
	//服务 超时时间
	ExpireTime int64
}

//GetHost 获取Host
func (s *Server) GetHost() string {
	return s.Host
}

//GetPost 获取post
func (s *Server) GetPost() int {
	return s.Port
}

func (s *Server) String() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

//KeepLive 判断是否存活
func (s *Server) KeepLive() bool {
	return true
}
