package onesite

import (
	"strings"
)

// IsValidRentableStatus checks that passed string contains valid rentable status
// acoording to rentroll system
func IsValidRentableStatus(s string) (bool, string, string) {
	found := false
	var tempRS, rentRollStatus string
	// first find that passed string contains any status key
	a := strings.ToLower(s)
	for k, v := range RentableStatusCSV {
		if strings.Contains(a, k) {
			tempRS = v
			rentRollStatus = k
			found = true
			break
		}
	}
	return found, rentRollStatus, tempRS
}
