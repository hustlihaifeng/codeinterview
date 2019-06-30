package utils

import (
	"context"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/docker/docker/api/types"
	dockercli "github.com/docker/docker/client"
)

func DockerPushImage(nameAndTag string, dockerVersion string) error {
	cli, err := dockercli.NewEnvClient()
	if err != nil {
		return err
	}

	dockercli.WithVersion(dockerVersion)(cli)
	defer cli.Close()
	_, err = cli.ImagePush(context.Background(), nameAndTag, types.ImagePushOptions{})
	if err != nil {
		beego.Warn(err)
		return err
	}
	return nil
}

func DockerImageList(dockerVersion string) ([]types.ImageSummary, error) {
	cli, err := dockercli.NewEnvClient()
	if err != nil {
		return nil, err
	}

	dockercli.WithVersion(dockerVersion)(cli)
	defer cli.Close()

	return cli.ImageList(context.Background(), types.ImageListOptions{})
}

func DockerTagImage(oldTagName, newTagName string) error {
	cmd := fmt.Sprintf("docker tag %s %s", oldTagName, newTagName)
	output, err := ExecuteCommandWithOutput(cmd)
	beego.Info(fmt.Sprintf(cmd))
	beego.Info(fmt.Sprintf(output))
	if err != nil {
		return err
	}
	return nil
}

func DockerRmImage(imageName string, force bool) error {
	cmd := fmt.Sprintf("docker image rmi %s", imageName)
	if force {
		cmd = cmd + " -f"
	}
	output, err := ExecuteCommandWithOutput(cmd)
	beego.Info(fmt.Sprintf(cmd))
	beego.Info(fmt.Sprintf(output))
	if err != nil {
		beego.Error("ErrorDockerRmi ", err)
		return err
	}
	return nil
}
