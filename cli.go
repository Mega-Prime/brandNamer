package brandNamer

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func StartCLI() {

	app := cli.NewApp()
	app.Name = "Random Brand Namer"
	app.Usage = "Create random brand names"

	// Using the same flag for all commands
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "namer",
			Value: "BrandNamer",
		},
	}

	// Create commands
	app.Commands = []cli.Command{
		{
			Name:  "bn",
			Usage: "Create random brandName",
			Flags: myFlags,

			// Code that will execute when bn command is executed
			Action: func(c *cli.Context) error {
				fmt.Println("Enter number of letters")

				var num int
				fmt.Scanln(&num)
				fmt.Println("Your new brand name is: ", StringDefault(num))
				return nil
			},
		},
	}

	// start app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
