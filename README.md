# Maven URL Resolver
[![Build Status](https://travis-ci.org/javicv/mvn_resolver.svg?branch=master)](https://travis-ci.org/javicv/mvn_resolver)

This small application is designed to resolve the URL of the last snapshot uploaded to a maven repository for a given version. The URL will be printed on the *standard output*.

## Use
An Environment Variable ***MAVEN_REPO*** with the Maven Repository base URL must be defined

## Example
`MAVEN_REPO=https://example.com/repo/snapshots ./mvn_resolver com.example.proj my-artifact 0.0.1-SNAPSHOT jar`
