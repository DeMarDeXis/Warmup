package main

import "fmt"

type Builder interface {
	Build()
}

type Person struct {
	Name string
	Age  int
}

type WorkExperience struct {
	Name string
	Age  int
}

func (p Person) printName() {
	fmt.Println(p.Name)
}

//
type WoodBuilder struct {
	Person
	//Name string //on 59
	//WorkExperience // on for 60 - 63
}

/* func (wb.WoodBuilder) printName() {
	fmt.Println(b.Name)
} */

func (wb WoodBuilder) Build() {
	fmt.Println("Строю дом из дерева")
}

type BrickBuilder struct {
	Person
}

func (bb BrickBuilder) Build() {
	fmt.Println("Строю из кирпича")
}

type Building struct {
	Builder
	Name string
}

func main() {
	//explanation()
	usecase()
}

func explanation() {
	builder := WoodBuilder{Person{Name: "Вася", Age: 30}} // alt 59
	//builder := WoodBuilder{Person{Name: "Вася"; Age: 30}, "Боб"} //on 25, but off 58
	/* builder := WoodBuilder{
	Person{Name: "Вася", Age: 30},
	"Боб",
	WorkExperience{Name: "Таксист", Age: 3}} */
	fmt.Printf("Type: %T Value: %#v\n", builder, builder)

	// shorthands
	fmt.Println(builder.Person.Age)
	fmt.Println(builder.Age) //if on 60 - 63 tnen  colliding, so off

	//shadowing
	fmt.Println(builder.Name)

	builder.printName() //ve look on 30 - 32 and watch
	//builder.Person.printName()
}

func usecase() {
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			Name: "Вася",
			Age:  40,
		}},
		Name: "Деревянная изба",
	}
	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{
			Person{
				Name: "Петя",
				Age:  30,
			},
		},
		Name: "Кирпичный дом",
	}
	brickBuilding.Build()
}
