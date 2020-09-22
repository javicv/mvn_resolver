# Maven Artifact URL Resolver

![License](https://img.shields.io/github/license/javicv/mvn-resolver.svg)
[![Build Status](https://travis-ci.com/javicv/mvn-resolver.svg?branch=master)](https://travis-ci.com/javicv/mvn-resolver)
![GitHub release](https://img.shields.io/github/release/javicv/mvn-resolver.svg)

This small application is designed to resolve the URL of the last snapshot uploaded to a maven repository for a given version. The URL will be printed on the *standard output*.

## Use

An Environment Variable ***MAVEN_REPO_URL*** with the Maven Repository base URL must be defined

If ***MAVEN_USERNAME*** and ***MAVEN_PASSWORD*** environment variables are set, connection with Maven repository will use Basic Athentication

## Example
* Public Maven Repository

$ `MAVEN_REPO_URL=https://example.com/repo/snapshots ./mvn_resolver com.example.proj my-artifact 0.0.1-SNAPSHOT jar`

* Basic Authentication

$ `export MAVEN_USERNAME=myuser`

$ `export MAVEN_PASSWORD=password`

$ `MAVEN_REPO_URL=https://example.com/repo/snapshots ./mvn_resolver com.example.proj my-artifact 0.0.1-SNAPSHOT jar`