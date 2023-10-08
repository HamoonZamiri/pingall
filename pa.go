package main
import (
	"fmt"
	"bufio"
	"os"
	"net/http"
)

func healthCheck(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println(url, res.Status)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pa <file>")
		return
	}

 	fileName := os.Args[1]
	file, err := os.Open(fileName)
	var urls []string

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for _, url := range urls {
		healthCheck(url)
	}

}