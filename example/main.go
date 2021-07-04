package main

import (
	"github.com/mhchlib/mflag"
	"log"
)

func main() {
	baseFlag := mflag.String("base", "i am base", "this is a base flag")
	etcdStoreFlagPlugin := mflag.FlagPlugin("store", "etcd", "this is a etcd store plugin")
	a1 := etcdStoreFlagPlugin.Int("testInt", 20, "dsada")
	b1 := etcdStoreFlagPlugin.Bool("testBool", true, "dsada")
	c1 := etcdStoreFlagPlugin.String("testString", "mconfig is ok", "dsada")
	mysqlStoreFlagPlugin := mflag.FlagPlugin("store", "mysql", "this is a mysql store plugin")
	a2 := mysqlStoreFlagPlugin.Int("testInt", 20, "dsada")
	b2 := mysqlStoreFlagPlugin.Bool("testBool", false, "dsada")
	c2 := mysqlStoreFlagPlugin.String("testString", "mconfig is ok", "dsada")
	mflag.Parse()
	log.Println(*baseFlag)
	log.Println(*a1, *b1, *c1)
	log.Println(*a2, *b2, *c2)
}
