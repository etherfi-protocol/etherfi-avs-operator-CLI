package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// ExportJSON provides a standardized output for cli commands that produce an artifact
func ExportJSON(operation string, operatorID int64, out any) error {

	buf, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		return fmt.Errorf("marshalling output: %w", err)
	}

	fname := fmt.Sprintf("%s-%d.json", operation, operatorID)
	fmt.Printf("Output written to %s\n%s\n", fname, string(buf))
	return os.WriteFile(fname, buf, 0666)
}
