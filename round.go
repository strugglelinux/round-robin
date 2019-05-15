package round

//ServiceWeight 权重结构
type ServiceWeight struct {
	Service      Server
	Weight       int
	EffectWeight int
}

// Services  服务
type Services struct {
	ServiceMap map[string]*ServiceWeight
}

//WeightRoundRobin 加权轮询调度算法 平滑
func (s *Services) WeightRoundRobin() string {
	var totalWeight int
	var name string
	for key, serviceWeight := range s.ServiceMap {
		s.ServiceMap[key].EffectWeight += serviceWeight.Weight
		totalWeight += serviceWeight.Weight
		if name == "" || serviceWeight.EffectWeight > s.ServiceMap[name].EffectWeight {
			name = key
		}
	}
	s.ServiceMap[name].EffectWeight -= totalWeight
	return name
}

//NewServices 初始化
func NewServices(sm map[string]*ServiceWeight) *Services {
	return &Services{ServiceMap: sm}
}
