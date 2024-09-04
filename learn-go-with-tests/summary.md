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



- Let's say we have a CountDown() functions which counts down from 3, printing each number on a new line. While this is a pretty trivial program, to test it fully we will need as always to take an iterative, test-driven approach.

**What do I mean by iterative? We make sure we take the smallest steps we can to have useful software.**

- We don't want to spend a long time with code that will theoretically work after some hacking because that's often how developers fall down rabbit holes. It's an important skill to be able to slice up requirements as small as you can so you can have working software.


#### Countdown example

Let's say tests pass and software works as intended but we have some problems:

- Our tests take 3 seconds to run.

	- Every forward-thinking post about software development emphasises the importance of quick feedback loops.

	- Slow tests ruin developer productivity.

	- Imagine if the requirements get more sophisticated warranting more tests. Are we happy with 3s added to the test run for every new test of Countdown?

We have not tested an important property of our function.

We have a dependency on Sleeping which we need to extract so we can then control it in our tests.

If we can mock time.Sleep we can use dependency injection to use it instead of a "real" time.Sleep and then we can spy on the calls to make assertions on them.

### But isn't mocking evil ?

You may have heard mocking is evil. Just like anything in software development it can be used for evil, just like DRY.

People normally get in to a bad state when they don't listen to their tests and are not respecting the refactoring stage.


If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should listen to that bad feeling and think about your code. Usually it is a sign of

- The thing you are testing is having to do too many things (because it has too many dependencies to mock)
	- Break the module apart so it does less

- Its dependencies are too fine-grained
	- Think about how you can consolidate some of these dependencies into one meaningful module

- Your test is too concerned with implementation details
	- Favour testing expected behaviour rather than the implementation

Normally a lot of mocking points to bad abstraction in your code.

What people see here is a weakness in TDD but it is actually a strength, more often than not poor test code is a result of bad design or put more nicely, well-designed code is easy to test.


### But mocks and tests are still making my life hard!

Ever run into this situation?

- You want to do some refactoring

- To do this you end up changing lots of tests

- You question TDD and make a post on Medium titled "Mocking considered harmful"

This is usually a sign of you testing too much implementation detail. Try to make it so your tests are testing useful behaviour unless the implementation is really important to how the system runs.

It is sometimes hard to know what level to test exactly but here are some thought processes and rules to follow:

- **The definition of refactoring is that the code changes but the behaviour stays the same.** If you have decided to do some refactoring in theory you should be able to make the commit without any test changes. So when writing a test ask yourself

	- Am I testing the behaviour I want, or the implementation details?

	- If I were to refactor this code, would I have to make lots of changes to the tests?

- Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour. Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.

- Feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design

- Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation. **Be sure you actually care about these details if you're going to spy on them**


#### Can't I just use a mocking framework?

Mocking requires no magic and is relatively simple; using a framework can make mocking seem more complicated than it is. We don't use automocking in this chapter so that we get:

- a better understanding of how to mock

- practice implementing interfaces

In collaborative projects there is value in auto-generating mocks. In a team, a mock generation tool codifies consistency around the test doubles. This will avoid inconsistently written test doubles which can translate to inconsistently written tests.

You should only use a mock generator that generates test doubles against an interface. Any tool that overly dictates how tests are written, or that use lots of 'magic', can get in the sea.




### Wrapping Up


#### More on TDD approach

- When faced with less trivial examples, break the problem down into "thin vertical slices". Try to get to a point where you have working software backed by tests as soon as you can, to avoid getting in rabbit holes and taking a "big bang" approach.

- Once you have some working software it should be easier to iterate with small steps until you arrive at the software you need.

`"When to use iterative development? You should use iterative development only on projects that you want to succeed."`

Martin Fowler.


#### Mocking

- **Without mocking important areas of your code will be untested.** In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that can fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.

- Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in **slow feedback loops.**

- By having to spin up a database or a webservice to test something you're likely to have **fragile tests** due to the unreliability of such services.


Once a developer learns about mocking it becomes very easy to over-test every single facet of a system in terms of the way it works rather than what it does. Always be mindful about **the value of your tests** and what impact they would have in future refactoring.

In these tests about mocking we have only covered **Spies**, which are a kind of mock. Mocks are a type of "test double."

[Test Double is a generic term for any case where you replace a production object for testing purposes](https://martinfowler.com/bliki/TestDouble.html)

Under test doubles, there are various types like stubs, spies and indeed mocks! Check out Martin Fowler's post for more detail.









