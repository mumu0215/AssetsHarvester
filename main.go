package main

import (
	"AssetsHarvester/module"
	"flag"
	"fmt"
)
var(
	SearchStr=flag.String("d","","search domain string")
	ExcludeHoneyPot=flag.Bool("h",true,"is exclude honeypots, default True")
	Config= module.Config{}
)

func main()  {
	myConfig:=module.ParseConfig("config.yml")
	fmt.Println(module.FofaMain(`domain="suzhoubank.com"`,myConfig))
}