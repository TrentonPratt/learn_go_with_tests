# learn_tests_with_go
######Learning golang by following https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go

##Notes:

t.Helper() should be added to helper functions for tests. It will cause a failing test to show the line number of the
main testing function rather than the helper function

Basic cycle discipline:
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

##Useful commands:

Navigate to repo:
```
cd "C:\Users\Trenton Pratt\Work\Practice\learn_go_with_tests"
```
Go doc:
```
cd "C:\Users\Trenton Pratt\go\bin"
godoc -http :8000
```
