package brandNamer

import (
	"fmt"
	"github.com/urfave/cli"
)

app := &cli.App{
	Name: "Brand Namer",
	Usage: "Create a random Brand Name",
	

	//Will use the same flags for all commands
	// so define it here
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name: "Brand Namer",
			Value: "Create Random Brand name"
		},

	}

	// create commands
	app.commands = []cli.Command{
		{
			Name: "bn",
			Usage: "Create random name w given # of letters",
			Flags: myFlags,
			// Action is code executed when we call "bn"
			Action: func(c *cli.Context) error {
				fmt.Println("Enter Number of letters")

				var num int
				fmt.Scanln(&num)
			}
		}
	}
}