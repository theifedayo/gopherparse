# GopherParse

GopherParse is a Go library that provides functionalities for parsing, manipulating, and formatting HTML and XML documents. It is inspired by the popular JavaScript library Cheerio and aims to offer a fast, flexible, and elegant way to work with HTML and XML content in Go.

## Installation
To use GopherParse, you need to have Go installed and set up. Then, you can install the package using the go get command:

```bash
go get github.com/theifedayo/gopherparse
```

## Usage
Import the GopherParse package in your Go code:

```go
import "github.com/theifedayo/gopherparse"
```

### Parsing HTML and XML
You can parse HTML and XML documents by using the LoadHTML and LoadXML functions, respectively:

```go
htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Sample HTML Document</title>
	</head>
	<body>
		<h1>Hello, World!</h1>
		<p>This is a sample HTML document.</p>
	</body>
	</html>
`

gpObj, err := gopherparse.LoadHTML(htmlContent)
if err != nil {
    // Handle error
}

// Now you can work with the GopherParse object.
```

```go
xmlContent := `
	<root>
		<item>Item 1</item>
		<item>Item 2</item>
	</root>
`

gpObj, err := gopherparse.LoadXML(xmlContent)
if err != nil {
    // Handle error
}

// Now you can work with the GopherParse object.
```

### Reading from Files
GopherParse also supports reading and parsing HTML and XML content from files:

```go
htmlFilePath := "path/to/your/html/file.html"
gpObj, err := gopherparse.LoadHTMLFile(htmlFilePath)
if err != nil {
    // Handle error
}

// Now you can work with the GopherParse object.
```

```go
xmlFilePath := "path/to/your/xml/file.xml"
gpObj, err := gopherparse.LoadXMLFile(xmlFilePath)
if err != nil {
    // Handle error
}

// Now you can work with the GopherParse object.
```

### Finding Elements
You can find elements in the parsed document using FindByTag and FindByClass methods:

```go
h1Elements := gpObj.FindByTag("h1")
// Get all h1 elements in the document

pElements := gpObj.FindByClass("paragraph")
// Get all elements with class "paragraph" in the document
```

### Manipulating Elements
You can manipulate elements in the document, such as setting their text content using SetText:

```go
gpObj.SetText("h1", "Updated Heading")
// Set text content of all h1 elements to "Updated Heading"
```

### Rendering the Document
To get the HTML representation of the document, you can use the Render method:

```go
htmlString := gpObj.Render()
// Get the HTML string representation of the document
```

### Contributing
Feel free to contribute to GopherParse by opening issues, submitting pull requests, or suggesting new features. Your contributions are highly appreciated!

### License
This project is licensed under the MIT License - see the LICENSE file for details.

### Roadmap
- [x] Parsing HTML and XML documents.
- [x] Reading and parsing content from HTML and XML files.
- [x] Finding elements by tag name within the document.
- [x] Finding elements by class name within the document.
- [x] Setting the text content of elements with the specified tag name.
- [x] Rendering the GopherParse object into an HTML string.
- [ ] Support for Selectors: Implement a selector engine for complex queries.
- [x] Element Manipulation: Add, remove, and modify attributes of elements.
- [ ] Traversing: Methods to traverse the document tree.
- [ ] Filtering: Filter selected elements based on specific conditions.
- [ ] Element Creation: Programmatically create and insert new elements.
- [ ] Text Searching and Highlighting: Search for specific text and highlight matches.
- [ ] Error Handling: Improve error messages and context for better debugging.
- [ ] Advanced Rendering Options: Control over rendering, indentation, and minification.
- [ ] Event Handling: Attach event listeners to elements and respond to interactions.
- [ ] Performance Optimization: Optimize parsing and manipulation for efficiency.
- [ ] Data Extraction: Extract data from structured documents automatically.
- [ ] XPath Support: Add support for XPath expressions as an alternative querying method.

Please note that the roadmap is subject to change based on feedback, contributions, and evolving requirements. Contributions and ideas from golang community are highly encouraged to shape the future development of GopherParse.





