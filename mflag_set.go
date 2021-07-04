package mflag

import (
	"fmt"
	"strings"
)

type MFlagSet struct {
	FlagSet
	plugins          map[string]map[string]*MFlagPlugin
	pluginPrefixList []string
}

func NewMFlagSet(name string, errorHandling ErrorHandling) *MFlagSet {
	f := &FlagSet{
		name:          name,
		errorHandling: errorHandling,
	}
	f.Usage = f.defaultUsage
	m := &MFlagSet{}
	m.FlagSet = *f
	return m
}

func (mflagSet *MFlagSet) FlagPlugin(pluginType string, pluginName string, usage string) *MFlagPlugin {
	if mflagSet.plugins == nil {
		mflagSet.plugins = make(map[string]map[string]*MFlagPlugin)
		mflagSet.pluginPrefixList = make([]string, 0)
	}
	plugin, ok := mflagSet.plugins[pluginType]
	if !ok {
		plugin = make(map[string]*MFlagPlugin)
		mflagSet.plugins[pluginType] = plugin
	}
	pluginPrefixs := []string{
		fmt.Sprintf("-%s.%s.", pluginType, pluginName),
		fmt.Sprintf("--%s.%s.", pluginType, pluginName),
	}
	plugin[pluginName] = NewMFlagPlugin(pluginType, pluginName, pluginPrefixs, usage, ExitOnError)
	mflagSet.pluginPrefixList = append(mflagSet.pluginPrefixList, pluginPrefixs...)
	return plugin[pluginName]
}

func (mflagSet *MFlagSet) removePluginFlag(arguments []string) []string {
	args := []string{}
	isPlugin := false
	for _, arg := range arguments {
		if arg[0] == '-' {
			isPlugin = false
			for _, prefix := range mflagSet.pluginPrefixList {
				if strings.HasPrefix(arg, prefix) {
					isPlugin = true
					break
				}
			}
		}
		if !isPlugin {
			args = append(args, arg)
		}
	}
	return args
}
