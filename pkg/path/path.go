package path

import (
	"os"
	"path/filepath"
)

// coredns
var (
	CoreDnsHome = os.Getenv("COREDNS_HOME")
)

// mate
var (
	CoreDnsMatePath    = filepath.FromSlash(CoreDnsHome + "/mate")
	CoreDnsScripts     = filepath.FromSlash(CoreDnsMatePath + "/scripts")
	CoreDnsStartScript = filepath.FromSlash(CoreDnsScripts + "/start-coredns.sh")
)
