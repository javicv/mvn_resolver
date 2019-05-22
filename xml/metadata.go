package maven

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
