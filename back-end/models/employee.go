package models

type DepartmentType struct {
	ID   int
	Type string
}

type PositionType struct {
	ID   int
	Type string
}

type Employee struct {
	ID         int
	Email      string
	Position   PositionType   `pg:"rel:has-one"`
	Department DepartmentType `pg:"rel:has-one"`
}
