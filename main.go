package main

import (
	"context"
	"log"

	"github.com/kihyun1998/go-aws/config"
	"github.com/kihyun1998/go-aws/presenter"
	"github.com/kihyun1998/go-aws/service"
)

func main() {
	configLoader := config.NewEnvConfigLoader()
	cfg, err := configLoader.LoadConfig()
	if err != nil {
		log.Fatal("설정을 불러오는 중 오류 발생:", err)
	}

	ec2Service, err := service.NewAWSEC2Service(cfg)
	if err != nil {
		log.Fatal("EC2 서비스 초기화 중 오류 발생:", err)
	}

	instance, err := ec2Service.ListInstances(context.Background())
	if err != nil {
		log.Fatal("인스턴스 조회 중 오류 발생: ", err)
	}

	presenter := presenter.NewConsolePresenter()
	presenter.Present(instance)

}
