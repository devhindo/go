package main

type Saiyan struct {
	Name string
	Power int
}

func (s *Saiyan) Super() {
	s.power += 10000
}

goku := Saiyan{
	Name: "ahmed",
	Power: 9000,
}

ob := Saiyan{}
ob.power = 33

ab = Saiyan{"ahmed",34}

func NewSaiyan (name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}

gouk := new(Saiyan)
gouk2 := &Saiyan{}

