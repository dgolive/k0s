/*
Copyright 2021 Mirantis, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func cmdFlagsToArgs(cmd *cobra.Command) []string {
	flagsAndVals := []string{}
	// Use visitor to collect all flags and vals into slice
	cmd.Flags().Visit(func(f *pflag.Flag) {
		switch f.Value.Type() {
		case "stringSlice", "stringToString":
			val := f.Value.String()
			flagsAndVals = append(flagsAndVals, fmt.Sprintf(`--%s="%s"`, f.Name, strings.Trim(val, "[]")))
		default:
			flagsAndVals = append(flagsAndVals, fmt.Sprintf("--%s=%s", f.Name, f.Value.String()))
		}
	})
	return flagsAndVals
}
