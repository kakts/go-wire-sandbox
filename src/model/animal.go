package model

type Animal interface {
	Bark() string
	Walk() int
}

type Dog struct {
	X int
}

func (d Dog) Bark() string {
	return "Bark: Dog"
}

func (d Dog) Walk() int {
	d.X += 1
	return d.X
}

func ProvideDog() Dog {
	return Dog{}
}

type Cat struct {
	X int
}

func (c Cat) Bark() string {
	return "Bark: Cat"
}

func (c Cat) Walk() int {
	c.X += 1
	return c.X
}

func ProvideCat() Cat {
	return Cat{}
}

type Owner struct {
	Animal Animal
}

func ProvideOwner(animal Animal) Owner {
	return Owner{Animal: animal}
}
