package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"ot/config"
	"ot/global"
	"ot/pkg/gormx"
	"ot/routers"
)

var (
	err error
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "c",
				Value:       "config/config.yaml",
				Usage:       "配置文件位置",
				Destination: &config.Opt.ConfigFile,
			},
		},
		Action: func(context *cli.Context) error {
			return nil
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	config.InitConfig(&config.Opt)

	global.GVA_DB = gormx.New()
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	router := routers.SetupRouter()
	err = router.Run(":9000")
	if err != nil {
		fmt.Println(err)
	}
}
