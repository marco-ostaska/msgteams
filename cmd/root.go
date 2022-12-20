package cmd

import (
	"fmt"
	"os"

	"github.com/marco-ostaska/msgteams/internal/teams"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "msgteams",
	Short: "msgteams allows you to send messages to a Microsoft Teams channel from the command line. ",
	Long: `msgteams allows you to send messages to a Microsoft Teams channel from the command line. It is a convenient way to send messages as part of automated processes or scripts.

To use msgteams, you will need to generate a URL for your Microsoft Teams channel. This can be done by following these steps:

- Open your Microsoft Teams channel.
- Click on the three dots in the top right corner of the channel.
- From the menu that appears, select "Connectors".
- Scroll down until you see the "Incoming Webhook" connector. If it is not appearing, you may need to contact your Teams administrator to have it added to your channel.
- Click on the "Configure" or "Add" button for the Incoming Webhook connector.
- Provide a name for the incoming webhook and click "Create".
- Click "Save" to save the URL provided.

The URL provided is a unique address that you will use to send messages to your Microsoft Teams channel. You will need to save the URL so that you can use it with the command.

Please note that Microsoft may change the process for creating incoming webhooks at any time. If the steps above do not work, you can find the most up-to-date instructions on the Microsoft Teams documentation page

	`,
	Version: "0.1.0",
	Example: `
	msgteams --url <your URL> --message "your message"

	or

	msgteams -u <your URL> -m "your message"
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: postMsg,
}

func postMsg(cmd *cobra.Command, args []string) error {
	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}

	msg, err := cmd.Flags().GetString("message")

	if err != nil {
		return err
	}

	if err := teams.NewPostMsg(url, msg); err == nil {
		fmt.Println("Message sent")
	}

	return err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd.DisableAutoGenTag = true
	// errDoc := doc.GenMarkdownTree(rootCmd, "./docs")
	// fmt.Println("Docs: ", errDoc.Error())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.msgteams.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringP("url", "u", "", "url to microsoft Teams channel")
	rootCmd.Flags().StringP("message", "m", "", "message to be sent to teams channel")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("message")

}
