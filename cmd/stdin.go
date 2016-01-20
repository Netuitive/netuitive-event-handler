// Copyright © 2016 Netuitive, Inc.
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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/netuitive/netuitive-event-handler/netuitive"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type sensuPipe struct {
	Client struct {
		Name          string   `json:"name"`
		Address       string   `json:"address"`
		Subscriptions []string `json:"subscriptions"`
		Timestamp     int64    `json:"timestamp"`
	} `json:"client"`
	Check struct {
		Name        string   `json:"name"`
		Issued      int      `json:"issued"`
		Output      string   `json:"output"`
		Status      int      `json:"status"`
		Command     string   `json:"command"`
		Subscribers []string `json:"subscribers"`
		Interval    int      `json:"interval"`
		Handler     string   `json:"handler"`
		History     []string `json:"history"`
		Flapping    bool     `json:"flapping"`
	} `json:"check"`
	Occurrences int    `json:"occurrences"`
	Action      string `json:"action"`
	ID          string `json:"id"`
}

var pipeJSON *sensuPipe

// stdinCmd represents the stdin command
var stdinCmd = &cobra.Command{
	Use:   "stdin",
	Short: "Post events to Netuitive from the stdin pipe",
	Long:  Name + " (" + Version + ")\n\n" + `Pipe in a properly formatted (sensu format) JSON payload`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Work your own magic here

		if viper.GetString("apikey") == "" {
			return errors.New("missing apikey\n")
		}

		if viper.GetString("url") == "" {
			return errors.New("missing url\n")
		}

		eventOccurrences := strconv.Itoa(pipeJSON.Occurrences)

		eventID := pipeJSON.ID
		eventAction := pipeJSON.Action

		element = pipeJSON.Client.Name

		checkName := pipeJSON.Check.Name
		checkCommand := pipeJSON.Check.Command
		checkOutput := pipeJSON.Check.Output
		status := pipeJSON.Check.Status

		switch status {
		case 0:
			level = "OK"
		case 1:
			level = "WARNING"
		case 2:
			level = "CRITICAL"
		default:
			level = "UNKNOWN"
		}

		title = element + " " + checkName + " is " + level

		message = "checkname: " + checkName
		message = message + " command: " + checkCommand
		message = message + " output: " + checkOutput
		message = message + " status: " + level
		message = message + " occurrences: " + eventOccurrences
		message = message + " id: " + eventID
		message = message + " action: " + eventAction

		tags = "checkname:" + checkName
		tags = tags + ",command:" + checkCommand
		tags = tags + ",output:" + checkOutput
		tags = tags + ",status: " + level
		tags = tags + ",occurrences:" + eventOccurrences
		tags = tags + ",id:" + eventID
		tags = tags + ",action:" + eventAction

		err := netuitive.PostEvent(viper.GetString("apikey"), viper.GetString("url"), source, element, eventType, title, message, level, tags, debug)

		return err

	},
}

func init() {
	RootCmd.AddCommand(stdinCmd)

	if !terminal.IsTerminal(0) {
		stdinBytes, err := ioutil.ReadAll(os.Stdin)
		if err == nil {
			json.Unmarshal(stdinBytes, &pipeJSON)
			if debug {
				fmt.Println(string(stdinBytes))
			}
		}
	}

}
