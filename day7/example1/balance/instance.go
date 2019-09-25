package balance

import "strconv"

type Instance struct {
	host string
	port int
}

func NewInstance(host string,port int) *Instance{	// 因为host,port是小写，外部无法直接使用，故自定义构造方法
	return &Instance{
		host:host,
		port:port,
	}
}

func (p *Instance) GetHost() string{
	return p.host
}

func (p *Instance) GetPort() int {
	return p.port
}

func (p *Instance) String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}
