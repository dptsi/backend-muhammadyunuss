package main

// function
var i int

func isEven(i int) bool {
	return i%2==0
}

// method
type myInt int //need a type to bind method to
				//DOESN'T HAVE TO BE A STRUCT
func(i myInt) isEven() bool { //method receiver
	return int(i)%2==0
}

//---------------------------------------

// function
var i int
func isEven(i int) bool  {
	return i%2==0
}
ans := isEven(i)

//method
type myInt int //need a type to bind method to
var mi myInt
func(i myInt) isEven() bool { //method receiver
	return int(i)%2==0
}
ans = mi.isEven()