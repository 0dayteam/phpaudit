package main

import (
	"context"
	"io/ioutil"
	"phpaudit/audit"
	"phpaudit/config"
)

func checkErr(err error) {
	if err != nil {

		panic(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	checkErr(err)
	conf, err := config.NewRunConfig(data)
	checkErr(err)

	ctx := context.TODO()
	audit.Run(ctx, *conf)
}
