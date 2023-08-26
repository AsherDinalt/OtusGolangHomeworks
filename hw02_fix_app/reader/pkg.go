package reader

import (
	"encoding/json"
	"fmt"
	"hw02_fix_app/types" //nolint:gci
	"io"
	"os" //nolint:gci
)

func ReadJSON(filePath string, _ int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f) //  было byte
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	json.Unmarshal(bytes, &data)

	res := data

	return res, nil
}
