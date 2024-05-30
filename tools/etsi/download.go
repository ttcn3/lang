package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ttcn3/lang/tools/etsi/internal/documents"
)

type Item struct {
	DocID      string   `json:"doc_id"`
	Year       string   `json:"year,omitempty"`
	WorkItemID int      `json:"wki_id"`
	Title      string   `json:"title"`
	Files      []string `json:"files,omitempty"`
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download list of ETSI TTCN-3 deliverables",
	RunE: func(cmd *cobra.Command, args []string) error {
		deliverables, err := documents.Deliverables()
		if err != nil {
			return err
		}
		b, err := json.MarshalIndent(deliverables, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))

		return nil

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
