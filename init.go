package main

import "rentroll/rlib"

func initLists() {
	App.asmt2str = make(map[int]string)
	App.asmt2int = make(map[string]int)

	s := "SELECT ASMTID,Name FROM assessmenttypes"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	var name string
	var id int
	for rows.Next() {
		rlib.Errcheck(rows.Scan(&id, &name))
		App.asmt2str[id] = name
		App.asmt2int[name] = id
	}
}
