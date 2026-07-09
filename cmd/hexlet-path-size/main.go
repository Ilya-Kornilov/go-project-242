// The main application entry point for hexlet-path-size.
package main

import (
	"code"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

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
				isAll := cmd.Bool("all")
				isRecursive := cmd.Bool("recursive")
				
				size, err := code.GetPathSize(path, isAll, isRecursive)
				if err != nil {
					return fmt.Errorf("failed to get path size: %w", err)
				}

				sizeStr := strings.TrimSuffix(size, "B")
				numSize, err := strconv.ParseInt(sizeStr, 10, 64)
				if err != nil {
					return fmt.Errorf("failed to parse size %s: %w", size, err)
				}
				formattedSize := code.FormatSize(numSize, cmd.Bool("human"))

				fmt.Printf("%s\t%s\n", formattedSize, path)
				return nil
			}
			fmt.Println("Hello from Hexlet!")
			return nil
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
