package main

import (
	"errors"
	"log"
	"os"

	"github.com/tflexsoom/gasm/internal/assembler"
	"github.com/tflexsoom/gasm/internal/linker"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gasm",
		Usage: "gasm tool suite for gasm projects and modules.",
		Commands: []*cli.Command{
			{
				Name:   "assemble",
				Usage:  "assemble an associated assembly project",
				Flags:  assembleFlags,
				Action: multiProjectCmd("assemble", assembleSubCmd),
			},
			{
				Name:   "parse",
				Usage:  "parse an associated assembly project",
				Flags:  baseFlags,
				Action: multiProjectCmd("parse", parseSubCmd),
			},
			{
				Name:  "debug",
				Usage: "utility tools for development",
				Subcommands: []*cli.Command{
					{
						Name:   "objectify",
						Usage:  "print parsed PE object file",
						Flags:  baseFlags,
						Action: multiProjectCmd("objectify", objectifySubCmd),
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var baseFlags = []cli.Flag{
	&cli.PathFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "Output file pathname for the backend output",
		Value:   "./a.out",
	},
	&cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "Print out debug information while performing work",
		Value:   false,
	},
}

func multiProjectCmd(
	cmdName string,
	cmdImpl func(cCtx *cli.Context) error,
) func(*cli.Context) error {
	return func(cCtx *cli.Context) error {
		if cCtx.Args().Len() < 1 {
			return errors.New("missing project destination! \"command [command options] [arguments...]\"")
		}

		return cmdImpl(cCtx)
	}
}

func parseSubCmd(cCtx *cli.Context) error {
	return assembler.ParseOnly(assembler.ParserOptions{
		Files:          cCtx.Args().Slice(),
		OutputLocation: cCtx.Path("output"),
		Verbose:        cCtx.Bool("verbose"),
	})
}

var assembleFlags = append(baseFlags,
	&cli.StringFlag{
		Name:    "language",
		Usage:   "Dialect of assembly desired to be assembled",
		Value:   "x86_64_windows",
		Aliases: []string{"l", "L"},
	},
	&cli.StringSliceFlag{
		Name:    "includes",
		Usage:   "Files to include for external symbols and libraries",
		Aliases: []string{"I", "include"},
	},
)

func assembleSubCmd(cCtx *cli.Context) error {
	return assembler.Assemble(assembler.AssemblerOptions{
		Files:          cCtx.Args().Slice(),
		Includes:       cCtx.StringSlice("includes"),
		OutputLocation: cCtx.Path("output"),
		Language:       cCtx.String("language"),
		Verbose:        cCtx.Bool("verbose"),
	})
}

func objectifySubCmd(cCtx *cli.Context) error {
	return linker.Objectify(linker.ObjectifyOptions{
		Files:          cCtx.Args().Slice(),
		OutputLocation: cCtx.Path("output"),
		Verbose:        cCtx.Bool("verbose"),
	})
}
