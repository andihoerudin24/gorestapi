package main

import (
	"context"
	"github.com/urfave/cli/v2"
	"gorestapi/config"
	"gorestapi/database/migration"
	"gorestapi/routes"
	"log"
	"os"
)

var (
	db    = config.SetUp()
	ctx   = context.Background()
	redis = config.InitRedis()
)

func InitialCommand() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "db:migrate",
				Action: func(cli *cli.Context) error {
					db.AutoMigrate(&migration.User{})
					db.AutoMigrate(&migration.Post{})
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer config.CloseDatabase(db)
	defer redis.Close()

	_, err := redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	InitialCommand()
	r := routes.Router()
	r.Run(os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT"))

}
