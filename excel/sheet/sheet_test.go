package sheet

import (
	"fmt"
	"testing"
)

func TestConvertToColName(t *testing.T) {
	for i := 1; i < 200; i++ {
		fmt.Println(ConvertToColName(i), " -- ", i)
	}
	return
}
