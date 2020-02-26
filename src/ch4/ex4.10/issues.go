package main

import (
	"ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	monthInHours = 24 * 30
	yearInHours  = 24 * 365
)

func printIssues(items []*github.Issue) {
	for _, item := range items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var issuesLessThanOneMonth, issuesLessThanOneYear, issuesMoreThanOneYear []*github.Issue
	for _, item := range result.Items {
		hours := time.Since(item.CreatedAt).Hours()
		if hours < monthInHours {
			issuesLessThanOneMonth = append(issuesLessThanOneMonth, item)
		} else if hours < yearInHours {
			issuesLessThanOneYear = append(issuesLessThanOneYear, item)
		} else {
			issuesMoreThanOneYear = append(issuesMoreThanOneYear, item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("%d less than 1 month old\n", len(issuesLessThanOneMonth))
	printIssues(issuesLessThanOneMonth)
	fmt.Printf("%d more than 1 month but less than 1 year old\n", len(issuesLessThanOneYear))
	printIssues(issuesLessThanOneYear)
	fmt.Printf("%d more than 1 year old\n", len(issuesMoreThanOneYear))
	printIssues(issuesMoreThanOneYear)
}
