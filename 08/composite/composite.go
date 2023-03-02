package composite

type IDepartment interface {
	Count() int
}

type FinanceDepartment struct {
	headCount int
}

func (dpt FinanceDepartment) Count() int {
	return dpt.headCount
}

type TechDepartment struct {
	headCount int
}

func (dpt TechDepartment) Count() int {
	return dpt.headCount
}

// HeadDepartment even if it implements IDepartment, it is acting as template, has departments
type HeadDepartment struct {
	departments []IDepartment
}

func (dpt HeadDepartment) Count() int {
	count := 0
	for _, d := range dpt.departments {
		count += d.Count()
	}
	return count
}
