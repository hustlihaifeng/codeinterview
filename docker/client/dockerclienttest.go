package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	dockercli "github.com/docker/docker/client"
)

func main() {
	localImageSummarySli, err := DockerImageList("v1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", localImageSummarySli)
	return
}

// DockerImageList 获取本地镜像列表，输出是image name和tag的数组
func DockerImageList(dockerVersion string) ([]types.ImageSummary, error) {
	cli, err := dockercli.NewEnvClient()
	if err != nil {
		return nil, err
	}

	dockercli.WithVersion(dockerVersion)(cli)
	defer cli.Close()

	return cli.ImageList(context.Background(), types.ImageListOptions{})
}
