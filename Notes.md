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



The `strings` package provides efficient utilities for manipulating text.
Common operations include trimming (`TrimSpace()`), splitting/joining (`Split()`, `Join()`, `Fields()`), searching (`HasPrefix()`, `HasSuffix()`), and replacing (`Replace()`, `ReplaceAll()`).

`strings.Builder` efficiently constructs large strings without repeated allocations.

The `fmt` package supports formatted strings and formatted errors through `Sprintf()` and `Errorf()`.

Regular expressions provide pattern matching for validation, searching and extraction.
Compile patterns once with `Compile()` or `MustCompile()`, then reuse methods like `MatchString()` and `FindAllString()`.

Templates generate dynamic text using placeholders.
`Parse()` compiles the template and `Execute()` renders it with data.

bufio.Scanner → Reads input line-by-line.






########## Concurrency Mastery ########### :-



A goroutine is a lightweight concurrent function started using the 'go' keyword.

WaitGroup is a shared counter used to wait for multiple goroutines.
`Add()` increments, `Done()` decrements, and `Wait()` blocks until the counter reaches zero.

A channel is a typed pipe used to safely send data between goroutines.
`channel <- value` sends data, while `value := <-channel` receives it.

Buffered channels store a limited number of values before blocking the sender.
Unbuffered channels require sender and receiver simultaneously; buffered channels allow temporary asynchronous communication.

`close(channel)` signals that no more values will be sent; buffered values can still be received.
`value, ok := <-channel` returns `ok=false` once the channel is closed and empty.

`select` waits for the first channel operation that becomes ready.
`context` coordinates cancellation or timeouts across goroutines, while `time.After()` creates a timeout channel after a specified duration.

A mutex protects shared data by allowing only one goroutine to access a critical section at a time.
Always `Lock()` before accessing shared mutable data and `Unlock()` (usually with `defer`) when finished.






########## File I/O And System Programming ########## :-



Go provides high-level helpers for reading and writing entire files, along with lower-level APIs for more control.
`os.WriteFile()`, `os.ReadFile()`, `os.Open()` and `os.OpenFile()` cover the most common file operations.

The `filepath` package provides OS-independent path manipulation utilities.
Common helpers include `Join()`, `Base()`, `Ext()` and `Clean()`.

Directories can be created recursively or removed together with all their contents.
`os.MkdirAll()` creates missing parents, while `os.RemoveAll()` removes an entire directory tree.

Temporary files and directories are useful for short-lived data and should usually be cleaned up with `defer`.
`os.CreateTemp()` and `os.MkdirTemp()` create unique temporary resources.

The `embed` package bundles files or directories into the compiled executable at build time.
Embedded resources are accessed through the read-only `embed.FS`.






########## Data Encoding & Security ########## :-



JSON marshalling converts Go values into JSON for storage or communication.
Struct tags control field names, while only exported fields participate in encoding.

JSON unmarshalling converts JSON back into Go values and requires a pointer to populate the destination.
Nested JSON maps naturally to nested structs, while `json:"-"` excludes fields from encoding and decoding.

JSON encoders stream encoded data directly to an `io.Writer` such as a file, terminal or HTTP response.
Unlike `Marshal()`, they avoid creating an intermediate `[]byte`.

JSON decoders read and decode JSON directly from an `io.Reader` such as a string, file or network stream.
`json.NewDecoder()` is preferred when processing streamed or continuously read JSON data.

Base64 converts arbitrary binary data into a text-safe representation and can be reversed without losing information.
`EncodeToString()` encodes bytes into Base64, while `DecodeString()` restores the original bytes.