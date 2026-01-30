package mcp

import (
	"encoding/base64"
	"encoding/json"

	"github.com/invopop/jsonschema"
)


type MCPProtocolVersion string


const (
	MCPProtocolVersion20250618 MCPProtocolVersion = "2025-06-18"
	MCPProtocolVersion20250326 MCPProtocolVersion = "2025-03-26"
	MCPProtocolVersion20241105 MCPProtocolVersion = "2024-11-05"
)


const LatestSupportedMCPVersion = MCPProtocolVersion20250618


const ToolResultStructuredContentMinVersion = MCPProtocolVersion20250618


const MCPProtocolVersionHeaderName = "MCP-Protocol-Version"


var SupportedMCPVersion = map[MCPProtocolVersion]bool{
	MCPProtocolVersion20250618: true,
	MCPProtocolVersion20250326: true,
	MCPProtocolVersion20241105: true,
}


type MCPInitializeRequest struct {
	ProtocolVersion string             `json:"protocolVersion"`
	ClientInfo      *MCPImplementation `json:"clientInfo"`
}


type MCPInitializeResponse struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    *MCPCapabilities   `json:"capabilities"`
	ServerInfo      *MCPImplementation `json:"serverInfo"`
}


type MCPCapabilities struct {
	Resources *MCPResourceCapabilities `json:"resources,omitempty"`
	Tools     *MCPToolCapabilities     `json:"tools,omitempty"`
	Prompts   *MCPPromptCapabilities   `json:"prompts,omitempty"`
}


type MCPImplementation struct {
	Name    string `json:"name"`
	Title   string `json:"title,omitempty"`
	Version string `json:"version"`
}


type MCPResourceCapabilities struct {
	Subscribe   bool `json:"subscribe"`
	ListChanged bool `json:"listChanged"`
}


type MCPToolCapabilities struct {
	ListChanged bool `json:"listChanged"`
}


type MCPPromptCapabilities struct {
	ListChanged bool `json:"listChanged"`
}


type MCPListResourcesResponse struct {
	Resources  []*MCPResource `json:"resources"`
	NextCursor string         `json:"nextCursor,omitempty"`
}


type MCPResource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Size        int    `json:"size,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}


type MCPReadResourceRequest struct {
	URI string `json:"uri"`
}


type MCPReadResourceResponse[T MCPTextResourceContents | MCPBlobResourceContents] struct {
	Contents []*T `json:"contents"`
}


type MCPTextResourceContents struct {
	URI      string `json:"uri"`
	Text     string `json:"text"`
	MimeType string `json:"mimeType,omitempty"`
}


type MCPBlobResourceContents struct {
	URI      string `json:"uri"`
	Blob     string `json:"blob"` 
	MimeType string `json:"mimeType,omitempty"`
}


type MCPListToolsResponse struct {
	Tools      []*MCPTool `json:"tools"`
	NextCursor string     `json:"nextCursor,omitempty"`
}


type MCPTool struct {
	Name         string             `json:"name"`
	InputSchema  *jsonschema.Schema `json:"inputSchema"`
	OutputSchema *jsonschema.Schema `json:"outputSchema,omitempty"`
	Title        string             `json:"title,omitempty"`
	Description  string             `json:"description,omitempty"`
}


type MCPCallToolRequest struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
}


type MCPCallToolResponse[T MCPTextContent | MCPImageContent | MCPAudioContent | MCPResourceLink | MCPEmbeddedResource] struct {
	Content           []*T `json:"content"`
	StructuredContent any  `json:"structuredContent,omitempty"`
	IsError           bool `json:"isError,omitempty"`
}


type MCPTextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}


type MCPImageContent struct {
	Type     string `json:"type"`
	MimeType string `json:"mimeType"`
	Data     string `json:"data"` 
}


type MCPAudioContent struct {
	Type     string `json:"type"`
	MimeType string `json:"mimeType"`
	Data     string `json:"data"` 
}


type MCPResourceLink struct {
	URI         string `json:"uri"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Size        int    `json:"size,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}


type MCPEmbeddedResource struct {
	Type     string `json:"type"`
	Resource any    `json:"resource"`
}


func NewMCPTextContent(text string) *MCPTextContent {
	return &MCPTextContent{
		Type: "text",
		Text: text,
	}
}


func NewMCPImageContent(data []byte, mimeType string) *MCPImageContent {
	return &MCPImageContent{
		Type:     "image",
		MimeType: mimeType,
		Data:     base64.StdEncoding.EncodeToString(data),
	}
}


func NewMCPAudioContent(data []byte, mimeType string) *MCPAudioContent {
	return &MCPAudioContent{
		Type:     "audio",
		MimeType: mimeType,
		Data:     base64.StdEncoding.EncodeToString(data),
	}
}


func NewMCPResourceLink(uri string, name string) *MCPResourceLink {
	return &MCPResourceLink{
		URI:  uri,
		Type: "resource_link",
		Name: name,
	}
}


func NewMCPEmbeddedResource[T MCPTextResourceContents | MCPBlobResourceContents](resource *T) *MCPEmbeddedResource {
	return &MCPEmbeddedResource{
		Type:     "resource",
		Resource: resource,
	}
}
