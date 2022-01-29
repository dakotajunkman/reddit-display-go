package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	"strings"
)

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

	var again string
	fmt.Print("Do you want to try again? y/n ")
	fmt.Scanln(&again)

	return strings.ToLower(again) == "y"
}