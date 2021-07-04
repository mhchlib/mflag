package mflag

import (
	"fmt"
	"strings"
)

type MFlagPlugin struct {
	pluginType    string
	pluginName    string
	pluginPrefixs []string
	pluginUsage   string
	FlagSet
}

func NewMFlagPlugin(pluginType string, pluginName string, pluginPrefixs []string, pluginUsage string, errorHandling ErrorHandling) *MFlagPlugin {
	f := NewFlagSet(fmt.Sprintf("%s.%s", pluginType, pluginName), errorHandling)
	plugin := &MFlagPlugin{
		pluginType:    pluginType,
		pluginName:    pluginName,
		pluginUsage:   pluginUsage,
		pluginPrefixs: pluginPrefixs,
	}
	plugin.FlagSet = *f
	return plugin
}

func (f *MFlagPlugin) PrintDefaults() {
	fmt.Printf("%s: %s\n", f.pluginName, f.pluginUsage)
	f.VisitAll(func(flag *Flag) {
		s := fmt.Sprintf("  -%s.%s.%s", f.pluginType, f.pluginName, flag.Name) // Two spaces before -; see next two comments.
		name, usage := UnquoteUsage(flag)
		if len(name) > 0 {
			s += " " + name
		}
		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if len(s) <= 4 { // space, space, '-', 'x'.
			s += "\t"
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			s += "\n    \t"
		}
		s += strings.ReplaceAll(usage, "\n", "\n    \t")

		if !isZeroValue(flag, flag.DefValue) {
			if _, ok := flag.Value.(*stringValue); ok {
				// put quotes on the value
				s += fmt.Sprintf(" (default %q)", flag.DefValue)
			} else {
				s += fmt.Sprintf(" (default %v)", flag.DefValue)
			}
		}
		fmt.Fprint(f.Output(), s, "\n")
	})
}

func (f *MFlagPlugin) extractArgs(arguments []string) []string {
	args := []string{}
	isPlugin := false
	for _, arg := range arguments {
		if arg[0] == '-' {
			isPlugin = false
		}
		for _, pluginPrefix := range f.pluginPrefixs {
			if strings.HasPrefix(arg, pluginPrefix) {
				isPlugin = true
				arg = "-" + arg[len(pluginPrefix):]
			}
		}
		if isPlugin {
			args = append(args, arg)
		}
	}
	return args
}
