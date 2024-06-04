package firewall

import "github.com/keyrotate/netbird/iface"

// IFaceMapper defines subset methods of interface required for manager
type IFaceMapper interface {
	Name() string
	Address() iface.WGAddress
	IsUserspaceBind() bool
	SetFilter(iface.PacketFilter) error
}
