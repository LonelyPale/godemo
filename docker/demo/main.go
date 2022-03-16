package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"os"
	"strconv"
	"time"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	cid := ""
	for _, cobj := range containers {
		fmt.Printf("%s %s %s\n", cobj.ID[:10], cobj.Names, cobj.Image)
		//fmt.Println(cobj.Names[0])
		// idpass:genkey hello-world redis
		//if cobj.Image == "idpass:genkey" {
		//	cid = cobj.ID
		//}
		if cobj.Names[0] == "/genkey-test" {
			cid = cobj.ID
		}
	}
	println(cid)

	//time.RFC3339
	now := time.Now()
	fmt.Println("now:", now.Unix(), now)

	if err := cli.ContainerStart(context.Background(), cid, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	//time.Sleep(time.Second * 3)

	out, err := cli.ContainerLogs(context.Background(), cid, types.ContainerLogsOptions{ShowStdout: true, Timestamps: true, Since: strconv.Itoa(int(now.Unix()))})
	if err != nil {
		panic(err)
	}

	n, err := stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	if err != nil {
		fmt.Println(n, err)
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(context.Background(), cid, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
		fmt.Println(time.Now().Local(), "exited")
		return
	}

	fmt.Println(time.Now().Local(), "quitted")
	select {}
}
