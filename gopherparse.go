package gopherparse

import (
	"bytes"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

// GopherParse represents a GopherParse object, similar to Cheerio.
type GopherParse struct {
	root *html.Node
}

// LoadHTML parses an HTML document and returns a GopherParse object.
func LoadHTML(htmlContent string) (*GopherParse, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}
	return &GopherParse{root: doc}, nil
}

// LoadXML parses an XML document and returns a GopherParse object.
func LoadXML(xmlContent string) (*GopherParse, error) {
	doc, err := html.Parse(strings.NewReader(xmlContent))
	if err != nil {
		return nil, err
	}
	return &GopherParse{root: doc}, nil
}



// FindByTag finds all elements with the specified tag name within the GopherParse object.
func (gp *GopherParse) FindByTag(tagName string) []*html.Node {
	var elements []*html.Node
	var find func(*html.Node)
	find = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			elements = append(elements, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			find(c)
		}
	}
	find(gp.root)
	return elements
}

// FindByClass finds all elements with the specified class name within the GopherParse object.
func (gp *GopherParse) FindByClass(className string) []*html.Node {
	var elements []*html.Node
	var find func(*html.Node)
	find = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "class" && containsClass(attr.Val, className) {
					elements = append(elements, n)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			find(c)
		}
	}
	find(gp.root)
	return elements
}

// SetText sets the text content of all elements with the specified tag name within the GopherParse object.
func (gp *GopherParse) SetText(tagName, text string) {
	elements := gp.FindByTag(tagName)
	for _, elem := range elements {
		elem.Data = text
	}
}

// Render renders the GopherParse object into an HTML string.
func (gp *GopherParse) Render() string {
	var buf bytes.Buffer
	html.Render(&buf, gp.root)
	return buf.String()
}

// containsClass checks if the class list contains the specified class name.
func containsClass(classList, className string) bool {
	classes := strings.Fields(classList)
	for _, c := range classes {
		if c == className {
			return true
		}
	}
	return false
}
