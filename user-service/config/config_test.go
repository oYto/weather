package config

import "testing"

func TestInit(t *testing.T) {
	Init()
	t.Log(GetGlobalConfig().DbConfig)
	t.Log(GetGlobalConfig().LogConfig)
	t.Log(GetGlobalConfig().SvrConfig)
}
