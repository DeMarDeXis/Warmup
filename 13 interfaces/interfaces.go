package interfaces

import "fmt"

//как объявляется интерфейс: type NAME interface{сигны наших методов()}
type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

//GOlang поддерживает встраивания
type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name) //Имплементируем интерфейс Runner
}

/* Значения интерфейсов. Есть значения интерфейсного типа
и есть значения конкретного типа.
До этого значения были конкретного типа Runner Swimmer Flyer.
InterfaceValues записано в тетрадке. */

func main() {
	InterfaceValues()
}

func InterfaceValues() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner) //nil*вый
	if runner == nil {
		fmt.Println("Runner is nil") // If the interface does not have a concrete type and a concrete value = nil
	}

	fmt.Println("")
	//Trying to assign a value
	//runner = int64(1) //panic У int64 нету метода Runner
	//runner.Run() //Panic, потому что интерфейсное и конкретное значение = nil

	// Имлементация интерфейса - это когда наш конкретный тип реализует все методы, которые перечисленны в интерфейсе.

	//trying to assign a value AGAIN

	//В ГО имплементация не явная, то есть без спец.слов, а используется DUCK TYPING

	// Для примера создадим struct Human ^ и для него объявляю метод Run(), которая возвращает строку

	var unnamedRunner *Human //указатель на Human
	fmt.Printf("Type: %T Value: %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("Runner is nil")
	}
	//Обрати внимание что Runner поменялся и проверка не прошла
	/* По поводу проверки, не смотря на то что конкретное
	значение конкретного типа = nil, само значение интерфейса != nil, т.к как минимум
	у нас есть инфа про конкретный тип. */

	fmt.Println("")
	//3 пример

	namedRunner := &Human{Name: "Jack"}
	fmt.Printf("Type: %T Value: %#v\n", namedRunner, namedRunner)
	runner = namedRunner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	/*Мы заполнили конкретное значение. Если смотреть по тетрадке то перешли на 3 стадию,
	то есть все конкретные данные(тип и значение) известны */

	//EmptyInterface{}
	//Это интерфейс у которого нет методов
	var emptyinterface interface{} = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", emptyinterface, emptyinterface)

	emptyinterface = runner
	fmt.Printf("Type: %T Value: %#v\n", emptyinterface, emptyinterface)

	emptyinterface = int64(1)
	fmt.Printf("Type: %T Value: %#v\n", emptyinterface, emptyinterface)

	emptyinterface = true
	fmt.Printf("Type: %T Value: %#v\n", emptyinterface, emptyinterface)
	//Ну и вот крч про интерфейсы
}

func typeAssertionAndPolymorphism() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	john := &Human{"John"}
	runner = john
	polymorphism(john)
	//typeAssertion(john)

	donald := &Duck("Donald", "Duck")
	runner = donald
	polymorphism(donald)
	//typeAssertion(donald)
}
func polymorphism(runner Runner) {
	fmt.Printf(runner.Run())
}

func typeAssertion(runner Runner) { //Утверждение типа
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if human, ok := runner.(*Human); ok { //ok это ФЛАГ удалось ли преобразовать значение interface типа в конкретный тип
		fmt.Printf("Type: %T Value: %#v\n", human, human)
		human.writeCode()
	}

	if duck, ok := runner.(*Duck); ok {
		fmt.Printf("Type: %T Value: %#v\n", duck, duck)
		fmt.Println(duck.Fly())
	}

	//Есть более короткий метод проверок-Type switch, но смотрит не на значение, а на тип
	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("Type: %T Value: %#v", v, v)
	}
}

func (h Human) writeCode() {
	fmt.Println("Dude write code.")
}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func (d Duck) Swim() string {
	return "Утка плавает"
}

func (d Duck) Fly() string {
	return "Утка плавает"
}
