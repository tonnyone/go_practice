package generic

import "testing"

func TestSumIntsOrFloats(t *testing.T) {
	i := map[int]int64{
		1: 5,
		4: 10,
		2: 6,
	}

	if sum := SumIntsOrFloats[int, int64](i); sum != 21 {
		t.Error()
	}
	// omit these type arguments
	if sum := SumIntsOrFloats(i); sum != 21 {
		t.Error()
	}

	j := map[float64]float64{
		1.0: 5.0,
		4:   10,
		2:   6,
	}
	sum := SumIntsOrFloats(j)
	if sum != 21 {
		t.Error()
	}

}

func TestSumNumberStruct(t *testing.T) {
	type Person struct {
		name string
	}
	ps := map[Person]int64{
		{name: "susan"}: 1,
		{name: "soch"}:  5,
	}
	if numbers := SumNumbers(ps); numbers != 6 {
		t.Error()
	}
}
