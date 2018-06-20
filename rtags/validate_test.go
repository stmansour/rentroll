package rtags

import (
	"fmt"
	"testing"
)

// TestUserStruct just a sample test with all combine tag rules in one struct
func TestUserStruct(t *testing.T) {
	type User struct {
		ID    int    `validate:"number,min=1,max=1000"`
		Name  string `validate:"string,min=8,max=50"`
		Bio   string `validate:"string"`
		Email string `validate:"email"`
	}

	tables := []struct {
		user      User
		errsCount int
	}{
		{User{ID: 0, Name: "short", Bio: "", Email: "foobar@.c"}, 4},
		{User{ID: 1, Name: "John Doe", Bio: "Hello, World!", Email: "john@doe.c"}, 0},
	}

	for _, tt := range tables {
		// call function
		errs := ValidateStructFromTagRules(tt.user)

		// should return list of errors
		if len(errs) != tt.errsCount {
			fmt.Println("Errors for debugging purpose:")
			for i, err := range errs {
				fmt.Printf("\t%d. %s\n", i+1, err.Error())
			}
			t.Fatalf("Expected length of errors %d, got %d", tt.errsCount, len(errs))
		}
	}
}
