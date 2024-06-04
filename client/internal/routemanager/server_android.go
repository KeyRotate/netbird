//go:build android

package routemanager

import (
	"context"
	"fmt"

	firewall "github.com/keyrotate/netbird/client/firewall/manager"
	"github.com/keyrotate/netbird/client/internal/peer"
	"github.com/keyrotate/netbird/iface"
)

func newServerRouter(context.Context, *iface.WGIface, firewall.Manager, *peer.Status) (serverRouter, error) {
	return nil, fmt.Errorf("server route not supported on this os")
}
