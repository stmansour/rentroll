package rlib

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// QueryClause normally holds select, where, order clauses
type QueryClause map[string]string

// GetQueryClauseCopy returns the new copy of given QueryClause
func GetQueryClauseCopy(qc QueryClause) QueryClause {
	nQC := make(QueryClause, len(qc))
	for k, v := range qc {
		nQC[k] = v
	}
	return nQC
}

// SelectQueryFields holds the list of fields, used to get those from sql select query
type SelectQueryFields []string

// SelectQueryFieldMap holds the map of one field to list of fields,
// one field would be mapped to multiple field, e.g, result come from combining those fields
// ex.
// Payor is mapped to `Transactant.FirstName, Transactant.LastName`
// so Payor is being shown as group concat over firstname, lastname
// also when someone search in payor, ultimately it'll search in firstname, lastname with
// requested input.
type SelectQueryFieldMap map[string][]string

// RenderSQLQuery accepets queryForm (text form), QueryClause (map)
// and executes text template with given map of clause and return
func RenderSQLQuery(queryForm string, qc QueryClause) string {
	b := &bytes.Buffer{}
	template.Must(template.New("").Parse(queryForm)).Execute(b, qc)
	return b.String()
}

// GetQueryCount returns the number of records fetched by execution of query
func GetQueryCount(query string, qc QueryClause) (int64, error) {

	// if query ends with ';' then remove it
	query = strings.TrimSuffix(strings.TrimSpace(query), ";")

	// replace select clause first and get count query
	queryForm := `
    SELECT
        COUNT(*)
    FROM ({{.query}}) as T;
    `

	countQuery := RenderSQLQuery(queryForm, map[string]string{"query": query})
	fmt.Println("Count Query: ", countQuery)

	// hit the query and get count from db
	count := int64(0)
	var err error
	de := RRdb.Dbrr.QueryRow(countQuery).Scan(&count)
	if de != nil {
		err = fmt.Errorf("GetQueryCount: query=\"%s\"    err = %s", countQuery, de.Error())
	}
	return count, err
}
