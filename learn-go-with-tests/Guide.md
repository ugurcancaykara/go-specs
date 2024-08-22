# Test Guide





#### Sections
- 04structsmethodsinterfaces/
	- This was more TDD practice, iterating over our solutions to basic mathematic problems and learning new language features motivated by our tests.
		- Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
		- Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
		- Adding methods so you can add functionality to your data types and so you can implement interfaces
		- Table driven tests to make your assertions clearer and your test suites easier to extend & maintain


## This is a guide for writing tests
I will collect some basic infos and approachs below this line

#### Writing tests
Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like xxx_test.go

- The test function must start with the word Test

- The test function takes one argument only t *testing.T

- To use the *testing.T type, you need to import "testing", like we did with fmt in the other file

For now, it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail.


#### Discipline
Let's go over the cycle again for TDD
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

Even this may seem tedious but sticking to the feedback loop is important. Not only does it ensure that you have relevant tests, it helps ensure you design good software by refactoring with safety of tests.

By ensuring your tests are fast and setting up your tools so that running tests is simple you can get in to a state of flow when writing your code.

By not writing tests, you are committing to manually checking your code by running your software, which breaks your state of flow. You won't be saving yourself any time, especially in the long run.


#### Benchmarking
Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```


The testing.B gives you access to the cryptically named b.N.

When the benchmark code is executed, it runs b.N times and measures how long it takes.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")

- NOTE by default Benchmarks are run sequentially.


#### Writing just enough test
It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.


Go's built-in testing toolkit features a coverage tool. Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway.

Try running
```
go test -cover 03arraysandslices/*.go
```

you should see
```
PASS
coverage: 100.0% of statements
```



#### Unchecked errors
Whilst the Go compiler helps you a lot, sometimes there are things you can still miss and error handling can sometimes be tricky.
To find them, run the following in a terminal to install `errcheck`, one of many linters available for Go.

```
go install github.com/kisielk/errcheck@latest
```

then, inside the directory with your code run
```
errcheck 05pointersanderros/*.go
```
this is what I got
```
wallet_test.go:17:18: wallet.Withdraw(Bitcoin(10))
```
what this is telling us is that we haven't checked the error being returned on that line of code.
