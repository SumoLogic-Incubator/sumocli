package delete_subdomain

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func NewCmdAccountDeleteSubdomain(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-subdomain",
		Short: "Delete the configured subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			deleteSubdomain(client)
		},
	}
	return cmd
}

func deleteSubdomain(client *cip.APIClient) {
	httpResponse, errorResponse := client.DeleteSubdomain()
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "The subdomain was successfully deleted.")
	}
}
