package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abondar24/RecipeCountAnalyzer/analyzer"
	"github.com/abondar24/RecipeCountAnalyzer/data"
	"log"
	"os"
)

func main() {
	fileName := flag.String("file", "", "path to fixtures file")
	recipeKeyword := flag.String("recipe", "veggie", "keyword contained in recipe")
	postcode := flag.String("postcode", "10120", "postcode to search")
	deliveryTime := flag.String("time", "10AM - 2PM", "delivery period")

	flag.Parse()

	if *fileName == "" {
		log.Fatal("Error: No file provided")
	}

	if len(*postcode) > 10 {
		log.Fatal("Error: Postcode must be not longer than 10 characters")
	}

	anl := analyzer.NewAnalyzer(fileName, recipeKeyword, postcode, deliveryTime)

	res := anl.Analyze()

	printResult(res)
}

func printResult(result *data.CountResult) {
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	_, _ = fmt.Fprintln(os.Stdout, string(prettyJSON))
}
