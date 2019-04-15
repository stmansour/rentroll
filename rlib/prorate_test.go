package rlib

import (
	"reflect"
	"testing"
)

// simpleProrateInput struct to test simple proprated amount
type simpleProrateInput struct {
	amount                  float64
	rentCycle, prorateCycle int64
	d1, d2, epoch           string
}

// simpleProrateOutput struct to hold expected output for a scenario
type simpleProrateOutput struct {
	total  float64
	np, tp int64
}

// simpleProrateCase struct holds input & output
type simpleProrateCase struct {
	simpleProrateInput
	simpleProrateOutput
}

// TestSimpleProrate tests all possible scenarios of simple prorate cycle
func TestSimpleProrate(t *testing.T) {
	var cases = []simpleProrateCase{
		{
			simpleProrateInput{1000.00, RECURMONTHLY, RECURDAILY, "2018-01-01", "2018-02-01", "2018-01-01"}, // INPUT
			simpleProrateOutput{1000.00, 31, 31}, // EXPECTED OUTPUT
		},
		{
			simpleProrateInput{1000.00, RECURMONTHLY, RECURDAILY, "2018-01-10", "2018-02-01", "2018-01-01"}, // INPUT
			simpleProrateOutput{709.68, 22, 31}, // EXPECTED OUTPUT
		},
		{
			simpleProrateInput{1000.00, RECURMONTHLY, RECURDAILY, "2018-01-10", "2018-02-01", "2018-01-03"}, // INPUT
			simpleProrateOutput{709.68, 22, 31}, // EXPECTED OUTPUT
		},
		{
			simpleProrateInput{1000.00, RECURMONTHLY, RECURDAILY, "2018-01-01", "2018-01-11", "2018-01-01"}, // INPUT
			simpleProrateOutput{322.58, 10, 31}, // EXPECTED OUTPUT
		},
	}

	for i := 0; i < len(cases); i++ {
		// GET THE INPUTS FROM STRUCT
		amt := cases[i].simpleProrateInput.amount
		rc := cases[i].simpleProrateInput.rentCycle
		pc := cases[i].simpleProrateInput.prorateCycle
		d1, _ := StringToDate(cases[i].simpleProrateInput.d1)
		d2, _ := StringToDate(cases[i].simpleProrateInput.d2)
		epoch, _ := StringToDate(cases[i].simpleProrateInput.epoch)
		tot, np, tp := SimpleProrateAmount(amt, rc, pc, &d1, &d2, &epoch)

		// LOG MESSAGE
		t.Logf("info: case[%d] : SimpleProrateAmount( %f, %s, %s, %s, %s, %s ) expect ( %f, %d, %d ), got ( %f, %d, %d )\n",
			i,
			amt, RentalPeriodToString(rc), ProrationUnits(pc), d1.Format(RRDATEFMT3), d2.Format(RRDATEFMT3), epoch.Format(RRDATEFMT3),
			cases[i].simpleProrateOutput.total, cases[i].simpleProrateOutput.np, cases[i].simpleProrateOutput.tp,
			tot, np, tp,
		)

		// IF NOT EQUAL THEN
		if !reflect.DeepEqual(cases[i].simpleProrateOutput, simpleProrateOutput{tot, np, tp}) {
			t.Errorf("error: case[%d] : SimpleProrateAmount( %f, %s, %s, %s, %s, %s ) expect ( %f, %d, %d ), got ( %f, %d, %d )\n",
				i,
				amt, RentalPeriodToString(rc), ProrationUnits(pc), d1.Format(RRDATEFMT3), d2.Format(RRDATEFMT3), epoch.Format(RRDATEFMT3),
				cases[i].simpleProrateOutput.total, cases[i].simpleProrateOutput.np, cases[i].simpleProrateOutput.tp,
				tot, np, tp,
			)
		}
	}
}
