// The main application entry point for hexlet-path-size.
package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			args := cmd.Args()

			if args.Len() > 0 {
				path := args.First()
				isRecursive := cmd.Bool("recursive")
				isHuman := cmd.Bool("human")
				isAll := cmd.Bool("all")

				size, err := code.GetPathSize(
					path, isRecursive, isHuman, isAll,
				)
				if err != nil {
					return fmt.Errorf("failed to get path size: %w", err)
				}

				fmt.Printf("%s\t%s\n", size, path)
				return nil
			}

			fmt.Println("Run the program with `-h` flag")
			return nil
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
