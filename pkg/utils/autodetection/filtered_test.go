package autodetection

import (
	"fmt"
	"testing"
)

func TestFilteredEnumeration(t *testing.T) {
	enumeration, c, err := FilteredEnumeration(nil, nil, nil, 4)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(*enumeration)
	fmt.Println(*c)
}
