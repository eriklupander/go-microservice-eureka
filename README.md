# A sample go microservice with Eureka integration

This very simplistic little http service provides hard-coded 'vendors' on /vendors/{id} and is used as an example project for a blog post on http://callistaenterprise.se/blogg/teknik/

## Building

Go SDK 1.5+ installed.

Clone this repo, set your GOPATH / GOROOT to the folder you cloned into.

    cd src/github.com/eriklupander/go-eureka/*.go
    go get

The docker image can be built using the .sh script:

    #!/usr/bin/env bash
    
    export GOARCH=amd64
    export GOOS=linux
    go build -o bin/goeureka src/github.com/eriklupander/goeureka/*.go
    docker build -t vendor .
    export GOARCH=amd64
    export GOOS=darwin

Update for your target OS accordingly.

## Running

This microservice can be ran standalone, though it will keep trying to contact the Eureka server at http://192.168.99.100:8761 until exit or success. I.e - this code is meant to be executed in the context of the blog post.

Anyway - you can hopefully start it using go run:

     > go run src/github.com/eriklupander/go-eureka/*.go

Have fun!

This software uses the MIT license, see LICENSE.md

Uses gorilla/context and gorilla/mux Copyright (c) 2012 Rodrigo Moraes
Uses twinj/uuid Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>, Copyright (C) 2014 by Daniel Kemp <twinj@github.com> Derivative work