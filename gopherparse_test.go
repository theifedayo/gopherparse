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
