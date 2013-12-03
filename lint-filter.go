package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var configPath = flag.String("config", "filterconfig.json", "Path to filter configuration file.")
	var prettyXml = flag.Bool("pretty", false, "If present, XML output will be indented.")
	flag.Parse()

	input := flag.Arg(0)
	output := flag.Arg(1)

	if len(input) == 0 || len(output) == 0 {
		fmt.Print("usage: lint-filter [options] input-file output-file\n\nOptions:\n")
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Input: %s\n", input)
	xmlContent, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		return
	}
	lintReport := LintReport{}
	err = xml.Unmarshal(xmlContent, &lintReport)
	if err != nil {
		fmt.Printf("Error parsing XML: %s\n", err)
		return
	}

	filterConfig := FilterConfig{}
	err = filterConfig.Read(*configPath)
	if err != nil {
		fmt.Printf("Error reading filter configuration: %s\n", err)
		return
	}

	fmt.Printf("lint format: %s\n", lintReport.Format)
	fmt.Printf("lint by: %s\n", lintReport.By)
	fmt.Printf("number of issues: %d\n", len(lintReport.Issues))

	if lintReport.Format != "4" {
		fmt.Println("Unknown report format.")
		return
	}

	filteredIssues := make([]Issue, 0, len(lintReport.Issues))
	for _, i := range lintReport.Issues {
		if filterConfig.Filter(&i) == FilterResultInclude {
			filteredIssues = append(filteredIssues, i)
		}
	}
	fmt.Printf("issues after filter: %d\n", len(filteredIssues))

	lintReport.Issues = filteredIssues

	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error opening output: %s", err)
		return
	}
	outputFile.WriteString(xml.Header)
	encoder := xml.NewEncoder(outputFile)
	if *prettyXml {
		encoder.Indent("", "  ")
	}
	err = encoder.Encode(lintReport)
	if err != nil {
		fmt.Printf("Error writing XML: %s", err)
	}
	outputFile.Close()
}
