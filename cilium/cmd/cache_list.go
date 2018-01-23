// Copyright 2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	identityApi "github.com/cilium/cilium/api/v1/client/policy"
	pkg "github.com/cilium/cilium/pkg/client"
	"github.com/cilium/cilium/pkg/policy"
	"github.com/spf13/cobra"
)

// identityListCmd represents the identity_list command
var cacheListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contents of endpoint IP cache",
	Run: func(cmd *cobra.Command, args []string) {
		listEndpointIPIdentityMapping(args)
	},
}

func init() {
	rootCmd.AddCommand(identityListCmd)
	AddMultipleOutput(identityListCmd)
}

func listEndpointIPIdentityMapping(args []string) {
	/*result := make(map[string]interface{})

	var params *identityApi.GetIdentityParams
	if len(args) != 0 {
		params = identityApi.NewGetIdentityParams().WithLabels(args)
	}

	identities, err := client.Policy.GetIdentity(params)
	if err != nil {
		if params != nil {
			Fatalf("Cannot get identities for given labels %v. err: %s\n", params.Labels, err.Error())
		} else {
			Fatalf("Cannot get identities. err: %s", pkg.Hint(err))
		}
	}
	result["identities"] = identities.Payload

	if len(dumpOutput) > 0 {
		if err := OutputPrinter(result); err != nil {
			os.Exit(1)
		}
		return
	}

	if reservedIDs {
		fmt.Println("Reserved identities:")
		for k, v := range result["reservedIDs"].(map[string]string) {
			fmt.Printf("\t %3s %-15s \n", v, k)
		}
	}

	fmt.Println("Identities in use by endpoints:\n" +
		"(Note: If labels have been provided as parameters, only matching identities will be displayed)")

	payload, err := json.MarshalIndent(result["identities"], "", "  ")
	if err != nil {
		Fatalf("Cannot marshal identities %s", err.Error())
	}
	fmt.Println(string(payload))*/
}
