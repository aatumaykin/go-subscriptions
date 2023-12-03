package entity

type Cycle struct {
	ID   uint
	Name string
	Days uint
}

var (
	Weekly  = Cycle{1, "Weekly", 7}
	Monthly = Cycle{2, "Monthly", 30}
	Yearly  = Cycle{3, "Yearly", 365}
)
