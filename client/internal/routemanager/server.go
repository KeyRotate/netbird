package routemanager

import "github.com/keyrotate/netbird/route"

type serverRouter interface {
	updateRoutes(map[route.ID]*route.Route) error
	removeFromServerNetwork(*route.Route) error
	cleanUp()
}
