package gopherparse

import (
	"testing"
)

func TestLoadHTML(t *testing.T) {
	htmlContent := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Sample HTML Document</title>
		</head>
		<body>
			<h1>Hello, World!</h1>
			<p class="intro">This is a sample HTML document.</p>
			<p class="description">This is a description.</p>
			<ul>
				<li>Item 1</li>
				<li class="highlight">Item 2</li>
				<li>Item 3</li>
			</ul>
		</body>
		</html>
	`

	// Test LoadHTML
	gpObj, err := LoadHTML(htmlContent)
	if err != nil {
		t.Fatalf("LoadHTML failed with error: %v", err)
	}

	// Test FindByTag
	h1Elements := gpObj.FindByTag("h1")
	if len(h1Elements) != 1 {
		t.Errorf("FindByTag - Expected 1 <h1> element, got %d", len(h1Elements))
	}

	// Test FindByClass
	introElements := gpObj.FindByClass("intro")
	if len(introElements) != 1 {
		t.Errorf("FindByClass - Expected 1 element with class 'intro', got %d", len(introElements))
	}

	// Test SetText
	// gpObj.SetText("h1", "Updated Heading")
	// h1Updated := gpObj.FindByTag("h1")[0]
	// if h1Updated.FirstChild.Data != "Updated Heading" {
	// 	t.Errorf("SetText - Expected heading text 'Updated Heading', got '%s'", h1Updated.FirstChild.Data)
	// }
}

func TestLoadXML(t *testing.T) {
	xmlContent := `
		<root>
			<item>Item 1</item>
			<item>Item 2</item>
			<item class="highlight">Item 3</item>
		</root>
	`

	// Test LoadXML
	gpObj, err := LoadXML(xmlContent)
	if err != nil {
		t.Fatalf("LoadXML failed with error: %v", err)
	}

	// Test FindByTag
	itemElements := gpObj.FindByTag("item")
	if len(itemElements) != 3 {
		t.Errorf("FindByTag - Expected 3 <item> elements, got %d", len(itemElements))
	}

	// Test FindByClass
	highlightElements := gpObj.FindByClass("highlight")
	if len(highlightElements) != 1 {
		t.Errorf("FindByClass - Expected 1 element with class 'highlight', got %d", len(highlightElements))
	}

	// Test SetText
	// gpObj.SetText("item", "Updated Item")
	// itemUpdated := gpObj.FindByTag("item")[0]
	// if itemUpdated.FirstChild.Data != "Updated Item" {
	// 	t.Errorf("SetText - Expected item text 'Updated Item', got '%s'", itemUpdated.FirstChild.Data)
	// }
}

func TestLoadHTMLFile(t *testing.T) {
	// Test LoadHTMLFile with a sample HTML file
	htmlFilePath := "testdata/sample.html"
	gpObj, err := LoadHTMLFile(htmlFilePath)
	if err != nil {
		t.Fatalf("LoadHTMLFile failed with error: %v", err)
	}

	// Test FindByTag
	h1Elements := gpObj.FindByTag("h1")
	if len(h1Elements) != 1 {
		t.Errorf("FindByTag - Expected 1 <h1> element, got %d", len(h1Elements))
	}
}

func TestLoadXMLFile(t *testing.T) {
	// Test LoadXMLFile with a sample XML file
	xmlFilePath := "testdata/sample.xml"
	gpObj, err := LoadXMLFile(xmlFilePath)
	if err != nil {
		t.Fatalf("LoadXMLFile failed with error: %v", err)
	}

	// Test FindByTag
	itemElements := gpObj.FindByTag("item")
	if len(itemElements) != 3 {
		t.Errorf("FindByTag - Expected 3 <item> elements, got %d", len(itemElements))
	}
}
