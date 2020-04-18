#!/bin/bash

go build -o integeration integeration.go
go install .
rm -rf integeration
sleep 1
jobdsigner
