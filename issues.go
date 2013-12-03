package main

type Location struct {
	File   string `xml:"file,attr"`
	Line   int    `xml:"line,attr"`
	Column int    `xml:"column,attr"`
}

type Issue struct {
	Id          string     `xml:"id,attr"`
	Severity    string     `xml:"severity,attr"`
	Message     string     `xml:"message,attr"`
	Category    string     `xml:"category,attr"`
	Priority    int        `xml:"priority,attr"`
	Summary     string     `xml:"summary,attr"`
	Explanation string     `xml:"explanation,attr"`
	ErrorLine1  string     `xml:"errorLine1,attr"`
	ErrorLine2  string     `xml:"errorLine2,attr"`
	Url         string     `xml:"url,attr"`
	Urls        string     `xml:"urls,attr"`
	QuickFix    string     `xml:"quickfix,attr"`
	Locations   []Location `xml:"location"`
}

type LintReport struct {
	Format string  `xml:"format,attr"`
	By     string  `xml:"by,attr"`
	Issues []Issue `xml:"issue"`
}
