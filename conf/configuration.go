package conf

import (
	"encoding/json"
	"os"
)

type ProducerConf struct {
	Addr      string
	Count     int
	TopicAuth string
}

type ConsumerConf struct {
	Addr     string
	Topic    string
	Channels []string
}

type NsqConfiguration struct {
	Producer *ProducerConf
	Consumer *ConsumerConf
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

var g_conf *Configuration

func ReadConfig(confpath string) (*Configuration, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	g_conf := &Configuration{}
	err := decoder.Decode(g_conf)

	return g_conf, err
}

func GetConf() *Configuration {
	return g_conf
}
