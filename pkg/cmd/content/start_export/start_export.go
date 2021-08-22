package start_export

import (
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdStartExport(client *cip.APIClient) *cobra.Command {
	var (
		id          string
		isAdminMode bool
	)
	cmd := &cobra.Command{
		Use: "start-export",
		Short: "Starts an asynchronous export of content with the given identifier. You will be given a job identifier" +
			"which can be used with the sumocli content export-status command." +
			"If the content is a folder everything under that folder is exported recursively.",
		Run: func(cmd *cobra.Command, args []string) {
			startExport(id, isAdminMode, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the content item to export")
	cmd.Flags().BoolVar(&isAdminMode, "isAdminMode", false, "Set to true if you want to perform the request as a content administrator")
	cmd.MarkFlagRequired("id")
	return cmd
}

func startExport(id string, isAdminMode bool, client *cip.APIClient) {
	var options types.ContentOpts
	if isAdminMode == true {
		options.IsAdminMode = optional.NewString("true")
	} else {
		options.IsAdminMode = optional.NewString("false")
	}
	apiResponse, httpResponse, errorResponse := client.BeginAsyncExport(id, &options)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse, errorResponse)
	} else {
		cmdutils.Output(apiResponse, httpResponse, errorResponse, "")
	}
}
