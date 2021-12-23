package coredns

import (
	"coredns_mate/pkg/config"
	"coredns_mate/pkg/path"
	"coredns_mate/pkg/util"
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
)

func Start() {
	if !config.RemoteMode {
		startCoreDns()
	}
}

func startCoreDns() {
	stdout, stderr, err := gutil.CallScript(path.CoreDnsStartScript)
	util.Logger().Info("output is ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	util.Logger().Error("run start coredns scripts failed ", zap.Error(err))
}
