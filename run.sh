#!/bin/bash
go build -o urlshortener cmd/web/*.go && ./urlshortener
