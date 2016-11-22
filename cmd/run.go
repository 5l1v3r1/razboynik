// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"errors"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [url]",
	Short: "Run reverse shell with configuration",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			request *core.REQUEST
			err     error
		)

		if len(args) < 1 {
			return errors.New("Not enough arguments.")
		}

		shl := worker.BuildShellConfig(shellmethod, "")
		php := worker.BuildPHPConfig(raw, false)
		srv := worker.BuildServerConfig(args[0], method, parameter, key, raw)

		request = worker.BuildRequest(shl, php, srv)

		err = worker.Run(request)

		return err
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
