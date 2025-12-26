//go:build !linux

package device

import (
	"github.com/ufovpn-ru/wireguard-go/conn"
	"github.com/ufovpn-ru/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
