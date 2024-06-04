package internal

import "github.com/keyrotate/netbird/client/internal/stdnet"

func (e *Engine) newStdNet() (*stdnet.Net, error) {
	return stdnet.NewNetWithDiscover(e.mobileDep.IFaceDiscover, e.config.IFaceBlackList)
}
