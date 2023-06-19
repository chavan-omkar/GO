package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type user struct {
	firstName string
	lastName  string
	age       int
	// contact   contactInfo  // here we dont need to explicitly give property name as contact even we dont pass any name it will make contactInfo as a variable
	contactInfo
}

func main() {

	//please uncomment one by one function and try it out

	declareStruct()
	embaddedStruct()

	mySlice := []string{"Hii", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)

}

func declareStruct() {
	/* This approach means the order we have used while creating this struct it will automatically follows that order
	if in future some one changes the above declared field order then declared values for alex will also changes
	*/
	// alex := user{"Alex", "Anderson", 30}

	// fmt.Println(alex)

	/*
		By following this approach we can change the order in struct declarion anytime
		and it wont affect anything further and also this gives more clearer image to reader
	*/
	amanda := user{firstName: "Amanda", lastName: "keith", age: 25}

	fmt.Println(amanda)

	/*
		At this stage we have declared a variable named aman of type person
		here go assign this variable zero values as we have not given any values
		eg :-
		string	 ""
		int		 0
		float	 0
		bool	 false
		so here it is aman{firstName : "",lastName:"",age:0}
	*/
	var aman user
	aman.firstName = "Aman"
	aman.lastName = "Mishra"
	aman.age = 22
	fmt.Println(aman) // the ouput is {  0} here we can see two blank spaces means blank string are printed

	fmt.Printf("%+v", aman) // o/p is {firstName: lastName: age:0}

}

func embaddedStruct() {

	/*
		we can declare any type of property inside the struct
		This is an example of struct inside struct

		here intresting thing is that even in the last line also we need to give comma (,) when we are declaring all values on new line
	*/

	amanda := user{
		firstName: "Amanda",
		lastName:  "keith",
		age:       25,
		// contact: contactInfo{
		// 	email:   "amandakeith1140@gmail.com",
		// 	zipCode: 213675,
		// },
		contactInfo: contactInfo{
			email:   "amandakeith1140@gmail.com",
			zipCode: 213675,
		},
	}

	// fmt.Printf("%+v", amanda)

	// this returns the memory address of the variable amanda to variable amandaPointer
	amandaPointer := &amanda // turned value into memory address
	amanda.updateName("Ammy")
	amanda.print()

	//go will take of the memory address here we explicitlly dont need to get and pass the memory address to the function
	amanda.updateNameP("jimmy chew")
	amanda.print()
	amandaPointer.updateNameP("AmmyPointer")
	amanda.print()

}

// user type as a receiver function
func (u user) print() {
	fmt.Printf("%+v", u)
}

/*
this will not work
here after passing the string value of newFirstName go will not exactly update the existing data here
it will copy the existing data in other place in memory and will make that data availbale for updateName function to modify
*/
func (u user) updateName(newFirstName string) {
	u.firstName = newFirstName
}

/*
This function will solve the above mentioned problem using pointers
function with receiver of type pointer to user
*/
func (u *user) updateNameP(newFirstName string) {
	// *u menas returning the value stored at that memory location, or turned memory address into value again
	(*u).firstName = newFirstName

}

/*
value type datastructures := int,float,string,bool,structs so here we need to use pointers to modify the data

reference type datastructures := slices, maps,channels,pointers,functions this datastructures maintain the referance to memory location in themselves so during modifying this we don't need to make use of pointers

*/

func updateSlice(slice []string) {
	slice[0] = "Sayo Nara"

}
