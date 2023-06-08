package consul_demo

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
)

var ConsulCli *api.Client

func initConsulClient() {
	config := api.DefaultConfig()
	config.Address = "http://175.178.59.92:8500"
	ConsulCli, _ = api.NewClient(config)
}

func GetConsulCli() *api.Client {
	if ConsulCli == nil {
		initConsulClient()
		return ConsulCli
	}
	return ConsulCli
}

type KafkaConfig struct {
	Brokers []string
}

func GetConfig() []string {
	CI := GetConsulCli()
	kv := CI.KV()
	pair, _, err := kv.Get("kafka", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	var kafkaConfig KafkaConfig
	err = json.Unmarshal(pair.Value, &kafkaConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	return kafkaConfig.Brokers
}
