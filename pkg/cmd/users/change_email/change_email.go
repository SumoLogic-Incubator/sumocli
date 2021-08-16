package change_email

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/pkg/cmdutils"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip/types"
)

func NewCmdUserChangeEmail(client *cip.APIClient) *cobra.Command {
	var (
		id    string
		email string
	)
	cmd := &cobra.Command{
		Use:   "change-email",
		Short: "Changes the email address of a Sumo Logic user.",
		Run: func(cmd *cobra.Command, args []string) {
			userChangeEmail(id, email, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the user that needs to have the email changed.")
	cmd.Flags().StringVar(&email, "email", "", "Specify the users new email address.")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("email")
	return cmd
}

func userChangeEmail(id string, email string, client *cip.APIClient) {
	httpResponse, errorResponse := client.RequestChangeEmail(types.ChangeEmailRequest{
		Email: email,
	},
		id)
	if errorResponse != nil {
		cmdutils.OutputError(httpResponse)
	} else {
		cmdutils.Output(nil, httpResponse, errorResponse, "Email change request was submitted successfully.")
	}
}