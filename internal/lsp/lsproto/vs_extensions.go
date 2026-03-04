package lsproto

// VS-specific LSP extensions for richer editor integration.
// These types implement the _vs_ prefixed extensions that Visual Studio's
// LSP client understands for enhanced hover and signature help rendering.
// Property names must be PascalCase and use _vs_type as the discriminator.

// ClassifiedTextRun represents a single classified run of text for VS rendering.
type ClassifiedTextRun struct {
	VSType             string  `json:"_vs_type"`
	ClassificationName string  `json:"ClassificationTypeName"`
	Text               string  `json:"Text"`
	MarkerTagType      *string `json:"MarkerTagType"`
	Style              int     `json:"Style"`
}

// NewClassifiedTextRun creates a ClassifiedTextRun.
func NewClassifiedTextRun(classification string, text string) *ClassifiedTextRun {
	return &ClassifiedTextRun{
		VSType:             "ClassifiedTextRun",
		ClassificationName: classification,
		Text:               text,
		Style:              0,
	}
}

// ClassifiedTextElement represents a sequence of classified text runs.
type ClassifiedTextElement struct {
	VSType string              `json:"_vs_type"`
	Runs   []*ClassifiedTextRun `json:"Runs"`
}

// NewClassifiedTextElement creates a ClassifiedTextElement with the given runs.
func NewClassifiedTextElement(runs []*ClassifiedTextRun) *ClassifiedTextElement {
	return &ClassifiedTextElement{
		VSType: "ClassifiedTextElement",
		Runs:   runs,
	}
}

// ImageID identifies a VS image.
type ImageID struct {
	VSType string `json:"_vs_type"`
	GUID   string `json:"Guid"`
	ID     int    `json:"Id"`
}

// ImageElement represents an image for VS rendering.
type ImageElement struct {
	VSType  string  `json:"_vs_type"`
	ImageID ImageID `json:"ImageId"`
}

// NewImageElement creates an ImageElement with the given GUID and ID.
func NewImageElement(guid string, id int) *ImageElement {
	return &ImageElement{
		VSType: "ImageElement",
		ImageID: ImageID{
			VSType: "ImageId",
			GUID:   guid,
			ID:     id,
		},
	}
}

// ContainerElement groups child elements for VS rendering.
type ContainerElement struct {
	VSType   string `json:"_vs_type"`
	Elements []any  `json:"Elements"`
	Style    int    `json:"Style"`
}

// NewContainerElement creates a ContainerElement with the given elements.
func NewContainerElement(style int, elements ...any) *ContainerElement {
	return &ContainerElement{
		VSType:   "ContainerElement",
		Elements: elements,
		Style:    style,
	}
}

// Classification type names matching VS's ClassificationTypeNames.
const (
	ClassificationKeyword       = "keyword"
	ClassificationIdentifier    = "identifier"
	ClassificationOperator      = "operator"
	ClassificationPunctuation   = "punctuation"
	ClassificationText          = "text"
	ClassificationStringLiteral = "string"
	ClassificationNumberLiteral = "number"
	ClassificationWhiteSpace    = "whitespace"
)

// Well-known VS image GUID.
const KnownImageGUID = "{ae27a6b0-e345-4288-96df-5eaf394ee369}"

// Image IDs from VS KnownImageIds catalog.
const (
	ImageIDClass      = 463
	ImageIDInterface  = 1356
	ImageIDEnum       = 745
	ImageIDEnumMember = 753
	ImageIDMethod     = 1874
	ImageIDProperty   = 2449
	ImageIDField      = 817
	ImageIDVariable   = 3345
	ImageIDConstant   = 585
	ImageIDModule     = 1938
	ImageIDType       = 3183
	ImageIDParameter  = 2186
	ImageIDKeyword    = 1452
)
