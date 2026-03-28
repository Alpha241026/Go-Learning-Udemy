########## General ########## :-


Every go executable is part of a main package...and this is where the program execution begins from a main function

Use go build to compile your source code into a single static binary and go run to run it.

You can't have more than one go files in a single folder as they'll be treated as the same package...and you can't have multiple main functions in a single package.

In a Go program, you must use every library that you import.




########## Core Language Fundamentals section ########## :-


The type of a variable is usually meentioned after its name when declaring.

A zero-value is assigned to variables that have been declared but not initialized.

Multiple variables can be delared on the same line

You can also declare and initialize on the same line without specifying datatype using type inference...wherein you let the compiler infer the type from the value you enter on the right

Constants are declared using the const keyword and can be of numeric,string or boolean types.

Enums are a set of constant values implemented using const and iota keywords in which each value is initialized by default in a successive auto-incremented way




########## Control Flow & Logic section ########## :-


For loop is the only way to loop in Go. It is quite versatile...can be used as classic C-style, a while loop, infinite loop (preferably with jump statements like break or continue) or as a range over arrays.

If (standalone) and else (also after else if) can be used as one kind of conditional statements and 
switch (having cases (separated by jump statements or fall through) and an optional default) as other.