package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"rentroll/rlib"
)

var backupDir = string("./bkup")

// AdmBkup is the HTTP handler for the Journal report request
func AdmBkup(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	sp := r.FormValue("sourcepage")

	// If the calling page was the admin bkup page then create the backup...
	if sp == "admbkup" {
		ok, _ := PathExists(backupDir)
		if !ok {
			err := os.Mkdir(backupDir, 0777)
			if err != nil {
				ui.ReportContent += fmt.Sprintf("Error creating directory %s: %s\n", backupDir, err.Error())
				return
			}
		}
		app := "./rrbkup"
		if err := exec.Command(app).Run(); err != nil {
			ui.ReportContent += fmt.Sprintf("*** Error *** running %s:  %v\n", app, err.Error())
			ui.ReportContent += fmt.Sprintf("*** Check the 'aws' command, does it need to be configured?\n")
		}
	}

	files, err := ioutil.ReadDir("./bkup")
	if err != nil {
		if os.IsNotExist(err) {
			ui.ReportContent += "no backup files"
			return
		}
		ui.ReportContent += "Error reading Database Backup directory: " + err.Error() + "\n"
		return
	}
	ui.ReportContent += "\nDatabase Backup Files\n"
	for _, file := range files {
		ui.ReportContent += file.Name() + "\n"
	}

}

// PathExists returns true if the path exists or false if it does not
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// AdmRestore is the HTTP handler for the Journal report request
func AdmRestore(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	sp := r.FormValue("sourcepage")
	// If the calling page was the admin bkup page then create the backup...
	if sp == "admrestore" {
		digits := r.FormValue("num")
		ok, _ := PathExists(backupDir)
		if !ok {
			err := os.Mkdir(backupDir, 0777)
			if err != nil {
				ui.ReportContent += fmt.Sprintf("Error creating directory %s: %s\n", backupDir, err.Error())
				return
			}
		}
		app := "./rrrestore"
		arg1 := "-d"
		if err := exec.Command(app, arg1, digits).Run(); err != nil {
			ui.ReportContent += fmt.Sprintf("*** Error *** running %s:  %v\n", app, err.Error())
			ui.ReportContent += fmt.Sprintf("*** Check the 'aws' command, does it need to be configured?\n")
		} else {
			ui.ReportContent += "Restore succeeded\n"
		}
	}

	files, err := ioutil.ReadDir(backupDir)
	if err != nil {
		if os.IsNotExist(err) {
			ui.ReportContent += "no backup files"
			return
		}
		ui.ReportContent += "Error reading Database Backup directory: " + err.Error() + "\n"
		return
	}
	ui.ReportContent += "\nDatabase Backup Files\n"
	for _, file := range files {
		ui.ReportContent += file.Name() + "\n"
	}
}
