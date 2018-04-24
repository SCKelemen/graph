package lex

import (
	"path/filepath"
)

type ITokenizer interface {
	CanTokenize(path string) bool
}

type TextTokenizer struct {
	Extensions []string
}

func (t TextTokenizer) CanTokenize(path string) bool {
	extension := filepath.Ext(path)
	switch extension {
	case "txt":
	case "text":
		break
	}
}

type RoslynTokenizer struct {
}

func (t RoslynTokenizer) CanTokenize(path string) bool {

}

type Extension string

const (
	TextExtension = "text"
	TxtExtension  = "txt"
	PDFExtension  = "pdf"
	DOCExtension  = "doc"
	DOCXExtension = "docx"
)

type ExtensionToken string

const (
	TextToken    ExtensionToken = "text"
	WordToken    ExtensionToken = "word"
	PDFToken     ExtensionToken = "pdf"
	RoslynToken  ExtensionToken = "roslyn"
	UnknownToken ExtensionToken = "unknown"
)

type ExtensionTokenizer struct {
}

func (et ExtensionTokenizer) Tokenize(path string) ExtensionToken {
	ext := filepath.Ext(path)
	switch ext {
	case "text":
	case "txt":
		return TextToken
	case "doc":
	case "docx":
		return WordToken
	case "pdf":
		return PDFToken
	case "vb":
	case "cs":
	case "csproj":
	case "sln":
	case "vbproj":
	case "asp":
	case "asmx": // https://github.com/dotnet/roslyn/blob/df30a634dabfdcf888f360acaa285b968759e4d2/src/Workspaces/Core/Portable/Workspace/Solution/Document.cs#L147
		return RoslynToken
	default:
		return UnknownToken
	}

	return UnknownToken
}

/*
The ability to tokenize a codefile depends on the project or solution being
compilable. This may require somework. If given a code file, we might be able
to walk up the directory to try and find the corresponding solution or project
but that could be complex, and indeterminate. Perhaps it would be better to only
look for project or solution files.
*/
