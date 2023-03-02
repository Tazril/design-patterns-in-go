package composite

import (
	"testing"
)

func TestHeadDepartment_Count(t *testing.T) {
	// given
	fd := FinanceDepartment{headCount: 30}
	td := TechDepartment{headCount: 50}
	hd := HeadDepartment{
		departments: []IDepartment{&fd, &td},
	}
	// when
	count := hd.Count()
	// then
	if count != 80 {
		t.Errorf("Got count = %v, Expected count = %v", count, 80)
	}
}
