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






########## Data Structs and Memory ########## :-



Arrays in Go have fixed size, and can be declared empty or have initialized directly.

Slices are abstarctions over arrays that dont need a size to be specified, and adding elements beyond their capacity doubles their size everytime.

Zero value of a map is nil.

A pointer is a variable that holds the address of another variable; in Go, pointer arithmetix is not supported unlike C and C++.






########## Functions and Error Handling ########## :-



Functions are used to group sequence of statements to perform a task, made using func keyword.

Variadic functions can be used if one is unsure about the number of parameters to take as input.

Functions can return multiple values, including errors.

Custom error types can be created with structs.

Defer makes the statement execute at end....and between two defers, the later one is executed first as per lifo behavior.

Panic is used to crash the program and recover called with defer when a panic occurs.






########## Object Oriented Programming ########## :-



Structs (structures) can be used to define custom concrete types having multiple primitive types inside them.

Methods are funtions but with a receiver that helps better implement encapsulation.

Interfaces specify 'what' (not how) a type can do help define a method signature without providing their implementation. If a type implements the method/s of the interface, it implicitly satisfies it. 

The Stringer interface can be used to improve readability of output by providing a custom alternative to standard printing format.

Generics help provide flexibility over types by allowing a function or type to accept multiple types of data.






########## Composition And Design Logic ########### :-



Composition builds larger types by combining smaller reusable types ("has-a" relationship) instead of inheritance.

Embedding is composition without an explicit field name, promoting the embedded type's fields and methods for direct access.

Fields and methods of an embedded type can be accessed directly through the outer struct as if they belonged to it.

If the outer struct defines a field or method with the same name as an embedded one, the outer definition takes precedence.






########## String Processing And Text ########## :-



strings.Builder → Efficiently builds large strings without repeated allocations.
TrimSpace() → Removes leading and trailing whitespace.
Split() ↔ Join() → Convert between strings and slices.
Fields() → Splits a string on whitespace.
HasPrefix()/HasSuffix() → Checks the beginning/end of a string.
Replace()/ReplaceAll() → Replaces substrings.

fmt.Sprintf() → Returns a formatted string.
Formatting Verbs:
%s  string
%d  integer
%f  float
%t  boolean
%v  default value
%+v struct with field names
%#v Go-syntax representation
fmt.Errorf() → Creates or wraps formatted errors (%w preserves the original error).

Strings in Go are UTF-8 encoded (variable-length encoding).
len(string) returns bytes, not characters (runes).
A rune represents one Unicode code point (character).
Use 'for range' to iterate over characters safely.

Regex → Pattern matching for searching, validating and extracting text.
Compile() / MustCompile() → Create regex patterns.
MatchString() → Match check
FindString() → First match
FindAllString() → All matches

Templates → Generate dynamic text using placeholders.
Parse() → Compile template.
Execute() → Render template with data.

bufio.Scanner → Reads input line-by-line.