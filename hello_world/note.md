1. ##### Code

    ```go
    package main
    
    import "fmt"
    
    func main() {
        fmt.Println("Hello World!")
    }
    ```

    Very simple go project with only 7 lines of code

1. ##### How to run the code

    ```shell
    go run main.go
    ```

    However, there is more than just ` go run main.go` to run the project

    1. `go build main.go`: compile the file, it will generate compiled file called `main`

        To run the compiled, run `./main` to run the file

    2. `go fmt main.go`: formats all the files in the directory

    3. `go install`: compile and install a package

    4. `go get`: downloads the raw source code of someone else's package

    5. `go test`: runs any tests associated with the current project

2. ##### What does `package main` mean

    package refers as workspace, it can have more than one go files in the package, as long as the package name is the same.

    By default, the main package in golang means execuatable package. Which means that if the program use it as the identity of the executable project. When run the command `go build` then it will generate `main` file, which is executable. Therefore, if we want to have a executable program, we have to have a main package. In the main package, there must be a `main` function for program entry point

    For those package whose name is not `main` are all considered as reusable package, which the code can be reused in large project.  When running `go build` reusable packag, then nothing will be created.

3. ##### What does `import "fmt"` mean

    `fmt` itself is one of the standard libraries provided by golang environment. 

    `import` means build a link or access to `fmt` library so that `main` package can use functions in `fmt` 

4. ##### What does `func` mean

    Tell golang runtime we're going to declare a function

5. ##### How is the `main.go` file organised

    1. Package declaration
    2. import other packages that we need
    3. declare functions, tell go to do something

