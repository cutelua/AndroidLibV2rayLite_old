package CoreI

import (
	v2core "v2ray.com/core"
)

type Status struct {
	IsRunning   bool

	Vpoint v2core.Server
}

func CheckVersion() int {
	return 21
}