//go:build !android

package peer

import (
	"github.com/keyrotate/netbird/client/internal/stdnet"
)

func (conn *Conn) newStdNet() (*stdnet.Net, error) {
	return stdnet.NewNet(conn.config.InterfaceBlackList)
}
