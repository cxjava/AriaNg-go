package cmd

import (
	"log"
	"net/http"

	"github.com/cxjava/AriaNg-go/assets"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server for AriaNg",
	Long:  `Start a server for AriaNg`,
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

var addr string

func server() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(assets.HTTP)))
	log.Println("starting server on " + addr)
	log.Println(http.ListenAndServe(addr, nil))
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&addr, "address", "a", ":18080", "bind address.")
}
