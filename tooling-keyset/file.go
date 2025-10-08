package tooling_keyset

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PubKey struct {
	Type  string `json:"@type,omitempty"`
	Value string `json:"key,omitempty"`
}

type Validator struct {
	Address string `json:"address"`
	PubKey  PubKey `json:"pub_key"`
}

type PubKeysFile struct {
	Validators []Validator `json:"validators"`
}

type PrivValidatorKey struct {
	PubKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"pub_key"`
	PrivKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"priv_key"`
}

func getValidatorsInfo(filePath, homeTooling string) (int, []string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read file: %w", err)
	}

	var file PubKeysFile
	if err := json.Unmarshal(data, &file); err != nil {
		return 0, nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	validators := file.Validators
	count := len(validators)
	paths := make([]string, 0, count)

	for i, v := range validators {
		pathEdit := filepath.Join(homeTooling, "priv_"+fmt.Sprintf("validator%d", i+1)+"_key.json")
		paths = append(paths, pathEdit)
		if err := editPubKey(pathEdit, v.PubKey.Value); err != nil {
			return count, paths, err
		}
	}

	return count, paths, nil
}

func editPubKey(filePath, newPubKey string) error {
	// Tạo thư mục nếu chưa có
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	var pv PrivValidatorKey

	if _, err := os.Stat(filePath); err == nil {
		// File tồn tại → đọc và cập nhật
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", filePath, err)
		}
		if err := json.Unmarshal(data, &pv); err != nil {
			return fmt.Errorf("invalid JSON format in %s: %w", filePath, err)
		}

		pv.PubKey.Value = newPubKey
	} else {
		// File chưa có → tạo mới
		pv.PubKey.Type = "tendermint/PubKeyEd25519"
		pv.PubKey.Value = newPubKey
		pv.PrivKey.Type = "tendermint/PrivKeyEd25519"
		pv.PrivKey.Value = newPubKey
	}

	// Ghi lại file
	data, err := json.MarshalIndent(pv, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
