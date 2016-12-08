package conf

import (
	"encoding/json"
	"os"
)

type ProducerConf struct {
	Addr         string
	Count        int
	TopicAuth    string
	TopicSetting string
	TopicPrice   string
	TopicStatus  string
	TopicWeiXin  string
}

type ConsumerConf struct {
	Addr string

	Topic    string
	Channels []string
}

type NsqConfiguration struct {
	Producer       *ProducerConf
	Consumer       *ConsumerConf
	ConsumerNotify *ConsumerConf
}

type ServerConfiguration struct {
	BindPort          string
	ReadLimit         uint16
	WriteLimit        uint16
	ConnTimeout       uint16
	ConnCheckInterval uint16
	ServerStatistics  uint16
}

type Configuration struct {
	Uuid   string
	Nsq    *NsqConfiguration
	Server *ServerConfiguration
}

var G_conf *Configuration

func ReadConfig(confpath string) (*Configuration, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	G_conf = &config

	return &config, err
}

func GetConf() *Configuration {
	return G_conf
}
