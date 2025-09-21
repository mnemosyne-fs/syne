package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Register() error {
	if Cli.Register.Force {
		if Cli.Register.Name == "" {
			return fmt.Errorf("Register requires name when using force.")
		}

		Reg[Cli.Register.Name] = &Server{Url: Cli.Register.Url}

		Pin.Stop("Done!")
		return Reg.Write()
	}

	Cli.Register.Name = strings.TrimSpace(Cli.Register.Name)

	if Cli.Register.Name == "" {
		path := fmt.Sprintf("http://%s/name", Cli.Register.Url)

		Pin.UpdateMessage("Getting Server Name...")
		name, err := http.Get(path)

		if err != nil {
			return err
		}
		if name.StatusCode != 200 {
			return fmt.Errorf("Server error")
		}
		name_bytes, err := io.ReadAll(name.Body)
		if err != nil {
			return err
		}
		Cli.Register.Name = string(name_bytes)
	}

	fmt.Printf("Adding %s %s\n", Cli.Register.Name, Cli.Register.Url)
	Reg[Cli.Register.Name] = &Server{Url: Cli.Register.Url}

	Pin.Stop("Done!")
	return Reg.Write()
}
