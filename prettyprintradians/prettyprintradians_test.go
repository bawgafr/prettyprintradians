package prettyprintradians

import (
	"math"
	"testing"
)

func Test_simpleMultiplesOf2Pi(t *testing.T) {
	fourPi := math.Pi * 4.0

	got := PrettyPrint(fourPi)
	expected := "4π"
	if got != expected {
		t.Errorf("not converted correctly, expected %s got %s", expected, got)
	}

}

// func TestPrettyPrint(t *testing.T) {
// 	t.Run(
// 		""
// 	)
// }
