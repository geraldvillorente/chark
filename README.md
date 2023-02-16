# chark (char-kee)
Minimalist command-line notes manager for Mac OS users. 

## Unit tests

To run the unit tests for this code, you can use the `go test` command in your terminal or command prompt.

First, navigate to the directory where your Go code and test file are located. Then, run the following command:

```
go test
```

This will run all of the test functions in the test file and report the results in the terminal. You should see output similar to the following:

```
PASS
ok      your/package/path    0.013s
```

If any of the tests fail, you will see an error message that describes the problem.

You can also run specific test functions by specifying their names as arguments to the `go test` command. For example, to run only the `TestAddReminder()` test function, you can run the following command:

```
go test -run TestAddReminder
```

This will only run the specified test function and will skip all of the other tests.

