package cmd

import (
	"github.com/Jhon-Henkel/go-lang-full-cycle-cobra-cli/internal/database"
	"github.com/spf13/cobra"
)

func NewCreateCommand(db database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Cria uma categoria",
		Long:  "",
		RunE:  RunCreate(db),
	}
}

func RunCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := NewCreateCommand(GetCategoryDb(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Nome da categoria")
	createCmd.Flags().StringP("description", "d", "", "Descrição da categoria")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
