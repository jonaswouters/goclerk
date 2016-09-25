package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jonaswouters/goclerk/models"
	"github.com/jonaswouters/goclerk/modules/setting"
	"github.com/jonaswouters/goclerk/modules/store"
	cliui "github.com/mitchellh/cli"
	"github.com/urfave/cli"
)

// Data command to import and export data from the database
var Data = cli.Command{
	Name:  "data",
	Usage: "Goclerk data tools",
	Subcommands: []cli.Command{
		{
			Name:   "export",
			Usage:  "Export goclerk data",
			Action: exportJson,
		},
		{
			Name:   "import",
			Usage:  "Import json data into goclerk",
			Action: importJson,
		},
	},
}

// exportJson will export all the data to json files
func exportJson(ctx *cli.Context) {
	setting.LoadSettings()
	ui := &cliui.BasicUi{Writer: os.Stdout, Reader: os.Stdin}

	path, _ := ui.Ask("Export folder (" + setting.Settings.ExportFolder + "):")

	if path != "" {
		setting.Settings.ExportFolder = path
	}

	// Check if the directory exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	fmt.Printf("Exporting to %s", setting.Settings.ExportFolder)

	exportInvoices(setting.Settings.ExportFolder)
	exportUsers(setting.Settings.ExportFolder)

	// Save updated settings
	setting.SaveSettings()
}

// reset will drop the database schema and run all migrations again
func importJson(ctx *cli.Context) {
	setting.LoadSettings()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exportUsers(path string) {
	var users []models.User
	err := store.GetDB().All(&users)

	if err != nil {
		fmt.Printf("Failed to retrieve users from the database %s", err.Error())
		return
	}

	for _, user := range users {
		b, _ := json.Marshal(user)

		err := ioutil.WriteFile(path+"users-"+user.ID.Hex()+".json", b, 0644)
		check(err)
	}
}

func exportInvoices(path string) {
	var invoices []models.Invoice
	err := store.GetDB().All(&invoices)

	if err != nil {
		fmt.Printf("Failed to retrieve invoices from the database %s", err.Error())
		return
	}

	for _, invoice := range invoices {
		b, _ := json.Marshal(invoice)
		// Convert bytes to string.

		err := ioutil.WriteFile(path+"invoice-"+invoice.ID.Hex()+".json", b, 0644)
		check(err)
	}
}
