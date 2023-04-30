/*
Copyright © 2023 Pilot Framework

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a Pilot Framework environment.",
	Long: `This command provisions the necessary resources for Pilot to function.

Pilot utilizes a customized version of HashiCorp's Waypoint as the base of the platform,
which can be deployed to either AWS EKS or GCP GKE. For more information about the Pilot Server,
visit https://pilotframework.com/architecture.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setup called")
	},
}

func init() {
	setupCmd.Flags().BoolP("aws", "", false, "Provision Pilot Server on AWS EKS.")
	setupCmd.Flags().BoolP("gcp", "", false, "Provision Pilot Server on GCP GKE.")
	setupCmd.Flags().BoolP("dev", "d", false, "Provision Pilot Server using our experimental image. Not guaranteed to work.")
	setupCmd.Flags().BoolP("bare", "", false, "Provision Pilot Server using HashiCorp's default Waypoint image. \nWARNING: you will not have access to Pilot's custom plugins.")

	setupCmd.Flags().MarkHidden("dev")
	setupCmd.Flags().SortFlags = false

	ServerCmd.AddCommand(setupCmd)
}
