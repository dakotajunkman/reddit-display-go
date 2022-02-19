package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"encoding/json"
)

// define each post
type post struct {
	Title string
	Username string
	Url string
}

// JSON is returned as data pointing to a slice of posts
type data struct {
	Data []post
}

func main() {
	printIntro()
	usernamePrompt()
	username := getInput()

	if len(username) == 0 {
		meanKeywordPrompt()
	} else {
		skippableKeywordPrompt()
	}

	keyword := getInput()
	
	if len(keyword) == 0 && len(username) == 0 {
		beMean()
		if tryAgain(){
			main()
			return  // stops execution from continuing after recursion
		}
		return
	}

	data := httpReq(username, keyword)
	displayPosts(data.Data)

	if tryAgain() {
		main()  // no need to stop since this is last line of function
	}
}

// introduce program
func printIntro() {
	fmt.Println("This neat program allows you to filter posts to the r/OSUOnlineCS subreddit by " +
	"username or keyword, just like you've always wanted.\n You must enter at least one search " +
	"parameter, but can do both if you want.\n If you enter none you will get in trouble and have " +
	"to do it again. To skip a parameter, simply hit Enter.")
}

// ask for username
func usernamePrompt() {
	fmt.Print("Username (can be skipped) (Press Enter to submit): ")
}

// get input from user and strip off whitespace
func getInput() string {
	var input string
	fmt.Scanln(&input)

	// trim off any whitespace
	return strings.TrimSpace(input)
}

// prompt for keyword when username entered
func skippableKeywordPrompt() {
	fmt.Print("Keyword (can be skipped) (Press Enter to submit): ")
}

// prompt for keyword when username was skipped
func meanKeywordPrompt() {
	fmt.Print("Keyword (skip at your own risk) " +
	"(Press Enter to submit): ")
}

// say mean thing to user
func beMean() {
	fmt.Println("Not cool. I specifically told you not to " +
	"leave both fields blank.")
}

// find out if user wants to run program again
func tryAgain() bool {
	fmt.Print("Do you want to try again? y/n ")
	userAnswer := getInput()

	return strings.ToLower(userAnswer) == "y"
}

// call Reddit scraper and extract response body
func httpReq(username string, keyword string) data {
	url := fmt.Sprintf("https://parser-service-361.herokuapp.com/reddit/get?username=%s&keyword=%s", username, keyword)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return decodeJson(body)
}

// unpack the JSON into a slice of structs
func decodeJson(body []byte) data {
	var data data

	// unmarshal will error on an empty byte array
	if len(body) == 0 {
		return data
	}

	// unpack the JSON into a slice of structs
	err := json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

// show posts to the user
func displayPosts(posts []post) {
	if len(posts) == 0 {
		fmt.Println("Your search has no results. Darn!")

	} else {
		fmt.Printf("\nYour search yielded %d result(s).\n", len(posts))
		for _, post := range posts {
			fmt.Printf("Title: %s\n", post.Title)
			fmt.Printf("Username: %s\n", post.Username)
			fmt.Printf("URL: %s\n\n", post.Url)
		}
	}
}