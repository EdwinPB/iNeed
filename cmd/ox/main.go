package main

import (
	"context"
	"fmt"
	"os"

	"ineedApp/migrations"
	_ "ineedApp/app/tasks"
	_ "ineedApp/app/models"

	"github.com/wawandco/ox/cli"
	"github.com/wawandco/ox/plugins/tools/soda"
)

// main function for the tooling cli, will be invoked by Ox
// when found in the source code. In here you can add/remove plugins that
// your app will use as part of its lifecycle.
func main() {
	cli.Use(soda.NewCommand(migrations.FS()))
	err := cli.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("[error] %v \n", err.Error())

		os.Exit(1)
	}
}