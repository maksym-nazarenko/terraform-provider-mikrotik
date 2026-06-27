package inspect

const (
	TypeSelf       Type = "self"
	TypeChild      Type = "child"
	TypeCompletion Type = "completion"
)

const (
	NodeTypeCmd  NodeType = "cmd"
	NodeTypeDir  NodeType = "dir"
	NodeTypePath NodeType = "path"
	NodeTypeArg  NodeType = "arg"
	NodeTypeSelf NodeType = "self"
)

type (
	Type     string
	NodeType string

	Argument struct {
		Name string
		// Options holds a list of valid options for this argument, if any.
		Options []string
	}

	// ConsoleItem is a raw representation of `/console/inspect` items
	ConsoleItem struct {
		Name       string   `mikrotik:"name"`
		NodeType   NodeType `mikrotik:"node-type"`
		Type       Type     `mikrotik:"type"`
		Completion string   `mikrotik:"completion"`
		Show       bool     `mikrotik:"show"`
	}

	Node struct {
		Self      string
		Name      string
		Type      NodeType
		Children  []*Node
		Arguments []*Argument
	}

	Config struct {
		Root  string
		Depth int
	}
)
