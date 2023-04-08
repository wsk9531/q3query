package q3query

import (
	"fmt"
	"net/netip"
)

type Server struct {
	IP   netip.AddrPort
	Info ServerInfo
}

type ServerInfo struct {
	protocol      int
	hostname      string
	mapname       string
	clients       int
	sv_maxclients int
	gametype      string
	pure          bool
	game          string
}

func NewServer(data []byte) (s *Server, err error) {
	ip, err := netip.ParseAddrPort(fmt.Sprintf("%d.%d.%d.%d:%d", data[0], data[1], data[2], data[3], uint16(data[4])<<8|uint16(data[5])))
	if err != nil {
		return nil, err
	}

	s = &Server{
		IP:   netip.AddrPortFrom(ip.Addr(), ip.Port()),
		Info: *new(ServerInfo),
	}

	return s, nil
}
