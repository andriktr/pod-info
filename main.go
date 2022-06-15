package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var podName string = os.Getenv("POD_NAME")
	var nodeName string = os.Getenv("NODE_NAME")
	var podIp string = os.Getenv("POD_IP")
	var podNamespace string = os.Getenv("POD_NAMESPACE")
	var podServiceAccount string = os.Getenv("POD_SERVICE_ACCOUNT")
	currentTime := time.Now()

	fmt.Fprintf(w, "Time: %s\n", currentTime.Format(time.UnixDate))
	fmt.Fprintf(w, "Pod Name: %s\n", podName)
	fmt.Fprintf(w, "Pod IP: %s\n", podIp)
	fmt.Fprintf(w, "Node Name: %s\n", nodeName)
	fmt.Fprintf(w, "Pod Namespace: %s\n", podNamespace)
	fmt.Fprintf(w, "Pod Service Account: %s\n", podServiceAccount)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
