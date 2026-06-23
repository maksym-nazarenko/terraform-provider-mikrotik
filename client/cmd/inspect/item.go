package main

const (
	NodeTypeCmd  NodeType = "cmd"
	NodeTypeDir  NodeType = "dir"
	NodeTypeArg  NodeType = "arg"
	NodeTypeSelf NodeType = "self"

	TypeSelf       Type = "self"
	TypeChild      Type = "child"
	TypeCompletion Type = "completion"
)

type (
	NodeType string
	Type     string

	Argument struct {
		Name    string
		Options []string
	}

	ConsoleItem struct {
		Name       string   `mikrotik:"name"`
		NodeType   NodeType `mikrotik:"node-type"`
		Type       Type     `mikrotik:"type"`
		Completion string   `mikrotik:"completion"`
		Show       bool     `mikrotik:"show"`
	}

	Item struct {
		Self      string
		Name      string
		NodeType  NodeType
		Children  []*Item
		Arguments []*Argument
	}

	InspectConfig struct {
		Root  string
		Depth int
	}
)
