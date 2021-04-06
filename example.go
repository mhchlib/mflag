package main

import "flag"

func main() {
	m := NewMflag(flag.FlagSet{},"example","This is for test")
	m.MflagSet.Int("testInt",123,"This is for testInt")
	m.MflagSet.String("testString","testString","This is for testString")
	m.MflagSet.Bool("testBool",false,"This is for testInt")


	BindPlugin(m,"plugin")

	PrintByPlugin()
}