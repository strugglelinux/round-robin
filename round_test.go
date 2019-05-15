package round

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	server1 := Server{Host: "192.168.1.0", Port: 8080}
	server2 := Server{Host: "192.168.1.1", Port: 8080}
	serviceWeight1 := ServiceWeight{Service: server1, Weight: 30}
	serviceWeight2 := ServiceWeight{Service: server2, Weight: 70}

	serviceMap := map[string]*ServiceWeight{"server1": &serviceWeight1, "server2": &serviceWeight2}
	services := NewServices(serviceMap)
	for a := 0; a < 10; a++ {
		serviceKey := services.WeightRoundRobin()
		if serviceKey == "" {
			t.Errorf("键a对应值未找到 %v", serviceKey)
		} else {
			service := serviceMap[serviceKey]
			fmt.Printf("服务 %s,Host:%s,Port:%d,Weight:%d,EffectWeight:%d\n", serviceKey, service.Service.Host, service.Service.Port, service.Weight, service.EffectWeight)
		}
	}

}
