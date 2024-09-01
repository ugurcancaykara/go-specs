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


### Section 6 - Maps

- Written tests at `06maps/`

Wrapping up

In this section, we covered a lot. We made a full CRUD (Create, Read, Update and Delete) API for our dictionary. Throughout the process we learned how to:

- Create maps
- Search for items in maps
- Add new items to maps
- Update items in maps
- Delete items from a map
- Learned more about errors
	- How to create errors that are constants
	- Writing error wrappers

### Section 7 - Dependency Injection

- Written tests at `07dependencyinjection/`

Wrapping up

Our first round of code was not easy to test because it wrote data to somewhere we couldn't control
```
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```

Motivated by written tests we refactored the code so we could control where the data was written by injecting a dependency which allowed us to:

- Test our code If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.
- Separate our concerns, decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilities (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.
- Allow our code to be re-used in different contexts The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.


#### What about mocking? I hear you need that for DI and also it's evil
- Mocking will be covered in detail at next section `08mocking/`
You use mocking to replace real things you inject with a pretend version that you can control and inspect in your tests. In our case though, the standard library had something ready for us to use.


### Section 8 - Mocking

- Written tests at `08mocking/`

Wrapping up


- Let's say we have a CountDown() functions which counts down from 3, printing each number on a new line. While this is a pretty trivial program, to test it fully we will need as always to take an iterative, test-driven approach.

**What do I mean by iterative? We make sure we take the smallest steps we can to have useful software.**

- We don't want to spend a long time with code that will theoretically work after some hacking because that's often how developers fall down rabbit holes. It's an important skill to be able to slice up requirements as small as you can so you can have working software.






