package main

import (
    "fmt"
    "os"
    "strings"
    "io/ioutil"
    "net/http"
    "encoding/xml"
    "github.com/javicv/mvn_resolver/xml"
)

func main() {

    var REPO, is_set = os.LookupEnv("MAVEN_REPO_URL")
    if !is_set {
        fmt.Fprintf(os.Stderr, "Error: Environment variable MAVEN_REPO_URL must be set\n")
        os.Exit(1)
    }

    if len(os.Args)<5 {
        fmt.Fprintf(os.Stderr, "Error: Not enough parameters provided\n")
        fmt.Fprintf(os.Stderr, "Uso: %s GROUP_ID ARTIFACT_ID VERSION PACKAGING\n", os.Args[0])
        os.Exit(2)
    }
    var group_id = strings.Split(os.Args[1], ".")
    var artifact_id = os.Args[2]
    var version = os.Args[3]
    var packaging = os.Args[4]

    var metadata_url = fmt.Sprintf("%s/%s/%s/%s/maven-metadata.xml", REPO, strings.Join(group_id,"/"),artifact_id,version)
    resp, err := http.Get(metadata_url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(3)
    }
    if resp.StatusCode != 200 {
        fmt.Fprintf(os.Stderr, "%s: %s\n", resp.Status, metadata_url)
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

    for _,snapshotVersion := range metadata.Versioning.SnapshotVersions.SnapshotVersion {
        if snapshotVersion.Extension == packaging {
            var artifact_url = fmt.Sprintf("%s/%s/%s/%s/%s-%s.%s", REPO, strings.Join(group_id,"/"), artifact_id, version, artifact_id, snapshotVersion.Value, snapshotVersion.Extension)
            fmt.Print(artifact_url)
        }
    }
}
