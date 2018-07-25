package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var build string

func main() {
	app := cli.NewApp()
	app.Name = "rancher publish"
	app.Usage = "rancher publish"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access-key",
			Usage:  "AWS access key",
			EnvVar: "PLUGIN_ACCESS_KEY,ECS_ACCESS_KEY,AWS_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "AWS secret key",
			EnvVar: "PLUGIN_SECRET_KEY,ECS_SECRET_KEY,AWS_SECRET_KEY",
		},
		cli.StringFlag{
			Name:   "region",
			Usage:  "aws region",
			EnvVar: "PLUGIN_REGION",
		},
		cli.StringFlag{
			Name:   "family",
			Usage:  "ECS family",
			EnvVar: "PLUGIN_FAMILY",
		},
		cli.StringFlag{
			Name:   "task-role-arn",
			Usage:  "ECS task IAM role",
			EnvVar: "PLUGIN_TASK_ROLE_ARN",
		},
		cli.StringFlag{
			Name:   "execution-role-arn",
			Usage:  "ECS execution IAM role",
			EnvVar: "PLUGIN_EXECUTION_ROLE_ARN",
		},
		cli.StringFlag{
			Name:   "service",
			Usage:  "Service to act on",
			EnvVar: "PLUGIN_SERVICE",
		},
		cli.StringFlag{
			Name:   "container-name",
			Usage:  "Container name",
			EnvVar: "PLUGIN_CONTAINER_NAME",
		},
		cli.StringFlag{
			Name:   "docker-image",
			Usage:  "image to use",
			EnvVar: "PLUGIN_DOCKER_IMAGE",
		},
		cli.StringFlag{
			Name:   "tag",
			Usage:  "AWS tag",
			EnvVar: "PLUGIN_TAG",
		},
		cli.StringFlag{
			Name:   "cluster",
			Usage:  "AWS ECS cluster",
			EnvVar: "PLUGIN_CLUSTER",
		},
		cli.StringFlag{
			Name:   "log-driver",
			Usage:  "The log driver to use for the container",
			EnvVar: "PLUGIN_LOG_DRIVER",
		},
		cli.StringSliceFlag{
			Name:   "log-options",
			Usage:  "The configuration options to send to the log driver",
			EnvVar: "PLUGIN_LOG_OPTIONS",
		},
		cli.StringSliceFlag{
			Name:   "port-mappings",
			Usage:  "ECS port maps",
			EnvVar: "PLUGIN_PORT_MAPPINGS",
		},
		cli.StringSliceFlag{
			Name:   "labels",
			Usage:  "A key/value map of labels to add to the container",
			EnvVar: "PLUGIN_LABELS",
		},
		cli.StringSliceFlag{
			Name:   "environment-variables",
			Usage:  "ECS environment-variables",
			EnvVar: "PLUGIN_ENVIRONMENT_VARIABLES",
		},
		cli.StringSliceFlag{
			Name:   "secret-environment-variables",
			Usage:  "Secret ECS environment-variables",
			EnvVar: "PLUGIN_SECRET_ENVIRONMENT_VARIABLES",
		},
		cli.Int64Flag{
			Name:   "cpu",
			Usage:  "The number of cpu units to reserve for the container",
			EnvVar: "PLUGIN_CPU",
		},
		cli.Int64Flag{
			Name:   "memory",
			Usage:  "The hard limit (in MiB) of memory to present to the container",
			EnvVar: "PLUGIN_MEMORY",
		},
		cli.Int64Flag{
			Name:   "memory-reservation",
			Usage:  "The soft limit (in MiB) of memory to reserve for the container. Defaults to 128",
			Value:  128,
			EnvVar: "PLUGIN_MEMORY_RESERVATION",
		},
		cli.StringFlag{
			Name:   "network-mode",
			Usage:  "The Docker networking mode to use for the containers in the task. Defaults to bridge if unspecified",
			EnvVar: "PLUGIN_TASK_NETWORK_MODE",
		},
		cli.StringFlag{
			Name:   "deployment-configuration",
			Usage:  "Deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks",
			EnvVar: "PLUGIN_DEPLOYMENT_CONFIGURATION",
		},
		cli.Int64Flag{
			Name:   "desired-count",
			Usage:  "The number of instantiations of the specified task definition to place and keep running on your cluster",
			EnvVar: "PLUGIN_DESIRED_COUNT",
		},
		cli.BoolTFlag{
			Name:   "yaml-verified",
			Usage:  "Ensure the yaml was signed",
			EnvVar: "DRONE_YAML_VERIFIED",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Key:                     c.String("access-key"),
		Secret:                  c.String("secret-key"),
		Region:                  c.String("region"),
		Family:                  c.String("family"),
		TaskRoleArn:             c.String("task-role-arn"),
		ExecutionRoleArn:        c.String("execution-role-arn"),
		Service:                 c.String("service"),
		ContainerName:           c.String("container-name"),
		DockerImage:             c.String("docker-image"),
		Tag:                     c.String("tag"),
		Cluster:                 c.String("cluster"),
		LogDriver:               c.String("log-driver"),
		LogOptions:              c.StringSlice("log-options"),
		PortMappings:            c.StringSlice("port-mappings"),
		Environment:             c.StringSlice("environment-variables"),
		SecretEnvironment:       c.StringSlice("secret-environment-variables"),
		Labels:                  c.StringSlice("labels"),
		CPU:                     c.Int64("cpu"),
		Memory:                  c.Int64("memory"),
		MemoryReservation:       c.Int64("memory-reservation"),
		NetworkMode:             c.String("network-mode"),
		DeploymentConfiguration: c.String("deployment-configuration"),
		DesiredCount:            c.Int64("desired-count"),
		YamlVerified:            c.BoolT("yaml-verified"),
	}
	fmt.Println("Desired count is ", plugin.DesiredCount)
	return plugin.Exec()
}
