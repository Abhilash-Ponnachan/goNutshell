/*
This gonutshell package will attempt to to present a quick tour of
golang programming concepts with some toy examples.
*/
// declare package as type main to make it executable
package main

// import required packages
import (
	"bytes"
	"fmt"
	_ "fmt"
)

// define entry point
func main() {
	// ==== Variable declaration =====
	var x int
	/*
		var <identifier> <type>
		Unlike C/C++/Java/C# etc the 'type'
		is specified After the variable name
	*/
	// assign a value to a declared variable
	x = 23

	// declaration with initialization
	y := 3.14159
	// NOTE: here the type is inferred

	// ==== Formatted Printing to Console ====
	// printing values using 'fmt.Printf' method
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %f\n", y)
	/*
		Printf - works like 'printf' in C
		It takes a string with 'format specifiers'
		(called verbs) and then values that are
		formatted and printed replacing the verbs
		(using the indicated format)
	*/
	// common format specifier verbs
	// integer
	fmt.Printf("%d\n", 43)
	// integer with padding
	fmt.Printf("%04d\n", 43)
	// integer padding with space
	fmt.Printf("% 4d\n", 43)
	// char, quoted char
	fmt.Printf("%c - %q\n", 'Z', 'Z')
	// binary , octal, hex
	fmt.Printf("%b - %o - %x - %#x\n", 32, 32, 32, 32)
	// float
	fmt.Printf("%f\n", 31.3)
	// float with precission
	fmt.Printf("%06.2f\n", 31.3)
	// float exponential notation
	fmt.Printf("%e\n", 31.3)
	// bool
	fmt.Printf("%t\n", false)
	// pointer
	fmt.Printf("%p\n", &x) // address of 'x' in this case
	// string, quoted string
	fmt.Printf("%s - %q\n", "Hello", "Hello")
	// *** default format
	fmt.Printf("%v\n", []int{1, 2}) // value is a 'slice'
	// [1 2]
	// *** Go code format
	fmt.Printf("%#v\n", []int{1, 2}) // value is a 'slice'
	// []int{1, 2}
	// *** print the 'type'
	fmt.Printf("%T\n", []int{1, 2}) // value is a 'slice'
	// []int
	// we can use this to print the type of variables
	fmt.Printf("Type of 'y' = %T\n", y)

	// ==== Formatted Print output to string ====
	fs := fmt.Sprintf("%x", 2047)
	fmt.Println(fs)
	// ==== Printing without format specification ====
	fmt.Print("Print Hello.. ")
	fmt.Println("Println Hello!")
	// ==== Invoking functions ====
	printDouble(23)
	fmt.Printf("Triple of %d = %d\n", 23, triple(23))
	// *** variable destructuring - multiple assignment
	var sum int
	var prod int
	sum, prod = sumAndProd(23, 23)
	// NOTE: sum, prod := sumAndProd - does NOT work!
	fmt.Printf("Sum = %d; Prod = %d\n", sum, prod)
	q, p := quadAndPentaple(23)
	fmt.Printf("Quad = %d; Pent = %d\n", q, p)
	/*
		NOTE: now direct assignment q, p := quadAndPentaple() works!
		This is because the return values in 'quadAndPentaple' are named.
	*/
	fmt.Printf("Hex = %d\n", hexaple(23))
	// ==== Type conversions ====
	b := byte('\n')
	fmt.Printf("Value of 'b' = %v; Type of 'b' = %T\n", b, b)
	// Value of 'b' = 10; Type of 'b' = uint8
	a := float32(3)
	fmt.Printf("Value of 'a' = %v\n", a)
	c := 'a'
	fmt.Printf("Type of 'c' = %T\n", c)
	// Type of 'c' = uint32

	// ==== Standard built-in collections ====
	// *** Arrays - fixed size collection
	var a4 [4]int
	fmt.Printf("a4 value = %v; a4 type = %T\n", a4, a4)
	// a4 value = [0 0 0 0]; a4 type = [4]int
	// NOTE: int32 array of 4 values, init to 0

	a5 := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("a5 value = %v; a5 type = %T\n", a5, a5)
	// a5 value = [10 20 30 40 50]; a5 type = [4]int
	/*
		NOTE: int32 array of 5 values, init to [10 20 30 40 50]
		NOTE: the ellipses '...' is required! It differentiates
		it from slices
	*/
	// Arrays are value types
	b5 := a5
	a5[0] = 100
	fmt.Printf("b5 = %v\n", b5)
	// b5 is not affected, it is a copy!

	// *** Slices - variable size collection
	// Similar to Lists in Python
	var s1 []int // declaration only, nothing allocated
	// --- append to slice - built-in function 'append'
	s1 = append(s1, 1, 3, 5, 7)
	fmt.Printf("s1 value = %v; s1 type = %T\n", s1, s1)
	// indexing range of values
	fmt.Printf("s1[0:1] = %v\n", s1[0:1])             // [1]
	fmt.Printf("s1[0:2] = %v\n", s1[0:2])             // [1 3]
	fmt.Printf("s1[0:len(s1)] = %v\n", s1[0:len(s1)]) // [1 3 5 7]
	// full length of the slice using 'len()'
	fmt.Printf("s1[1:] = %v\n", s1[1:]) // [3 5 7]
	// [<low>:] = <low> - till -> <end>
	fmt.Printf("s1[:3] = %v\n", s1[:3]) // [1 3 5]
	// [:<high>] = <0> - till -> <high>

	// --- allocate a slice uisng - make()
	a1 := make([]string, 5)
	// allocate a slice of strings with size 5, inited to ""
	fmt.Printf("a1 = %v\n", a1)

	// --- copy - copy (destination <- source)
	o1 := []int{1, 2, 3, 4, 5}
	e1 := []int{10, 20, 30, 40}
	copy(o1, e1)
	fmt.Printf("o1 = %v\n", o1) // [10 20 30 40 5]
	// NOTE: Any overflow from source will be ignored
	// --- copy - with sub-range
	o2 := []int{1, 2, 3, 4, 5}
	copy(o2[1:4], e1)
	fmt.Printf("o2 = %v\n", o2) // [1 10 20 30 5]
	// NOTE: elements at indices 1, 2, 3 are replaced by e1
	copy(o2[1:4], e1[1:])
	fmt.Printf("o2 = %v\n", o2) // [1 20 30 40 5]
	copy(o2[1:4], e1[2:])
	fmt.Printf("o2 = %v\n", o2) // [1 30 40 30 5]
	// NOTE: source slice [30, 40] copied cyclically!

	// --- delete from slice - fast - order not preserved
	i := 1
	s1[i] = s1[len(s1)-1]                             // copy last element to position 'i'
	s1[len(s1)-1] = 0                                 // "zero" last element
	s1 = s1[:len(s1)-1]                               // truncate the slice without last element
	fmt.Printf("s1 with 2nd item deleted = %v\n", s1) // [1 7 5]
	// NOTE: This has constant time complexity

	// --- delete from slice - slow - order preserved
	s2 := []int{1, 2, 3, 4, 5}
	copy(s2[i:], s2[i+1:])
	// copied over everything i+1 to one place left
	s2[len(s2)-1] = 0
	// reset last value as it is redundant now
	s2 = s2[:len(s2)-1]
	// truncate slice without last element
	fmt.Printf("s2 with 2nd item deleted = %v\n", s2) // [1 3 4 5]
	// NOTE: This has linear time complexity

	// *** Maps - variable size associative arrays
	/*
		Similar to Dictionaries/Hash Tables
		Maps in Go has the structure
		map[KeyType]ValueType
	*/
	// --- creating maps ---
	// declare a map variable
	var scores map[string]int // now this variable points to 'nil'
	// allocate and initialize map data
	scores = make(map[string]int)

	// declare and initialize together - {key: value, ..}
	days := map[string]int{"sun": 0, "mon": 1, "tue": 2, "wed": 3, "thu": 4, "fri": 5, "sat": 6}

	// --- adding elements to a map - assign with key - like JS ---
	scores["Alan"] = 83
	scores["Bob"] = 72
	scores["Cathy"] = 91

	// --- accessing a value - [key] ---
	var sr int
	var found bool
	sr, found = scores["Bob"]
	fmt.Printf("Bob's score = %d; found = %v\n", sr, found)
	// Bob's score = 72; found = true
	sr, found = scores["Ron"]
	fmt.Printf("Ron's score = %d; found = %v\n", sr, found)
	// Ron's score = 0; found = false
	/*
		NOTE: In idiomatic Go style, accessing a map element is
		returns two values. The second is a boolean value that
		indicates if the key was found, and the first is the
		actual value. If the value does not exist, it will have
		a 'zero' value.
		To ignore a returned value use '_'
	*/
	// --- number of items - len() ---
	fmt.Printf("Num of days = %d\n", len(days))

	// --- delete from a map - delete() ---
	delete(scores, "Bob")
	fmt.Println(scores)
	// map[Alan: 83 Cathy: 91]
	// NOTE: If the key is not found, 'delete' does nothing

	// ==== Control-flow commands ====
	// *** conditionals
	// --- if / else ---
	if 2 == 3 {
		// NOTE: the 'condition' does not need ()
		// The body requires {}
		fmt.Println("Inside '2 == 3'")
	} else if 2 == 2.0 {
		fmt.Println("Inside 2 == 2.0")
	} else {
		// NOTE: 'else' has to be inline with the } .. {
		fmt.Println("Inside 'else'")
	}
	// --- if with initialization! ---
	if i1, i2 := 2.0*22/7, 3.414*2; i1 > i2 {
		fmt.Printf("%v > %v\n", i1, i2)
	} else {
		fmt.Printf("%v > %v\n", i2, i1)
	}
	//6.828 > 6.285714285714286
	// --- switch / case ---
	sw1 := 20
	var r1 string
	switch sw1 {
	case 0:
		r1 = "Zero"
	case 10:
		r1 = "Too low"
	case 20:
		r1 = "Just right"
	case 30:
		r1 = "Too high"
	default:
		r1 = "Undefined"
	}
	fmt.Println(r1)
	// NOTE: Switch in Go has no break!
	// switch with expression cases
	switch {
	case sw1 < 20:
		r1 = "Below 20"
	case sw1 == 20:
		r1 = "Equal 20"
	default:
		r1 = "Above 20"
	}
	fmt.Println(r1)
	switch {
	case sw1 >= 10:
		r1 = "At 10"
	case sw1 >= 20:
		r1 = "At 20"
	case sw1 >= 30:
		r1 = "At 30"
	default:
		r1 = "Out of range"
	}
	fmt.Println(r1)
	// At 10
	/*
		NOTE: Switch in Go has no fall-through, which is why in the
		above example even though sw1=20, the first case condition
		is satisfied and then it ignores everything else!
	*/
	// *** Iteration
	// --- for loop ---
	/*
		The general syntax is -
		for <initialization>; <condition>; <post>{
			<body>
		}
	*/
	fmt.Println()
	for sw1 := 1; sw1 < 10; sw1++ {
		fmt.Printf("%d ", sw1)
	}
	fmt.Println()
	//1 2 3 4 5 6 7 8 9
	/*
		NOTE:  The variable 'sw1' in the for loop is different from the one
		outside. It is in the scope of the for loop only, which is why
		we can do sw1 := 1 and not sw1 = 1
	*/
	// --- initialization & post can be separate
	sw1 = 1
	for sw1 < 10 {
		fmt.Printf("%d ", sw1)
		sw1++
	}
	fmt.Println()
	// --- 'break' and 'continue'
	sw1 = 0
	for {
		sw1++ // inc loop variable
		fmt.Printf("%d ", sw1)
		if sw1 >= 10 {
			break // exit loop
		}
		if sw1 > 5 {
			continue // skip the rest
		}
		sw1++ // inc loop variable again!
	}
	fmt.Println()
	// 1 3 5 7 8 9 10
	// --- multiple variables
	for i, j := 1, 10; i <= 10 || j <= 30; i, j = i+1, j+10 {
		fmt.Printf("(%d, %d)", i, j)
	}
	fmt.Println()
	/*
		NOTE:The loop will execute till the 'condition' becomes false
		 => in this case loop will terminate only when (i > 10) AND (j > 30)
		 => we get (1, 10)(2, 20)(3, 30)(4, 40)(5, 50)(6, 60)(7, 70)(8, 80)(9, 90)(10, 100)
		 if the loop condition was (i <= 10) && (j <= 30)
		 => the loop will terminate if either (i > 10) OR (j > 30)
		 => we get (1, 10)(2, 20)(3, 30)

		 Note also the use of comma (,) for separating the variables!
	*/
	// --- 'range' to iterate over collections
	for i, v := range [...]rune{'A', 'B', 'C', 'D', 'E'} {
		fmt.Printf("%d:%c ", i, v)
	}
	fmt.Println()
	// 0:A 1:B 2:C 3:D 4:E
	// NOTE: 'range' returns an index and a value!
	// --- 'range' over map
	for k, v := range map[int]rune{1: 'A', 2: 'E', 3: 'I', 4: 'O', 5: 'U'} {
		fmt.Printf("(%d = %c) ", k, v)
	}
	fmt.Println()
	// (4 = O) (5 = U) (1 = A) (2 = E) (3 = I)
	// NOTE: Order is not preserved for maps

	// === Place-holder identifier ====
	s3, _ := sumAndProd(23, 23)
	fmt.Printf("s3 (sum only) = %d\n", s3)
	// NOTE: can be used to ignore some return values
	// sometimes used to bypass unused variable check !
	i1 := 2
	_ = i1

	// === Variadic functions - Invocation ===
	fmt.Println(fullName("John", "Doe"))
	// John Doe
	fmt.Println(fullName("Jon", "Von", "Neumann"))
	// Jon Von Neumann
	// NOTE: variable number of names passed in
	showVar(1, 20, 45, 34, 78)
	// Type of varidic argument 'prm' = []int
	/*
		It is converted to a 'Slice of type int' inside
		the function, however we cannot directly pass
		in a 'Slice of int' as an argument here!
	*/

	// --- Passing a Slice as a variadic arument ---
	/*
		Since a variadic argument is cast into a 'slice' at runtime,
		one might assume that it would be alright to pass in a 'slice'
		as a variadic argument...
	*/
	// fullName([]string{"The", "ghost", "who", "walks"})
	// the above will result in a compile time type error
	/*
		We have to be cognizant that 'variadic function parameter'
		and 'Slice parameter' are two different type signatures

		However if we "suffix the slice with ..." then the compiler
		accepts it. The slice will be directly passed in without a
		new slice being created!
	*/
	nm := []string{"The", "ghost", "who", "walks"}
	fmt.Println(fullName(nm...))
	// --- Gotcha - Note that the 'slice' can get modified ---
	fmt.Println(nm) // [The ghost who walks]
	change(nm...)
	fmt.Println(nm) // [Modified! ghost who walks]
	// --- Variable type of argument ----
	/*
		This can be achived using (empty) 'interface'.
		More on this later after introducing interfaces!

		for example the 'Printf' function in Go lib 'fmt' is
		declared as follows -

		func Printf(format string, a ...interface{}) (n int, err error){
			...
		}
	*/
	// === First-class functions ===
	/*
		Functions in Go are first-class citizens just like
		any other datatypes such as string or int.
		That is, functions can be treated as values that
		can be -
			1) assigned to variables
			2) passed in as arguments to functions
			3) returned from other functions
			The functions that accept other functions and/or
			return functions are called Higher Order Functions (HOF)
	*/
	// --- function literals & closures ---
	circClosed := false
	switchAction := func() {
		if circClosed {
			fmt.Println("The ligt is ON!")
		} else {
			fmt.Println("Light is OFF.")
		}
	}
	switchAction() // Light is OFF.
	circClosed = true
	switchAction() // The light is ON!
	/*
		NOTE: The function literal is a 'Closure'. The referenced outer variable
		'circClosed' is wrapped up in the closure.
		Changing the variable value affets the closure functionality!
		NOTE: Also how we used an 'anonymous' function!
	*/
	// --- HOF - Passing functions as arguments
	fmt.Println(calc(2, 3, func(x, y int) int { return x + y }))
	// 2 + 3 = 5
	fmt.Println(calc(2, 3, func(x, y int) int { return x * y }))
	// 2 * 3 = 6
	// NOTE: The behaviour is injected

	// --- HOF - Returning functions from HOF / function factory
	c1 := counterFact() // counter 1
	c2 := counterFact() // counter 2
	/*
		NOTE: c1 & c2 are closures over the counter
		variable 'i'
	*/
	fmt.Printf("counter 1 - value = %d\n", c1())
	// counter 1 - value = 1
	fmt.Printf("counter 2 - value = %d\n", c2())
	// counter 2 - value = 1
}

