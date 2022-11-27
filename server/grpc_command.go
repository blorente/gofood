// A hacky re-creation of the serve command.
// We want to implement a command that is "serve+grpc", so we need to enhance serve.
package main

import (
	"github.com/fatih/color"
	"github.com/pocketbase/pocketbase"
	"github.com/spf13/cobra"
)
func NewGrpcCommand(app *pocketbase.PocketBase) *cobra.Command {
	var port string

	command := &cobra.Command{
		Use:   "serve-grpc",
		Short: "Starts the web server (default to localhost:8080)",
		Run: func(command *cobra.Command, args []string) {
			// reload app settings in case a new default value was set with a migration
			// (or if this is the first time the init migration was executed)
			if err := app.RefreshSettings(); err != nil {
				color.Yellow("=====================================")
				color.Yellow("WARNING - Settings load error! \n%v", err)
				color.Yellow("Fallback to the application defaults.")
				color.Yellow("=====================================")
			}
			CreateAndStartServer(*app, port)
		},
	}

	// command.PersistentFlags().StringSliceVar(
	// 	&allowedOrigins,
	// 	"origins",
	// 	[]string{"*"},
	// 	"CORS allowed domain origins list",
	// )

	command.PersistentFlags().StringVar(
		&port,
		"port",
		"8080",
		"grpc server port",
	)

	return command
}
