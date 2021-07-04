package mflag

var Mflags map[string][]*Mflag

type Mflag struct {
	MFlagSet
	FlagName    string
	Disturbuilt string
	PluginType  string
}

func BindPlugin(mflag *Mflag, pluginType string) {
	mflag.PluginType = pluginType
	if Mflags == nil {
		Mflags = make(map[string][]*Mflag)
	}
	Mflags[pluginType] = append(Mflags[pluginType], mflag)
}

func NewMflag(flag FlagSet, FlagName, Disturbuilt string) *Mflag {
	f := &Mflag{
		FlagName:    FlagName,
		Disturbuilt: Disturbuilt,
	}
	f.FlagSet = flag
	return f
}

func FlagPlugin(pluginType string, name string, usage string) *MFlagPlugin {
	return CommandLine.FlagPlugin(pluginType, name, usage)
}
