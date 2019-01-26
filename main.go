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

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.Fatal(err)
	}

	// containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var id string
	// for _, container := range containers {
	// 	// fmt.Printf("%s %s\n", container.ID[0:10], container.Status)
	// 	id = container.ID
	// }

	if len(os.Args) != 2 {
		fmt.Println("usage: ", os.Args[0], "<container_id>")
		os.Exit(1)
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
