package lua

import "github.com/yuin/gopher-lua"

func RunLuaFile(path string) error {
	L := lua.NewState()
	defer L.Close()
	return L.DoFile(path)
}