// ==== Function declaration ====
/*
	Function declaration uses the keyword 'func'.
	Like other lamguages, they take paramters in parantheses '()',
	and return values using 'return' keyword.
*/
func printDouble(x int) {
	fmt.Printf("Double of %d = %d\n", x, 2*x)
}

/*
NOTE: If a function returns a value it has to specify
the return type.
*/
func triple(x int) int {
	return 3 * x
}

/*
Go functions can return more than one value, in fact it
is idiomatic in Go to use this for returning 'errors'
from functions as there are no exceptions.
*/
func sumAndProd(x, y int) (int, int) {
	return x + y, x * y
}

/*
Unlike most other languages, the return values can be
named!
*/
func quadAndPentaple(x int) (quad, pent int) {
	return x * 4, x * 5
	// quad = x * 4
	// pent = x * 5
}

/*
The named return values can be assigned to, and then 'return' will
automatically return these.
*/
func hexaple(x int) (r int) {
	r = 6 * x
	return
	// value assigned to 'r' will be returned
	// 'r' is implicit here
}

// --- Variadic functions ---
/*
Variadic functions accept a variable number of parameters.
They are useful when the number of arguments cannot be pre-determined.
For a function to be variadic, the "last" parameter should be
of the form (...T).
This will imply that the function can take any number of
arguments of the type 'T'

For "variable type variadic parameters" (such as used by Println, Printf etc.),
we have to rely on the empty interface 'interface{}'. More on this later!
*/
func fullName(names ...string) string {
	var buffer bytes.Buffer
	l := len(names) - 1
	for i, n := range names {
		buffer.WriteString(n)
		if i != l {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}

// show variadic argumenyt type
func showVar(prm ...int) {
	fmt.Printf("Type of varidic argument 'prm' = %T\n", prm)
	// prm becomes a new Slice within the function
}

// 'slice' passed in as variadic args can get modified
func change(text ...string) {
	if len(text) > 0 {
		text[0] = "Modified!"
	}
}

// --- HOF - function type parameter ---
func calc(x int, y int, oper func(int, int) int) int {
	return oper(x, y)
}

// --- User Defined Function Type! ---
type counter func() int

// --- HOF - returning function / function factory ---
func counterFact() counter {
	i := 0
	return func() int {
		i++
		return i
	}
}
