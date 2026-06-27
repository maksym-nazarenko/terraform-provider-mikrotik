package inspect

import (
	"encoding/json"
	"fmt"
	"io"
)

func WriteJSON(w io.Writer, node *Node) error {
	b, err := json.MarshalIndent(node, "", "    ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%s", b)

	return err
}
