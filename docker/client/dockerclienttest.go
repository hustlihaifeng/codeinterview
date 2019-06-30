package main

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/docker/docker/api/types"
	dockercli "github.com/docker/docker/client"
)

func main() {
	localImageSummarySli, err := DockerImageList("v1.39")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", localImageSummarySli)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	err = DockerPushImage("172.17.0.2:5000/hello-world:tag1", "v1.39")
	if err != nil {
		panic(err)
	}
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

func DockerPushImage(nameAndTag string, dockerVersion string) (err error) {
	defer func() {
		fmt.Printf("%v return %v\n", FileLineFunc(1), err)
	}()
	cli, err := dockercli.NewEnvClient()
	if err != nil {
		return err
	}

	err = dockercli.WithVersion(dockerVersion)(cli)
	if err != nil {
		return err
	}
	defer cli.Close()
	_, err = cli.ImagePush(context.Background(), nameAndTag, types.ImagePushOptions{RegistryAuth: "NotValid"})
	if err != nil {
		beego.Warn(err)
		return err
	}
	fmt.Println("push success")
	return nil
}

func FileLineFunc(offset int) string {
	if offset < 0 {
		offset = 0
	}
	depth := 1 + offset
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		return "FileLineFunc"
	}

	baseFile := filepath.Base(file)
	name := runtime.FuncForPC(pc).Name()

	return fmt.Sprintf("%s:%v\t%s", baseFile, line, name)
}
