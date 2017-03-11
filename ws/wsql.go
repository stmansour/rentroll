package ws

import (
	"fmt"
	"reflect"
	"rentroll/rlib"
)

// gridBuildQuery builds a query from the supplied base and the sort / search parameters
// in the supplied w2ui grid structure.  To play with this routine in isolation
// use this:
//				https://play.golang.org/p/HOkP77h0Ts
//
// Parameters:
// 		table - the name of the table to query
// 		srch  - the default where clause. Used if the Search info is empty. Does not require
//              the keyword "WHERE".  That is, flter == "BID=1" when you want the where clause
//              to be "WHERE BID=1"
// 		order - default sorting clause. Used when Sort is empty
//      p     - pointer to the struct associated with the database table. It is used to match
//              the fields passed in by the UI.  We need to determine what type of fields
//              they are in order to properly construct the WHERE clause
//
// Returns:
//     string - the full query
//     string - the WHERE clause suitable for a COUNT(*) query
//----------------------------------------------------------------------------------------------
func gridBuildQuery(table, srch, order string, d *ServiceData, p interface{}) (string, string) {
	// Handle Search
	q := "SELECT * FROM " + table + " WHERE"
	return gridBuildQueryWhereClause(q, table, srch, order, d, p)
}

func gridBuildQueryWhereClause(q, table, srch, order string, d *ServiceData, p interface{}) (string, string) {
	qw := ""
	if len(d.wsSearchReq.Search) > 0 {
		val := reflect.ValueOf(p).Elem() // reflect value of input p
		count := 0
		for i := 0; i < len(d.wsSearchReq.Search); i++ {
			if d.wsSearchReq.Search[i].Field == "recid" || len(d.wsSearchReq.Search[i].Value) == 0 {
				continue
			}
			// look for this field in p
			for j := 0; j < val.NumField(); j++ {
				field := val.Field(j)                   // this is field[j] of p
				n := val.Type().Field(j).Name           // variable name for field(i)
				if n != d.wsSearchReq.Search[i].Field { // is this the field we're looking for?
					continue
				}
				t := field.Type().String() // Is it a type we can handle?
				if t != "string" {         // TODO: handle all data types
					continue
				}
				switch d.wsSearchReq.Search[i].Operator {
				case "begins":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s like '%s%%'", &count)
				case "ends":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s like '%%%s'", &count)
				case "is":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s='%s'", &count)
				case "between":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s like '%%%s%%'", &count)
				default:
					fmt.Printf("Unhandled search operator: %s\n", d.wsSearchReq.Search[i].Operator)
				}
			}
		}
		if len(qw) > 0 {
			qw = fmt.Sprintf(" BID=%d AND (%s)", d.BID, qw)
		}
		q += qw         // add the WHERE information
		if count == 0 { // if we didn't match any of the search criteria...
			q += " " + srch // then revert to the default search clause
			qw = srch
		}
	} else {
		q += " " + srch // no search info supplied, use the default
		qw = srch
	}

	// Handle any Sorting requests
	q += " ORDER BY "
	if len(d.wsSearchReq.Sort) > 0 {
		for i := 0; i < len(d.wsSearchReq.Sort); i++ {
			if i > 0 {
				q += ","
			}
			q += d.wsSearchReq.Sort[i].Field + " " + d.wsSearchReq.Sort[i].Direction
		}
	} else {
		q += order
	}

	// now set up the offset and limit
	q += fmt.Sprintf(" LIMIT %d OFFSET %d", d.wsSearchReq.Limit, d.wsSearchReq.Offset)
	return q, qw
}

func gridHandleField(q, logic, field, value, format string, count *int) string {
	if *count > 0 {
		q += " " + logic
	}
	q += fmt.Sprintf(format, field, value)
	*count++
	return q
}

// GetRowCount returns the number of database rows in the supplied table with the supplied where clause
func GetRowCount(table, where string) (int64, error) {
	count := int64(0)
	var err error
	s := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", table, where)
	de := rlib.RRdb.Dbrr.QueryRow(s).Scan(&count)
	if de != nil {
		err = fmt.Errorf("GetRowCount: query=\"%s\"    err = %s", s, de.Error())
	}
	return count, err
}
