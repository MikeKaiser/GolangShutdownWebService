// compile with "go build shutdownrestart.go"

package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func execute(command string) {
	out, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	if out != nil {
		fmt.Printf("out: %s\n", out)
	}
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
}

func executeNonBlocking(command string) {
	err := exec.Command("/bin/sh", "-c", command).Run()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=/shutdown>Shutdown RPi</a>\n<a href=/restart>Restart RPi</a>\n")
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	executeNonBlocking("sudo shutdown -h now &")
	fmt.Fprintf(w, "Shutting Down")
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	executeNonBlocking("sudo shutdown -r now &")
	fmt.Fprintf(w, "Restarting")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/shutdown", shutdownHandler)
	http.HandleFunc("/restart", restartHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
