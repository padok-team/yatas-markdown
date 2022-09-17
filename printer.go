package main

import (
	"os"

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
	}
}
