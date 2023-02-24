package cmd

import (
	"log"

	"github.com/ebizno/Ula/internal/email"
	"github.com/ebizno/Ula/internal/events"
	"github.com/ebizno/Ula/internal/file"
	"github.com/spf13/cobra"
)

var (
	from        string
	password    string
	portEmail   int
	host        string
	subject     string
	body        string
	contentType string
	fileName    string
)

func init() {
	CmdEvent.Flags().StringVarP(&from, "from", "f", "", "Email from")
	CmdEvent.Flags().StringVarP(&password, "password", "p", "", "Email password")
	CmdEvent.Flags().IntVarP(&portEmail, "port", "P", 0, "Email port")
	CmdEvent.Flags().StringVarP(&host, "host", "H", "", "Email host")
	CmdEvent.Flags().StringVarP(&subject, "subject", "s", "", "Email subject")
	CmdEvent.Flags().StringVarP(&body, "body", "b", "", "Email body")
	CmdEvent.Flags().StringVarP(&contentType, "content-type", "c", "", "Email content type")
	CmdEvent.Flags().StringVarP(&fileName, "file-name", "n", "", "File name")
}

var CmdEvent = &cobra.Command{
	Use:   "event",
	Short: "Ula watcher of system event",
	Long: `Ula watcher of system event
			send email when is created or modificed`,
	RunE: ExecuteEvent,
}

func ExecuteEvent(cmd *cobra.Command, args []string) error {
	emailCredentials, err := email.NewEmailCredential(from, password, portEmail, host)
	if err != nil {
		log.Println(err)
	}

	file, err := file.NewFilePath(fileName)
	if err != nil {
		log.Println(err)
	}

	events.NewEvent(emailCredentials, subject, body, contentType, file)

	return nil
}
