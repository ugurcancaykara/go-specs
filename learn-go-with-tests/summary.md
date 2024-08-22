# Summary



### Section 1 - Integers 

- Written tests at `01integers` folder path

Wrapping up

- More practice of the TDD workflow
- Integers, addition
- Writing better documentation so users of our code can understand its usage quickly
- Examples of how to use our code, which are checked as part of our tests

### Section 2 - Iteration

- Written tests at `02iterators` folder path

Wrapping up

- More TDD practice
- Learned `for`
- Learned how to write benchmarks

### Section 3 - Arrays and slices

- Written tests at `03arraysandslices` folder path

Wrapping up

- Arrays
- Slices
	- The various ways to make them
	- How they have a fixed capacity but you can create new slices from old ones using append
	- How to slice, slices!
- `len` to get the length of an array or slice
- Test coverage tool
- reflect.DeepEqual and why it's useful but can reduce the type-safety of your code

I've used slices and arrays with integers but they work with any other type too, including
arrays/slices themsleves. So you can declare a variable of `[][]string` if you need to.


### Section 4 - Structs and Methods and Interfaces

- Written tests at `04structsmethodsinterfaces` folder path

Wrapping up

This was more TDD practice, iterating over our solutions to basic mathematic problems and learning new language features motivated by our tests.
- Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
- Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
- Adding methods so you can add functionality to your data types and so you can implement interfaces
- Table driven tests to make your assertions clearer and your test suites easier to extend & maintain


### Section 5 - Pointers and Errors

- Written tests at `05pointersanderros` folder path

Wrapping up

- Pointers
	- Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
	- The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

- nil
	- Pointers can be nil
	- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
	- Useful for when you want to describe a value that could be missing


- Errors
	- Errors are the way to signify failure when calling a function/method.
	- By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
 

- Create new types from existing ones
	- Useful for adding more domain specific meaning to values
	- Can let you implement interfaces
