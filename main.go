package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/yarlson/pin"
)

var Reg ServerRegistry

var Cli struct {
	Guest  bool   `short:"g" help:"Execute as guest"`
	Server string `short:"s" help:"Server to use"`

	Register struct {
		Force bool   `short:"f" help:"Don't check if server works"`
		Name  string `short:"n" help:"Give a name for the server" default:""`

		Url string `arg:"" help:"Url to the server"`
	} `cmd:"" help:"Registers a server"`

	Login  struct{} `cmd:"" help:"Login to account"`
	Logout struct {
		Others bool `short:"o" help:"Logout all others"`
	} `cmd:"" help:"Logout from account"`

	Init struct {
		Path string `arg:"" help:"Subdirectory on server" default:"/"`
	} `cmd:"" help:"Initialises a directory for syncing"`
	Sync struct {
		Path string `arg:"" help:"File or Folder to upload" default:"~"`
	} `cmd:"" help:"Syncs the active folder"`

	Get struct {
		Path string `arg:"" help:"Path to get" default:"/"`
	} `cmd:"" help:"Download a file"`
	Add struct {
		From string `arg:"" help:"What file to upload" type:"path"`
		To   string `arg:"" help:"Where to upload the file" default:"/"`
	} `cmd:"" help:"Add a File or Folder to the server"`
	Mv struct {
		From string `arg:"" help:"The File / Folder to rename"`
		To   string `arg:"" help:"The new name / path"`
	} `cmd:"" help:"Renames a File or Folder on the server"`
	Del struct {
		Path string `arg:"" help:"File or Folder to delete"`
	} `cmd:"" help:"Deletes a File or Folder"`
	Info struct {
		Path string `arg:""`
	} `cmd:"" help:"Get info on a file or folder"`

	Share struct {
		Path      string `arg:""`
		Password  string `short:"p"`
		NumUses   int    `short:"n"`
		AliveTime int    `short:"t"`
	} `cmd:"" help:"Temporarily share a file or folder"`
}

var Pin *pin.Pin

func main() {
	var err error
	ctx := kong.Parse(&Cli,
		kong.HelpOptions{
			Compact: true,
		})

	Reg, err = ParseRegistry()
	if err != nil {
		log.Fatal(err)
	}

	Pin = pin.New("Syne...")
	cancel := Pin.Start(context.Background())
	defer cancel()

	switch ctx.Command() {
	case "register <url>":
		err = Register()
	default:
	}

	if err != nil {
		fmt.Println("Error:")
		fmt.Print(err)
	}
}
