Basics of writing a module in go and resuing it in a different module

- `hello` module uses `greetings` module here

## pointing to a local module
the greetings module used here is not published, to use module from local dir run
```
go mod edit --replace shivanikanal/greetings=../greetings
```
and sync dependencies by running
```
go mod tidy
```
this generates a pseudo-version number in go.mod for your local dependency

## run
```
cd hello && go run .
```