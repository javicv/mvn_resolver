/*
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
 */

package xml

import "encoding/xml"

// Metadata node
type Metadata struct {
	XMLName      xml.Name   `xml:"metadata"`
	ModelVersion string     `xml:"modelVersion,attr"`
	GroupID      string     `xml:"groupId"`
	ArtifactID   string     `xml:"artifactId"`
	Version      string     `xml:"version"`
	Versioning   Versioning `xml:"versioning"`
}

// Versioning node
type Versioning struct {
	XMLName          xml.Name         `xml:"versioning"`
	Snapshot         Snapshot         `xml:"snapshot"`
	LastUpdated      string           `xml:"lastUpdated"`
	SnapshotVersions SnapshotVersions `xml:"snapshotVersions"`
}

// Snapshot node
type Snapshot struct {
	XMLName     xml.Name `xml:"snapshot"`
	Timestamp   string   `xml:"timestamp"`
	BuildNumber string   `xml:"buildNumber"`
}

// SnapshotVersions node
type SnapshotVersions struct {
	XMLName         xml.Name          `xml:"snapshotVersions"`
	SnapshotVersion []SnapshotVersion `xml:"snapshotVersion"`
}

// SnapshotVersion node
type SnapshotVersion struct {
	XMLName   xml.Name `xml:"snapshotVersion"`
	Extension string   `xml:"extension"`
	Value     string   `xml:"value"`
	Updated   string   `xml:"updated"`
}

// Filter SnapshotVersion
func Filter(vs []SnapshotVersion, f func(SnapshotVersion) bool) []SnapshotVersion {
	vsf := make([]SnapshotVersion, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
