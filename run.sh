#!/bin/zsh

go build -o bookings cmd/web/*.go
./bookings -dbname=bookings -dbuser=lottoley -cache=false -production=false