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

3. ##### Functions

    ```go
    func newCard() string {
    	return "Five of Diamonds"
    }
    ```

    1. `string` is the return type of the function. We need to declare it expiclitly in Golang
    2. `newCard` starts from lower case, which means it's a private function
    3. `return` end execution and return the value of computation

4. ##### Slice and Array

    Array: fixed length list of things

    Slice: an array that can grow and shrink

    1. need to have same type

    2. ```go
        var cards []string = []string {"Ace of Diamonds", "Ace of Spades"}
        cards = append(cards, "Six of Spades")
        ```

        append will not change the slice, instead, it will return a new slice

    3. How to iterator slice

        ```go
        for index, card := range cards {
            fmt.Println(index, card)
        }
        ```

        `index`: the index of the element that current iterating

        `card`: the content of the element

        `range`: `i` and `card` are initialized everytime it starts iterating, the previous one will be  dropped after iteration.