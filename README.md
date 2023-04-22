# GoLang-Tutorials
golang is compiled , it  compiles everything and runtime is baked into the compiled code

Features of golang :
 1) Go does not have generics, Generics allow you to write code that can work with different types of data, without having to write separate code for each type.
 2) Go does not have support for exception handling
 3) Go does not support traditional inheritance, which is a common feature in object-oriented programming languages.
 4) Go does not support function overloading
 5) first-class support for concurrency and parallelism
 6) It has a compiler that produces efficient machine code, and its garbage collector is optimized for low latency and high throughput and garbage collection happens automatically
 7) Go provides memory safety features that help prevent common programming errors, such as buffer overflows and null pointer dereferences.
 8) Go is a statically-typed language, which means that variable types are checked at compile-time.
 9) well-suited for writing networked and distributed systems
 10) Instead of classes we have structs in golang
 11) one entry point i.e, main function 
 12) everything considered as types in golang , ex: function is types in golang so we can also pass it on another function
 13) Memory allocationand deallocation happens automatically
 14) The garbage collector in Go uses a mark-and-sweep algorithm to perform garbage collection. The algorithm works by marking all memory that is currently in use, then sweeping through the memory to identify and free up memory that is not marked. Garbage collection is a memory management technique that automatically frees up memory that is no longer in use by the program. 
 15) array is very less used and slices is mostly used in golang
 16) In the root directory there can be only 1 main.go file



 ** lexer
 In programming, a lexer(lexical analyzer) is a tool that reads your code and breaks it down into individual "tokens" or "lexemes". Each token represents a meaningful unit of your code, such as a keyword, identifier, or symbol. The lexer then passes these tokens on to the next step of the compilation process.A lexer breaks down your code into individual tokens to help the computer understand it better. This is an important step in the process of turning your code into a program that the computer can execute.

 ** 
  Use GOOS from env to build for specific OS