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

// AdmNewDB is the HTTP handler for the Journal report request
func AdmNewDB(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	app := "./rrnewdb"
	if err := exec.Command(app).Run(); err != nil {
		ui.ReportContent += fmt.Sprintf("*** Error *** running %s:  %v\n", app, err.Error())
		ui.ReportContent += fmt.Sprintf("*** Check the 'aws' command, does it need to be configured?\n")
	} else {
		ui.ReportContent = "New database created!"
	}
}

// CreateDBBackupFileList returns a string table of backup files and timestamps
func CreateDBBackupFileList() string {
	var t rlib.Table
	errmsg := ""
	t.Init()
	t.AddColumn("Filename", 30, rlib.CELLSTRING, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Modified", 23, rlib.CELLDATETIME, rlib.COLJUSTIFYLEFT)
	t.AddColumn("Size (bytes)", 12, rlib.CELLINT, rlib.COLJUSTIFYRIGHT)
	t.SetTitle("Database Backup Files\n\n")
	files, err := ioutil.ReadDir("./bkup")
	if err != nil {
		if os.IsNotExist(err) {
			errmsg += "no backup files"
			return errmsg
		}
		errmsg += "Error reading Database Backup directory: " + err.Error() + "\n"
		return errmsg
	}
	for _, file := range files {
		t.AddRow()
		t.Puts(-1, 0, file.Name())
		t.Putdt(-1, 1, file.ModTime())
		t.Puti(-1, 2, file.Size())
	}
	return t.String()
}

// AdmBkup is the HTTP handler for Backing up a database
func AdmBkup(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	sp := r.FormValue("sourcepage")
	fname := r.FormValue("filename")

	if len(fname) > 0 {
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
			args := []string{"-f", fname}
			if err := exec.Command(app, args...).Run(); err != nil {
				ui.ReportContent += fmt.Sprintf("*** Error *** running %s:  %v\n", app, err.Error())
				ui.ReportContent += fmt.Sprintf("*** Check the 'aws' command, does it need to be configured?\n")
			}
		}
	} else {
		ui.ReportContent += "*** Please enter a filename, then press Backup. ***\n\n"
	}
	ui.ReportContent += CreateDBBackupFileList()
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

// AdmRestore is the HTTP handler for restoring a database
func AdmRestore(w http.ResponseWriter, r *http.Request, xbiz *rlib.XBusiness, ui *RRuiSupport) {
	sp := r.FormValue("sourcepage")
	fname := r.FormValue("filename")
	// If the calling page was the admin bkup page then create the backup...
	if sp == "admrestore" && len(fname) > 0 {
		ok, _ := PathExists(backupDir)
		if !ok {
			err := os.Mkdir(backupDir, 0777)
			if err != nil {
				ui.ReportContent += fmt.Sprintf("Error creating directory %s: %s\n", backupDir, err.Error())
				return
			}
		}
		app := "./rrrestore"
		args := []string{"-f", fname}
		if err := exec.Command(app, args...).Run(); err != nil {
			ui.ReportContent += fmt.Sprintf("*** Error *** running %s:  %v\n", app, err.Error())
			ui.ReportContent += fmt.Sprintf("*** Check the 'aws' command, does it need to be configured?\n")
		} else {
			ui.ReportContent += "Restore succeeded\n"
		}
	}
	ui.ReportContent += CreateDBBackupFileList()
}
