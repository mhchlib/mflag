package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var Mflags map[string][]*Mflag

type Mflag struct {
	MflagSet flag.FlagSet
	FlagName string
	Disturbuilt string
	PluginType	string
}

func BindPlugin(mflag *Mflag,pluginType string)  {
	mflag.PluginType = pluginType
	if Mflags == nil {
		Mflags = make(map[string][]*Mflag)
	}
	Mflags[pluginType] = append(Mflags[pluginType], mflag)
}

func NewMflag(flag flag.FlagSet,FlagName,Disturbuilt string) *Mflag {
	return &Mflag{
		MflagSet: flag,
		FlagName: FlagName,
		Disturbuilt: Disturbuilt,
	}
}


func PrintByPlugin()  {
	if Mflags == nil {
		Mflags = make(map[string][]*Mflag)
	}
	for s, mflags := range Mflags {
		fmt.Fprint(os.Stderr,s+"\n")
		for _, mflag := range mflags {
			fmt.Println(" " + mflag.FlagName + " \t"+mflag.Disturbuilt)
			mflag.printDefaults()
		}
	}

}

func (m *Mflag) printDefaults() {
	m.MflagSet.VisitAll(func(flag2 *flag.Flag) {
		s := fmt.Sprintf("  -%s", flag2.Name)
		name, usage := flag.UnquoteUsage(flag2)
		if len(name) > 0 {
			s += " " + name
		}
		s += "\t"
		s += strings.ReplaceAll(usage, "\n", "\n    \t")
		fmt.Fprint(m.MflagSet.Output(), s, "\n")
	})
}
func (m *Mflag) PrintDefaults() {
	m.MflagSet.PrintDefaults()
}
