package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "abtest/abtest_proto"

	"google.golang.org/grpc"
)

var (
	// GRPCAddress is the remote grpc server address
	GRPCAddress string = "127.0.0.1:9527"

	// Users is the users number
	Users int = 100

	// HashKey is the hash key
	HashKey string = ""
)

var (
	userLabPath = make(map[string][]string, 0)
	offCount    = 0
	onENCount   = 0
	onCNENCount = 0
)

func init() {
	flag.IntVar(&Users, "users", 100, "参与测试的用户总数")
	flag.StringVar(&GRPCAddress, "grpc", "127.0.0.1:9527", "ABTest server GRPC server address")
	flag.StringVar(&HashKey, "key", "", "HashKey for once AB test")
}

func main() {
	// AB test
	flag.Parse()

	if HashKey != "" {
		serviceSubtitleLayerSwitch(HashKey)
	}

	for i := 0; i < Users; i++ {
		serviceSubtitleLayerSwitch(fmt.Sprintf("userID00%d", i))
	}

	for k, v := range userLabPath {
		log.Println(k, v)
	}

	fmt.Println("offCount", offCount)
	fmt.Println("onENCount", onENCount)
	fmt.Println("onCNENCount", onCNENCount)
}

func serviceSubtitleLayerSwitch(user string) {
	conn, err := grpc.Dial(GRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial error: %+v", err)
	}
	defer conn.Close()

	abtClient := pb.NewABTestServiceClient(conn)
	ctx := context.Background()

	resp, err := abtClient.GetABTestZone(ctx, &pb.GetABTestZoneRequest{
		ProjectId: "Subtitle",
		LayerId:   "Switch",
		HashKey:   user,
	})

	if err != nil {
		log.Fatalf("abtClient.GetABTestZone error: %+v", err)
	}

	// Push ab test data
	defer abtClient.PushABTestData(ctx, &pb.PushABTestDataRequest{
		ProjectId: "Subtitle",
		HashKey:   user,
	})
	userLabPath[user] = append(userLabPath[user], resp.GetZone().GetLabel())
	switch resp.GetZone().GetLabel() {
	case "ON":
		fmt.Printf("The user goes into Lab %s\n", resp.GetZone().GetLabel())

		// go to language
		serviceSubtitleLayerLanguage(user)
	case "OFF":
		fmt.Printf("The user goes into Lab %s\n", resp.GetZone().GetLabel())

		offCount++

		// go to language by wrong
		// serviceSubtitleLayerLanguage(user)
	default:
		fmt.Printf("The user goes into defualt and the label is [%s\n", resp.GetZone().GetLabel()+"]")
	}
}

func serviceSubtitleLayerLanguage(user string) {
	conn, err := grpc.Dial(GRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial error: %+v", err)
	}
	defer conn.Close()

	abtClient := pb.NewABTestServiceClient(conn)
	ctx := context.Background()

	resp, err := abtClient.GetABTestZone(ctx, &pb.GetABTestZoneRequest{
		ProjectId: "Subtitle",
		LayerId:   "Language",
		HashKey:   user,
	})

	if err != nil {
		log.Fatalf("abtClient.GetABTestZone error: %+v", err)
	}
	userLabPath[user] = append(userLabPath[user], resp.GetZone().GetLabel())
	switch resp.GetZone().GetLabel() {
	case "CN_EN":
		onCNENCount++
		fmt.Printf("The user goes into Lab %s\n", resp.GetZone().GetLabel())
	case "EN":
		onENCount++
		fmt.Printf("The user goes into Lab %s\n", resp.GetZone().GetLabel())
	default:
		fmt.Printf("The user goes into defualt and the label is [%s\n", resp.GetZone().GetLabel()+"]")
	}
}

/* Subtitle 实验配置 如下：
Subtitle:
    -
        id: "ON"
        project:
            id: "Subtitle"
        layer:
            id: "Switch"
            totalweight: 100
        weight:
            min: 1
            max: 50
        label: "ON"
        description: "有字幕"
        usergroups: ["VIP","Frequent"]
    -
        id: "OFF"
        project:
            id: "Subtitle"
        layer:
            id: "Switch"
            totalweight: 100
        weight:
            min: 51
            max: 100
        label: "OFF"
        description: "无字幕"
        usergroups: ["VIP","Frequent"]
    -
        id: "CN_EN"
        project:
            id: "Subtitle"
        layer:
            id: "Language"
            totalweight: 100
            parentzones:
                -
                    id: "ON"
                    project:
                        id: "Subtitle"
                    layer:
                        id: "Switch"
                        totalweight: 100
                    weight:
                        min: 1
                        max: 50
                    label: "ON"
                    description: "有字幕"
                    usergroups: ["VIP","Frequent"]
        weight:
            min: 1
            max: 50
        label: "CN_EN"
        description: "中英字幕"
    -
        id: "EN"
        project:
            id: "Subtitle"
        layer:
            id: "Language"
            totalweight: 100
            parentzones:
                -
                    id: "ON"
                    project:
                        id: "Subtitle"
                    layer:
                        id: "Switch"
                        totalweight: 100
                    weight:
                        min: 1
                        max: 50
                    label: "ON"
                    description: "有字幕"
                    usergroups: ["VIP","Frequent"]
        weight:
            min: 51
            max: 100
        label: "EN"
        description: "英文字幕"
*/
