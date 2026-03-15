#!/bin/bash
# LOG_LEVEL=debug ENV=dev (default) — text format, debug level
# ENV=prod — JSON format, info level
go run *.go "$@"
