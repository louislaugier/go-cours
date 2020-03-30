package main

import "fmt"

type house struct {
	Area  float32
	Price float32
	Sims  []sims
}

type human struct {
	FirstName string
	LastName  string
	Age       int
}

type animal struct {
	Name string
	Age  int
}

type sims interface { // underlying type: *human / *animal
	Birthday()
}

func (a *animal) Birthday() {
	a.Age = a.Age + 1
	fmt.Println(a.Name, "est content")
}

func (h *human) Birthday() {
	h.Age = h.Age + 1
	fmt.Println(h.FirstName, "reçoit des cadeaux")
}

func main() {
	arthur := &human{
		FirstName: "Arthur",
	}
	// arthur.Birthday()
	myHome := &house{
		Sims: []sims{
			arthur,
			&animal{
				Name: "Felix",
			},
		},
	}
	fmt.Println(myHome)
	for _, s := range myHome.Sims {
		s.Birthday()
	}
}
