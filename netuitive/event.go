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

package netuitive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var events []event

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type event struct {
	Data struct {
		Elementid string `json:"elementId"`
		Level     string `json:"level"`
		Message   string `json:"message"`
	} `json:"data"`
	Source string `json:"source"`
	Tags   []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"tags,omitempty"`
	Timestamp int64  `json:"timestamp"`
	Title     string `json:"title"`
	Type      string `json:"type"`
}

func PostEvent(apikey string, apiurl string, source string, element string, eventType string, title string, message string, level string, tags string, debug bool) error {

	url := apiurl + "/" + apikey

	var e event
	var es []event

	e.Source = source
	e.Title = title
	e.Type = eventType
	e.Data.Elementid = element
	e.Data.Level = level
	e.Data.Message = message
	e.Timestamp = int64(time.Now().UnixNano() / int64(time.Millisecond))

	if len(tags) > 0 {
		tag := strings.Split(tags, ",")

		for i := range tag {

			ts := strings.Split(tag[i], ":")

			e.Tags = append(e.Tags, struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			}{
				ts[0],
				ts[1],
			})

		}
	}

	es = append(es, e)

	jsonStr, _ := json.Marshal(&es)

	if debug {
		b, _ := json.MarshalIndent(es, "", "    ")
		fmt.Println("=== JSON PAYLOAD ===")
		fmt.Println(string(b) + "\n")
		fmt.Println("=== JSON PAYLOAD ===\n\n")
		fmt.Println("url = " + url)

	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return err
	}

	defer resp.Body.Close()

	if debug {
		fmt.Println("response Status:\n", resp.Status, "\n")
		fmt.Println("response Headers:\n", resp.Header, "\n")
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:\n", string(body), "\n")
	}

	if strings.Contains(resp.Status, "OK") {

		if debug {
			fmt.Println("response Status:\n", resp.Status, "\n")
			fmt.Println("response Headers:\n", resp.Header, "\n")
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("response Body:\n", string(body), "\n")
		}

	} else {
		fmt.Printf("\n\n\nError posting to %s\n", url)
		fmt.Errorf("Response: %s\n", resp.Status)
	}
	return nil
}
