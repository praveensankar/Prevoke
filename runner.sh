#!/bin/sh

runIssuer = "go run  github.com/praveensankar/Revocation-Service -issuer"
runVerifier = "go run github.com/praveensankar/Revocation-Service -verifier"
runHolder = "go run github.com/praveensankar/Revocation-Service -holder"

currentDirectory=$(pwd)

# Open a new terminal window and execute commands
osascript -e "tell application \"Terminal\"" \
          -e "do script \"cd '$currentDirectory' && go run  main.go -issuer\"" \
          -e "activate" \
          -e "end tell" &

osascript -e "tell application \"Terminal\"" \
          -e "do script \"cd '$currentDirectory' && go run  main.go -verifier\"" \
          -e "activate" \
          -e "end tell" &


osascript -e "tell application \"Terminal\"" \
          -e "do script \"cd '$currentDirectory' && go run  main.go -holder\"" \
          -e "activate" \
          -e "end tell" &

