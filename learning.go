package main

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }
// func sayBye(n string) {
// 	fmt.Printf("Good Bye! %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}
// 	return initials[0], "_"
// }

// var score = 99.8

// func updateName(x string) {
// 	x = "yoshi"
// }

// func updateName(x string) string {
// 	x = "wedge"
// 	return x
// }

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.99
// }

func main() {

	// fmt.Println("Hello!")

	// var nameOne string = "mario"
	// var nameTwo = "luigi"
	// nameour := "Sushant"
	// fmt.Println(nameOne, nameTwo, nameour)

	// var age1 int = 20
	// var age2 = 30
	// age3 := 40
	// fmt.Println(age1, age2, age3)

	// var num1 int8 = 25
	// var num2 int8 = -128
	// var num3 uint = 1234
	// fmt.Println(num1, num2, num3)

	// var score1 float32 = 25.98
	// var score2 float64 = 8888888888888888.98
	// score3 := 1.5
	// fmt.Println(score1, score2, score3)

	// fmt.Print("hello")
	// fmt.Print("world! \n")

	// fmt.Println("my name is", nameour)

	// fmt.Printf("my name is %v and my age is %v \n", nameour, age2)
	// fmt.Printf("my name is %q and my age is %q \n", nameour, age2)
	//%f for float
	//%T for type of variable

	//arrays and slices
	//arrays
	//var ages [3]int = [3]int{20, 22, 25}
	// var ages = [3]int{20, 22, 25}
	// names := [4]string{"Sushant", "Ramesh", "Ram", "Shyam"}
	// names[1] = "Suresh"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))
	//slices

	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	//slice ranges
	// rangeOne := names[1:3]
	// range2 := names[1:]
	// range3 := names[:3]
	// fmt.Println(rangeOne)
	// fmt.Println(range2)
	// fmt.Println(range3)

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello!"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))

	// fmt.Println(strings.Index(greeting, "th"))

	// fmt.Println(strings.Split(greeting, " "))

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "bowser"))

	//loops

	// 	x := 0
	// 	for x < 5 {
	// 		fmt.Println("value of x is:", x)
	// 		x++
	// 	}
	//
	//
	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	//
	//
	//
	// for index, value := range names {
	// 	fmt.Printf("the position at index %v is %v \n", index, value)
	// }
	// for _, value := range names {
	// 	fmt.Printf("the value is %v \n", value)
	// 	value = "new string"
	// }

	// fmt.Println(names)

	//
	//booleans and conditionals
	//
	// age := 25

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	//
	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("Continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break
	// 	}
	// 	fmt.Printf("the value at position %v is %v \n", index, value)
	// }
	//

	//
	//Functions in go
	//
	// sayGreeting("Sushant")
	// sayBye("Sushant")

	// cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	// cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	//
	//
	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1, a2)
	// fmt.Printf("area of circle 1 is %0.3f and area of circle 2 is %0.3f", a1, a2)

	// fn1, sn1 := getInitials("Sushant Khanal")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("Trishan")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("Prasiddha Raj Gautam")
	// fmt.Println(fn3, sn3)

	// sayHello("Sushant")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	// menu := map[string]float64{
	// 	"soup":  320,
	// 	"Pie":   70,
	// 	"salad": 120,
	// }
	// fmt.Println(menu)
	// fmt.Println(menu["Pie"])

	// //looping maps
	// for k, v := range menu {

	// 	fmt.Println(k, "-", v)
	// }
	// // ints as key type
	// phonebook := map[int]string{
	// 	1234: "Sushant",
	// 	5678: "Trishan",
	// 	1020: "samrat",
	// }

	// fmt.Println(phonebook)
	// fmt.Println(phonebook[1234])

	// phonebook[1020] = "Pandey"
	// fmt.Println(phonebook)

	// phonebook[1020] = "Goredai"
	// fmt.Println(phonebook)

	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	// name := "tifa"

	// // updateName(name)
	// name = updateName(name)

	// fmt.Println(name)

	// // group B types -> slices, maps, functions
	// // pointer wrapper values
	// menu := map[string]float64{
	// 	"pie":       5.95,
	// 	"ice cream": 3.99,
	// }

	// updateMenu(menu)
	// fmt.Println(menu)

}
