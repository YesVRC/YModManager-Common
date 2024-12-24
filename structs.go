package common

import (
	"errors"
	"fmt"
)

type Mod struct {
	Name              string `json:"name"`
	CurseForgeId      string `json:"curseforge_id"`
	CurseForgeVersion string `json:"curseforge_version"`
}

type InstanceType int

const (
	Vanilla InstanceType = iota
	Forge
	Fabric
)

var ModLoader = map[InstanceType]string{
	Vanilla: "Vanilla",
	Forge:   "Forge",
	Fabric:  "Fabric",
}

type Instance struct {
	Name      string       `json:"name"`
	Version   string       `json:"version"`
	Path      string       `json:"path"`
	Server    bool         `json:"server"`
	ModLoader InstanceType `json:"mod_loader"`
	Mods      []Mod        `json:"mods"`
}

func (i Instance) String() string {
	return fmt.Sprintf("%s (%s)", i.Name, i.Version)
}

func (i Instance) Install() error {
	return errors.New("notImplemented")
}
