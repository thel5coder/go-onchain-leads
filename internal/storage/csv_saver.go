package storage

import (
	"encoding/csv"
	"fmt"
	"go-onchain-leads/internal/domain"
	"os"
)

type CSVSaver struct {
	filePath string
}

func NewCSVSaver(filePath string) *CSVSaver {
	return &CSVSaver{
		filePath: filePath,
	}
}

func (s *CSVSaver) SaveLead(lead domain.Lead) error {
	// 1. Check if the file already exists
	fileExists := true
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		fileExists = false
	}

	// 2. Open the file in "Append" mode (create it if it doesn't exist)
	file, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 3. If it's a brand new file, write the Headers first!
	if !fileExists {
		headers := []string{"ENS Name", "Twitter", "Email", "GitHub", "Wallet Address", "Contract Address", "Gas Used"}
		if err := writer.Write(headers); err != nil {
			return err
		}
	}

	// 4. Write our golden lead data as a new row
	row := []string{
		lead.ENSName,
		lead.Twitter,
		lead.Email,
		lead.GitHub,
		lead.DeveloperWallet,
		lead.ContractAddress,
		fmt.Sprintf("%d", lead.GasUsed),
	}

	return writer.Write(row)
}
