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
	var USERNAME, isSetUsername = os.LookupEnv("MAVEN_USERNAME")
	var PASSWORD, isSetPassword = os.LookupEnv("MAVEN_PASSWORD")

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
	client := &http.Client{}
	req, err := http.NewRequest("GET", metadataURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(3)
	}
	if isSetUsername && isSetPassword {
		req.SetBasicAuth(USERNAME, PASSWORD)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(4)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		processSnapshot(*resp, REPO, groupID, artifactID, version, packaging)
		break
	case 400:
		processRelease(REPO, groupID, artifactID, version, packaging)
		break
	case 404:
		processRelease(REPO, groupID, artifactID, version, packaging)
		break
	default:
		fmt.Fprintf(os.Stderr, "%s: %s\n", resp.Status, metadataURL)
		os.Exit(4)
	}
}

func processRelease(repo string, groupID []string, artifactID string, version string, packaging string) {
	var artifactURL = fmt.Sprintf("%s/%s/%s/%s/%s-%s.%s", repo, strings.Join(groupID, "/"), artifactID, version, artifactID, version, packaging)

	client := &http.Client{}
	req, err := http.NewRequest("HEAD", artifactURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(5)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(6)
	}
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Error %s: %s\n", resp.Status, artifactURL)
		os.Exit(7)
	}
	fmt.Print(artifactURL)
}

func processSnapshot(resp http.Response, repo string, groupID []string, artifactID string, version string, packaging string) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing response\n")
		os.Exit(8)
	}
	var metadata maven.Metadata
	err = xml.Unmarshal(body, &metadata)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(9)
	}

	snapshotVersion := maven.Filter(metadata.Versioning.SnapshotVersions.SnapshotVersion, func(v maven.SnapshotVersion) bool { return v.Extension == packaging })
	for _, sv := range snapshotVersion {
		var artifactURL = fmt.Sprintf("%s/%s/%s/%s/%s-%s.%s", repo, strings.Join(groupID, "/"), artifactID, version, artifactID, sv.Value, sv.Extension)
		fmt.Print(artifactURL)
	}
}
