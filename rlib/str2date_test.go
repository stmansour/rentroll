package rlib

import "testing"

func TestStringToDate(t *testing.T) {
	var m = []string{
		"04/03/18",
		"4/3/18",
		"4/3/2018",
		"04/03/2018",
		"2018-04-03",
		"2018-04-03 17:04:00 PDT",
		"Sat, 20 Jan 2018 08:00:00 GMT",
		"4/3/2018 5:04 pm",
		"Apr 3, 2018",
		"April 3, 2018",
	}

	for i := 0; i < len(m); i++ {
		dt, err := StringToDate(m[i])
		//t.Logf("StringToDate( %s ), got %s\n", m[i], dt.Format(RRDATETIMEINPFMT))
		if err != nil {
			t.Errorf("*** ERROR *** StringToDate( %s ): err = %s", m[i], err.Error())
		}
		if dt.Year() != 2018 {
			t.Errorf("*** ERROR *** StringToDate( %s ): year = %d", m[i], dt.Year())
		}
	}
}
