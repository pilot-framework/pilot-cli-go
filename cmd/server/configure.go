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

var (
	gcpProject string
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure remote Pilot Server.",
	Long: `Configure remote Pilot Server with credentials for the selected cloud provider(s).

This typically only needs to be run once for each provider, as it is stored within the cluster
as part of its ConfigMap or as a Secret.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
	},
}

func init() {
	configureCmd.Flags().BoolP("aws", "", false, "Configure server with AWS Pilot IAM credentials.")
	configureCmd.Flags().BoolP("gcp", "", false, "Configure server with GCP Pilot IAM credentials.")
	configureCmd.Flags().StringVarP(&gcpProject, "project", "p", "", "Project ID for GCP Project that the service account and IAM role will be created for.")

	ServerCmd.AddCommand(configureCmd)
}
