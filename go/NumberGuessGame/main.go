package main
import(
"bufio"
"fmt"
"math/rand"
"os"
"strconv"
"strings"
"time"
)

func main() {
rand.Seed(time.Now().UnixNano())
secretNo := rand.Intn(100)+1
attempts := 3


fmt.Println("**************  Welcome to the Number Guessing Game  **************")
fmt.Println("I have picked a number between 1 and 100. Can you guess it ?")

reader := bufio.NewReader(os.Stdin)
for(attempts > 0){
	attempts--
	fmt.Println("Enter your guess: ")

	input, err := reader.ReadString('\n')

	if err != nil{
		fmt.Println("Error reading your input")
		continue
	}
	guess, err := strconv.Atoi(strings.TrimSpace(input))

	if guess < secretNo{
		fmt.Println("Too low! Try again. You have left with ",attempts," attempts")
		continue
	}else if guess > secretNo{
		fmt.Println("Too high! Try again. You have left with ",attempts," attempts")
		continue
	}else{
	fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
	break
	}
}

fmt.Printf("Game Over! The secret number was %d.\n",secretNo)


}