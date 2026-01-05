#!/bin/bash
# golangci-lint avtomatik ishga tushirish skripti
export PATH=$PATH:~/go/bin
golangci-lint run --fix ./day-01/ ./day-02/
