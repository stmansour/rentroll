package ws

import (
	"encoding/json"
	"fmt"
	"strings"
)

// scan forward through b, starting at index i, look for char c
// return the index where found. Or return -1 if not found.
func scanFwd(b []byte, k int, c byte) int {
	for i := k; i < len(b); i++ {
		if b[i] == c {
			return i
		}
	}
	return -1
}

// JSONchangeParseUtil is used to handle the "updated" records from a grid in the UI and store only the
// the changes.  Go expects a well-defined struct for json.Unmarshal. But in this case, we cannot
// use a well defined struct.  For example, say we have a struct like PaymentType which looks like this:
//
//		type PaymentType struct {
//			PMTID       int64
//			BID         int64
//			Name        string
//			Description string
//			LastModTime rlib.JSONDateTime
//			LastModBy   int64
//		}
//
//	And here is a real example of a change from the UI (remember there is lots more data before and
//  after what's shown here, I'm just showing the relevant part for updates to a single PaymentType
//  record. Note that the client sends back an array of updates, but we'll just look at one of the
//  updates within the array:
//
//		b = []byte( ... [{"recid":2,"Name":"Check"} ...] )
//                       +------------^-----------+
//  If we unmarshal this data into a new struct variable p as follows...
//
// 		var p PaymentType
//		json.Unmarshal(b,&p)
//
//	...it will work, recid will essentially be ignored, and the Name field will update to the value
//  "Check".  However, the other values (PMTID, BID, Description, LastModTime, LastModBy) will remain
//  at their nil values. So instead, we will read the actual record that is being updated first into
//  a variable, then unmarshal to the same struct, this will essentially update the structure. Since
//  we don't know the ID of the struct that we need to read from the db until the data is unmarshaled,
//  we need to unmarshal twice... the first time to get the relevant ID(s) so we know what record to
//  load, the second time to update the actual database struct in memory then save it back to the db.
//
//  Parameters:
//		s = byte array holding the array of updates from the client. This will be parsed and the
//          individual array elements will be unmarshaled separately. This string should be
//			exactly what the client sends as its POST payload
//     fn = function to call with the string representing an individual record update. This
//          function will be called for each record found in the array. If it returns an error
//          then processing will be stopped.
//      d = Context information about this web service call
//
//  Go Playground sample app:  https://play.golang.org/p/Qvp7FQ9Lgg
//
func JSONchangeParseUtil(s string, F func(string, *ServiceData) error, d *ServiceData) error {
	chgs := "\"changes\":"
	i := strings.Index(s, chgs)
	if i < 0 {
		return fmt.Errorf("PANIC: cannot find changes!  0")
	}

	b := []byte(s[i+len(chgs):])
	i = 0
	i = scanFwd(b, i, '[') // search for beginning of array
	if i < 0 {
		return fmt.Errorf("format error. 1")
	}
	i++
	done := false
	for !done {
		i = scanFwd(b, i, '{') // find the opening brace of the next record to parse
		if i < 0 {
			return fmt.Errorf("format error. 2")
		}

		t := 1
		var j int
		for j = i + 1; j < len(b); j++ { // find matching closing brace...
			switch b[j] {
			case '{':
				t++
			case '}':
				t--
			}
			if t == 0 {
				break
			}
		}
		// fmt.Printf("terminated scan on j = %d, b[j] = %c\n", j, b[j])
		if t != 0 { // didn't find a closing brace ????  uh oh
			return fmt.Errorf("format error. 3")
		}
		b1 := b[i : j+1]
		// fmt.Printf("Found string: %s\n", string(b1))

		var f interface{}
		json.Unmarshal(b1, &f)
		m := f.(map[string]interface{})
		js := "{"
		thisIsMore := ""
		for k, v := range m {
			// fmt.Printf("k = %s,  v is type %T:  %v\n", k, v, v)
			vtype := fmt.Sprintf("%T", v)
			js += thisIsMore + fmt.Sprintf("%q:", k)
			if vtype == "string" {
				js += fmt.Sprintf("\"%v\"", v)
			} else {
				js += fmt.Sprintf("%v", v)
			}
			thisIsMore = ","
		}
		js += "}"

		err := F(js, d)
		if err != nil {
			return err
		}

		openParen := false
		for i = j + 1; i < len(b) && !openParen; i++ {
			switch b[i] {
			case '{':
				openParen = true
			case ']':
				done = true
			}
			if openParen {
				break
			}
		}
		if i >= len(b) {
			done = true
		}
	}
	return nil
}
