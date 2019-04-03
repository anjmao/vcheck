package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/anjmao/vcheck/client"
	"log"
	"os"
	"time"
)

var (
	target          = flag.String("target", "", "-target staging.mpost-api.guardtek.net:443")
	expectedVersion = flag.String("version", "", "Expected version")
	method          = flag.String("method", "/version.Version/GetVersion", "Version check endpoint path")
	clientType      = flag.String("client", "grpc", "Version check client type")
	checkCount      = flag.Int("count", 12, "Check count")
	sleepAfterCheck = flag.Int("sleep", 5, "Sleep after check in seconds")
	showHelp        = flag.Bool("help", false, "Show help")
)

const checkTimeout = 10 * time.Second

const usageStr = `
Usage: vcheck [options]
Options:
	--target <target>				Target host including port. (e.g --target service.mydomain.com:443)
	--version <version>				Expected version (e.g -v 1.2.3)
	--method <method>				Version check endpoint (default: /debug.Debug/GetVersion)
	--client <client>				Client type (grpc, http)
	--count	<count>					Check count (default: 12)
	--sleep	<sleep>					Sleep duration after check in seconds (default: 5)

Other options:
	--help
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func main() {
	flag.Parse()
	if *showHelp || *target == "" || *expectedVersion == "" {
		usage()
	}

	sleep := time.Duration(*sleepAfterCheck) * time.Second

	var actualVersion string
	c := getClient(*clientType, *target, *method)
	for i := 0; i < *checkCount; i++ {
		v, err := getVersion(c)
		if err != nil {
			fmt.Printf("could not get version: %v\n", err)
			time.Sleep(sleep)
			continue
		}
		actualVersion = v.BuildVersion
		fmt.Printf("expected version %s, got %s\n", *expectedVersion, actualVersion)
		if actualVersion == *expectedVersion {
			fmt.Println("deployment successful")
			return
		}

		time.Sleep(sleep)
	}

	log.Fatalf("deployment failed: expected version %s, got %s\n", *expectedVersion, actualVersion)
}

func getVersion(c client.Client) (*client.GetVersionReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), checkTimeout)
	defer cancel()
	v, err := c.GetVersion(ctx)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func getClient(clientType, target, method string) client.Client {
	switch clientType {
	case "http":
		return client.NewHTTP(target, method)
	case "grpc":
		return client.NewGRPC(target, method)
	default:
		log.Fatalf("unknow client type: %s", clientType)
	}
	return nil
}
