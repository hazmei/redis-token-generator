package main

import (
	"fmt"
	"os"
    "log"
    "crypto/tls"

	"github.com/build-on-aws/aws-redis-iam-auth-golang/auth"
	"github.com/redis/go-redis/v9"
)

var (
	clusterEndpoint string
	username        string
	client          *redis.ClusterClient
	region          string
	clusterName     string
)

const defaultRegion = "ap-southeast-1"

func init() {
	clusterEndpoint = os.Getenv("CLUSTER_ENDPOINT")
	username = os.Getenv("USERNAME")
	region = os.Getenv("AWS_REGION")
	clusterName = os.Getenv("CLUSTER_NAME")

	generator, err := auth.New("elasticache", clusterName, username, region)
    if err != nil {
        log.Fatal("failed to initialise token generator", err)
    }
    fmt.Println("Initialised token generator")

    token, err := generator.Generate()
    if err != nil {
        log.Fatal("failed to generate auth token", err)
    }
    fmt.Println(token)

	// client = redis.NewClusterClient(
	// 	&redis.ClusterOptions{
	// 		Username: username,
	// 		Addrs:    []string{clusterEndpoint},
	// 		NewClient: func(opt *redis.Options) *redis.Client {

	// 			return redis.NewClient(&redis.Options{
	// 				Addr: opt.Addr,
	// 				CredentialsProvider: func() (username string, password string) {

	// 					token, err := generator.Generate()
	// 					if err != nil {
	// 						log.Fatal("failed to generate auth token", err)
	// 					}

	// 					fmt.Println("auth token generated successfully")
    //                     fmt.Println(token)

	// 					return opt.Username, token
	// 				},
	// 				TLSConfig: &tls.Config{InsecureSkipVerify: true},
	// 			})
	// 		},
	// 	})

    // fmt.Println("Done creating new redis client")
    // fmt.Println(client)
	// err = client.Ping(context.Background()).Err()
	// if err != nil {
	// 	log.Fatal("failed to connect to memorydb -", err)
	// }

	// fmt.Println("successfully connected to cluster", clusterEndpoint)

	fmt.Println("Hello, init world.")
}

func main() {
	fmt.Println("Hello, main world.")
}