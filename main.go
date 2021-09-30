package main

import (
	"AssetsHarvester/module"
	"flag"
	"fmt"
)
var(
	SearchStr=flag.String("d","","search domain string")
	ExcludeHoneyPot=flag.Bool("b",true,"is exclude honeypots, default True")
	Config= module.Config{}
)
func init(){

}

func main()  {
	myConfig:=module.ParseConfig("config.yml")
	a,b:=module.FofaMain(`domain="suzhoubank.com"`,myConfig)
	fmt.Println(a)
	fmt.Println(b)
}