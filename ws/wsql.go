package ws

import (
	"fmt"
	"reflect"
	"rentroll/rlib"
	"strings"
)

// Requirements:
// 1. Output is targeted for use with W2UI grid. So, the server needs to take as input the
//    search query as described here: http://w2ui.com/web/docs/1.5/grid . In a nutshell, the
//    W2UI grid describes the query in this JSON structure:
//
// 			{
// 			    "cmd"         : "get-records",
// 			    "limit"       : 100,
// 			    "offset"      : 0,
// 			    "selected"    : [1, 2],
// 			    "searchLogic" : "AND",
// 			    "search": [
// 			        { "field": "fname", "type": "text", "value": "vit", "operator": "is" },
// 			        { "field": "age", "type": "int", "value": [10, 20], "operator": "between" }
// 			    ],
// 			    "sort": [
// 			        { "field": "fname", "direction": "ASC" },
// 			        { "field": "Lname", "direction": "DESC" }
// 			    ]
// 			}
//
// 2. The successful reply is a JSON solution set that is of this form:
//
//          {
//              "status"  : "success",
//              "total"   : 873,			// the total number of records that match the query
//              "records" : [
//                  { "recid": 1, "field-1": "value-1", ... "field-N": "value-N" }
//					...
//              ]
//          }
//
// 		a) Note that we must be able to produce a count of the total number of records that match the
// 		   query.  This value is independent of the LIMIT and OFFSET values.  In other words, suppose
// 		   that the solution set for a particular query has 600 rows. Regardless of the values for
// 		   LIMIT and OFFSET, we must return a value of 600 for "total".  This suggests that the coded
//         solution for these queries will be able to return both a "COUNT(*)" query as well as a
//         a query that provides the record fields.
//

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
					rlib.Console("Unhandled search operator: %s\n", d.wsSearchReq.Search[i].Operator)
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
	var qOrder string
	if len(d.wsSearchReq.Sort) > 0 {
		for i := 0; i < len(d.wsSearchReq.Sort); i++ {
			// do not sorting over recid, no such field exists actually on any datatype
			if d.wsSearchReq.Sort[i].Field == "recid" || len(d.wsSearchReq.Sort[i].Direction) == 0 {
				continue
			}
			if i > 0 {
				qOrder += ","
			}
			qOrder += d.wsSearchReq.Sort[i].Field + " " + d.wsSearchReq.Sort[i].Direction
		}
	} else {
		qOrder += order
	}

	// if any parameters are there for order by clause then
	if len(qOrder) > 0 {
		q += " ORDER BY " + qOrder
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

// GetSQLOrderClause builds order by clause of SQL query
func GetSQLOrderClause(fieldMap map[ColSort][]string) string {
	var orderClause string
	var count int
	for gsrt, fieldList := range fieldMap {
		var fieldOrder string
		for i, mf := range fieldList {
			fieldOrder += fmt.Sprintf("%s %s", mf, strings.ToUpper(gsrt.Direction))
			if i != len(fieldList)-1 {
				fieldOrder += ", "
			}
		}
		orderClause += fieldOrder
		if count != len(fieldMap)-1 {
			orderClause += ", "
		}
		count++
	}
	return orderClause
}

// GetSQLWhereClause builds where clause of sql query
func GetSQLWhereClause(fieldMap map[GenSearch][]string, searchLogic string) string {
	var whereClause string
	var count int

	// TODO: handle date type proper for LIKE operator

	for gsrch, fieldList := range fieldMap {
		var fieldSearch string
		var likeFmt string

		switch gsrch.Operator {
		case "begins":
			likeFmt = "%s like '%s%%'"
		case "ends":
			likeFmt = "%s like '%%%s'"
		case "is":
			likeFmt = "%s='%s'"
		case "between":
			likeFmt = "%s like '%%%s%%'"
		case "contains":
			likeFmt = "%s like '%%%s%%'"
		default:
			rlib.Console("Unhandled search operator: %s\n", gsrch.Operator)
		}

		for i, mf := range fieldList {
			fieldSearch += fmt.Sprintf(likeFmt, mf, gsrch.Value)
			if i != len(fieldList)-1 {
				fieldSearch += " OR "
			}
		}
		whereClause += fieldSearch
		if count != len(fieldMap)-1 {
			whereClause += " " + searchLogic + " "
		}
		count++
	}
	return whereClause
}

// GetSearchAndSortSQL returns where and order by clause
func GetSearchAndSortSQL(d *ServiceData, fieldMap map[string][]string) (string, string) {
	rlib.Console("Entered GetSearchAndSortSQL\n")
	var (
		reqWhereClause     string
		reqOrderClause     string
		reqSearchClauseMap = make(map[GenSearch][]string)
		reqOrderClauseMap  = make(map[ColSort][]string)
	)

	// get search clause first
	for _, gsrch := range d.wsSearchReq.Search {
		if gsrch.Field == "recid" || len(gsrch.Value) == 0 {
			continue
		}

		// if requested search doesn't exist in map then skip it
		if _, ok := fieldMap[gsrch.Field]; !ok {
			continue
		}
		reqSearchClauseMap[gsrch] = fieldMap[gsrch.Field]
	}
	reqWhereClause = GetSQLWhereClause(reqSearchClauseMap, d.wsSearchReq.SearchLogic)

	// get order clause then
	for _, gsrt := range d.wsSearchReq.Sort {
		if gsrt.Field == "recid" || len(gsrt.Direction) == 0 {
			continue
		}

		// if requested search doesn't exist in map then skip it
		if _, ok := fieldMap[gsrt.Field]; !ok {
			continue
		}
		reqOrderClauseMap[gsrt] = fieldMap[gsrt.Field]
	}
	reqOrderClause = GetSQLOrderClause(reqOrderClauseMap)

	return reqWhereClause, reqOrderClause
}
