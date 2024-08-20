# Test Guide




## This is a guide for writing tests

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



