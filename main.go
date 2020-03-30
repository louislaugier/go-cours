package main

import "fmt"

type user struct {
	FirstName string
	LastName  string
	Age       int
	Parents   []*user
}

type animal struct {
	Name string
	Age  int
}

func (u *user) Birthday() { // pointeur = référence à une valeur qui existe déjà en mémoire
	fmt.Println("Avant :", u.Age)
	u.Age = u.Age + 1
	fmt.Println("Après :", u.Age)
}

func main() {
	// fmt.Println("Hello World")
	// var hello string
	// hello = "hello world"
	// hello := "hello world"
	// fmt.Println(hello)
	// users := [2]string{
	// 	"arthur",
	// 	"test",
	// }
	// users := []string{
	// 	"arthur",
	// 	"test",
	// 	"test b",
	// }
	// fmt.Println(users)
	// for _, user := range users {
	// 	fmt.Println(user)
	// }
	arthur := &user{
		FirstName: "Arthur",
		// LastName:  "De Test",
		Age: 14,
		// Parents: []string{
		// 	"toto",
		// 	"tata",
		// },
		Parents: []*user{
			&user{
				FirstName: "Mom",
				Age:       45,
			},
			&user{
				FirstName: "Dad",
				Age:       49,
			},
		},
	}
	fmt.Println("First name:", arthur.FirstName)
	fmt.Println("First name:", arthur.LastName)
	fmt.Println("Age before birthday:", arthur.Age)
	if arthur.Parents != nil {
		for _, u := range arthur.Parents {
			u.Birthday()
			fmt.Println("Age de", u.FirstName, u.Age)
		}
	}
	arthur.Birthday()
	fmt.Println("Age after birthday:", arthur.Age)
	fmt.Println(arthur)
}
