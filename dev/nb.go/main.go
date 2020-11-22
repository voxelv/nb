package main

// revive:disable:error-strings Errors should match existing `nb` formatting.

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	// "strings"
	"syscall"
)

type config struct {
	nbDir  string
	nbPath string
}

type subcommand struct {
	name  string
	usage string
}

func cmdRun(cfg config, args []string, env []string) error {
	var arguments []string

	if len(args) < 2 {
		return errors.New("Command required.")
	} else if 2 <= len(args) {
		arguments = args[2:]
	} else {
		arguments = []string{}
	}

	cmd := exec.Command(args[1], arguments...)
	cmd.Dir = cfg.nbDir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func configure() (config, error) {
	cfg := config{}

	cfg.nbDir = os.Getenv("NB_DIR")

	if cfg.nbDir == "" {
		if runtime.GOOS == "windows" {
			cfg.nbDir = os.Getenv("APPDATA")

			if cfg.nbDir == "" {
				cfg.nbDir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data", "nb")
			}

			cfg.nbDir = filepath.Join(cfg.nbDir, "nb")
		} else {
			cfg.nbDir = filepath.Join(os.Getenv("HOME"), ".nb")
		}
	}

	var err error

	if cfg.nbPath, err = exec.LookPath("nb"); err != nil {
		return config{}, err
	}

	return cfg, nil
}

// run loads the configuration and environment, then runs the subcommand.
func run() error {
	var cfg config
	var err error

	if cfg, err = configure(); err != nil {
		return err
	}

	args := os.Args
	env := os.Environ()

	if len(args) > 1 && args[1] == "run" {
		if err := cmdRun(cfg, args[1:], env); err != nil {
			return err
		}
	} else {
		if err := syscall.Exec(cfg.nbPath, args, env); err != nil {
			return err
		}
	}

	return nil
}

// main is the primary entry point for the program.
func main() {
	os.Exit(presentError(run()))
}

// presentError translates an error into a message presnted to the user and
// returns the appropriate exit code.
func presentError(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	return 0
}
