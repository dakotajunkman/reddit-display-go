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
	Author string
	Url string
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
		if tryAgain(){
			main()
			return  // stops execution from continuing after recursion
		}
		return
	}
}

func printIntro() {
	fmt.Println("This neat program allows you to filter posts to the r/OSUOnlineCS subreddit by " +
	"username or keyword, just like you've always wanted. You must enter at least one search " +
	"parameter, but can do both if you want. If you enter none you will get in trouble and have " +
	"to do it again. To skip a parameter, simply hit Enter.")
}

func usernamePrompt() {
	fmt.Print("Username (Press Enter to submit): ")
}

func getInput() string {
	var input string
	fmt.Scanln(&input)

	// trim off any whitespace
	return strings.TrimSpace(input)
}

func skippableKeywordPrompt() {
	fmt.Print("Keyword (can be skipped) (Press Enter to submit): ")
}

func meanKeywordPrompt() {
	fmt.Print("Keyword (skip at your own risk) " +
	"(Press Enter to submit): ")
}

func tryAgain() bool {
	fmt.Println("Not cool. I specifically told you not to " +
	"leave both fields blank.")

	fmt.Print("Do you want to try again? y/n ")
	userAnswer := getInput()

	return strings.ToLower(userAnswer) == "y"
}

func httpReq(username string, keyword string) []post {
	url := fmt.Sprintf("www.url.com?username=%s&keyword=%s", username, keyword)
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

func decodeJson(body []byte) []post {
	var posts []post

	err := json.Unmarshal(body, &posts)

	if err != nil {
		log.Fatal(err)
	}

	return posts
}