package create_subdomain

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdAccountCreateSubdomain(client *cip.APIClient, log *zerolog.Logger) *cobra.Command {
	var subdomain string
	cmd := &cobra.Command{
		Use:   "create-subdomain",
		Short: "Create a subdomain. Only the Account Owner can create a subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			createSubdomain(subdomain, client, log)
		},
	}
	cmd.Flags().StringVar(&subdomain, "subdomain", "", "Specify a subdomain (minimum 4 and maximum 63 characters)")
	cmd.MarkFlagRequired("subdomain")
	return cmd
}

func createSubdomain(subdomain string, client *cip.APIClient, log *zerolog.Logger) {
	apiResponse, httpResponse, errorResponse := client.CreateSubdomain(types.ConfigureSubdomainRequest{
		Subdomain: subdomain,
	})
	if errorResponse != nil {
		log.Error().Err(errorResponse).Msg("failed to create subdomain")
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
