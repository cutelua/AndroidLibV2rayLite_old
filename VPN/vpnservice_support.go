package VPN

import (
	"context"
	"net"
	"syscall"
	"time"

	v2net "v2ray.com/core/common/net"
	v2internet "v2ray.com/core/transport/internet"
)

type protectSet interface {
	Protect(int) int
}

type VPNProtectedDialer struct {
	SupportSet protectSet
}

func (d VPNProtectedDialer) Dial(ctx context.Context,
	src v2net.Address, dest v2net.Destination, sockopt *v2internet.SocketConfig) (net.Conn, error) {

	dialer := &net.Dialer{
		Timeout:   time.Second * 16,
		DualStack: true,
		LocalAddr: resolveSrcAddr(dest.Network, src),
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				d.SupportSet.Protect(int(fd))
			})
		},
	}

	return dialer.DialContext(ctx, dest.Network.SystemString(), dest.NetAddr())
}

func resolveSrcAddr(network v2net.Network, src v2net.Address) net.Addr {
	if src == nil || src == v2net.AnyIP {
		return nil
	}

	if network == v2net.Network_TCP {
		return &net.TCPAddr{
			IP:   src.IP(),
			Port: 0,
		}
	}

	return &v2net.UDPAddr{
		IP:   src.IP(),
		Port: 0,
	}
}
