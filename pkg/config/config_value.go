package config

import "github.com/paashzj/gutil"

// coredns config
var (
	RemoteMode bool
)

func init() {
	RemoteMode = gutil.GetEnvBool("REMOTE_MODE", true)
}
