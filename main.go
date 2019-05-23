/**
*  Copyright 2019 Francisco Javier Collado Valle
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
**/

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	maven "github.com/javicv/mvn_resolver/xml"
)

func main() {

	var REPO, isSet = os.LookupEnv("MAVEN_REPO_URL")
	if !isSet {
		fmt.Fprintf(os.Stderr, "Error: Environment variable MAVEN_REPO_URL must be set\n")
		os.Exit(1)
	}

	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Error: Not enough parameters provided\n")
		fmt.Fprintf(os.Stderr, "Uso: %s GROUP_ID ARTIFACT_ID VERSION PACKAGING\n", os.Args[0])
		os.Exit(2)
	}
	var groupID = strings.Split(os.Args[1], ".")
	var artifactID = os.Args[2]
	var version = os.Args[3]
	var packaging = os.Args[4]

	var metadataURL = fmt.Sprintf("%s/%s/%s/%s/maven-metadata.xml", REPO, strings.Join(groupID, "/"), artifactID, version)
	resp, err := http.Get(metadataURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(3)
	}
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "%s: %s\n", resp.Status, metadataURL)
		os.Exit(3)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing response\n")
		os.Exit(4)
	}
	var metadata maven.Metadata
	err = xml.Unmarshal(body, &metadata)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(5)
	}

	for _, snapshotVersion := range metadata.Versioning.SnapshotVersions.SnapshotVersion {
		if snapshotVersion.Extension == packaging {
			var artifactURL = fmt.Sprintf("%s/%s/%s/%s/%s-%s.%s", REPO, strings.Join(groupID, "/"), artifactID, version, artifactID, snapshotVersion.Value, snapshotVersion.Extension)
			fmt.Print(artifactURL)
		}
	}
}
