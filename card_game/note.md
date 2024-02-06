1. ##### Code

    ```go
    package main
    
    import "fmt"
    
    func main() {
        var card string = "Ace of Spades"
        fmt.Println(card)
    }
    ```

2. ##### Syntax

    ```go
    var card string = "Ace of Spades"
    ```

    1. `string`: basic data type
    2. `bool`: true or false
    3. `int`: numbers 0, -1000, 9999
    4. `float64`: 10.00001, 0.0009, -100.003
    5. need to give data type to variable

    ```go
    card := "Ace of Spades"
    ```

    the code above, would do the same job as the previous syntax. Go will interfer data type based on the right-hand value.

    Don't have to change use `:=` if we just want to re-assign a variable. Otherwise, if we need initilize vairable, then we can use it.