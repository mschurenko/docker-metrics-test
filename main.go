package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ", os.Args[0], "<container_id> [<client_version]")
		os.Exit(1)
	}

	clientVer := "1.38"

	if len(os.Args) == 3 {
		clientVer = os.Args[2]
	}

	cli, err := client.NewClientWithOpts(client.WithVersion(clientVer))
	if err != nil {
		log.Fatal(err)
	}

	id := os.Args[1]

	stats, err := cli.ContainerStats(context.Background(), id, false)
	if err != nil {
		log.Fatal(err)
	}

	var parsed types.StatsJSON

	err = json.NewDecoder(stats.Body).Decode(&parsed)
	if err != nil {
		log.Fatal(err)
	}
	defer stats.Body.Close()

	// fmt.Println(parsed.MemoryStats)

	bs, err := json.MarshalIndent(parsed.MemoryStats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

}
