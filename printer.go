package main

import (
	"os"
	"strconv"

	"github.com/stangirard/yatas/plugins/commons"
)

func WriteMarkdown(tests []commons.Tests) {
	// Write markdown report
	// Open the file for writing
	file, err := os.Create("report.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the header
	file.WriteString("# Yatas \n")

	// Write the checks in a table

	for _, test := range tests {
		file.WriteString("\n## " + test.Account + "\n\n")
		file.WriteString("| Id | Name | Status | \n")
		file.WriteString("| ---- | ---- | ------ | \n")
		for _, check := range test.Checks {
			status := "✅"
			if check.Status == "FAIL" {
				status = "❌"
			}

			file.WriteString("| " + check.Id + "| " + check.Name + " | " + status + " | \n")
		}
		file.WriteString("\n")

		WriteCategoriesSuccess(test, file)
	}
}

func WriteCategoriesSuccess(test commons.Tests, file *os.File) {
	// Find the categories
	categories := []string{}
	categoriesSuccess := map[string]int{}
	categoriesFailure := map[string]int{}
	for _, check := range test.Checks {
		for _, category := range check.Categories {
			if !contains(categories, category) {
				categories = append(categories, category)
				categoriesSuccess[category] = 0
				categoriesFailure[category] = 0
			}
			if check.Status == "OK" {
				categoriesSuccess[category]++
			} else {
				categoriesFailure[category]++
			}
		}
	}

	// Write the categories
	file.WriteString("\n## Categories \n\n")
	file.WriteString("| Category | % Completion | \n")
	file.WriteString("| ---- | ---- | \n")
	for _, category := range categories {
		file.WriteString("| " + category + "| " + CalculatePercent(categoriesSuccess[category], categoriesFailure[category]) + "% | \n")
	}
	file.WriteString("\n")

}

func CalculatePercent(success int, failure int) string {
	total := success + failure
	if total == 0 {
		return "0"
	}
	return strconv.Itoa((success * 100) / total)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
