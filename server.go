package charging_pile_das

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/mq"
	"github.com/giskook/gotcp"
	"log"
	"net"
	"time"
)

type ServerConfig struct {
	Listener      *net.TCPListener
	AcceptTimeout time.Duration
	NsqConf       *conf.NsqConfiguration
}

type Server struct {
	config           *ServerConfig
	srv              *gotcp.Server
	mq               *mq.NsqSocket
	checkconnsticker *time.Ticker
}

var Gserver *Server

func SetServer(server *Server) {
	Gserver = server
}

func GetServer() *Server {
	return Gserver
}

func NewServer(srv *gotcp.Server, nsq_socket *mq.NsqSocket, config *ServerConfig) *Server {
	serverstatistics := conf.GetConf().Server.ServerStatistics
	return &Server{
		config:           config,
		srv:              srv,
		mq:               nsq_socket,
		checkconnsticker: time.NewTicker(time.Duration(serverstatistics) * time.Second),
	}
}

func (s *Server) Start() {
	go s.mq.Start()
	go s.checkStatistics()

	s.srv.Start(s.config.Listener, s.config.AcceptTimeout)
}

func (s *Server) Stop() {
	s.srv.Stop()
	s.checkconnsticker.Stop()
	s.mq.Stop()

}

func (s *Server) checkStatistics() {
	for {
		<-s.checkconnsticker.C
		log.Printf("---------------------Total Connections : %d---------------------\n", conn.NewConns().GetCount())
	}
}
