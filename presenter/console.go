package presenter

import (
	"fmt"

	"github.com/kihyun1998/go-aws/service"
)

type InstancePresenter interface {
	Present(instances []service.EC2Instance)
}

type ConsolePresenter struct{}

func NewConsolePresenter() InstancePresenter {
	return &ConsolePresenter{}
}

func (p *ConsolePresenter) Present(instances []service.EC2Instance) {
	fmt.Printf("총 인스턴스 수: %d\n", len(instances))
	for _, instance := range instances {
		fmt.Println("----------------------------------------")
		fmt.Printf("인스턴스 ID: %s\n", instance.InstanceID)
		fmt.Printf("인스턴스 타입: %s\n", instance.InstanceType)
		fmt.Printf("상태: %s\n", instance.State)

		fmt.Println("태그:")
		for key, value := range instance.Tags {
			fmt.Printf("  %s: %s\n", key, value)
		}

		if instance.PublicIP != "" {
			fmt.Printf("퍼블릭 IP: %s\n", instance.PublicIP)
		}
		if instance.PrivateIP != "" {
			fmt.Printf("프라이빗 IP: %s\n", instance.PrivateIP)
		}
	}
}
