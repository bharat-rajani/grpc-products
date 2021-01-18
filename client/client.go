package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "github.com/bharat-rajani/grpc-products-demo/gen/proto/products"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	addr = flag.String("addr", "localhost", "The address of the server to connect to")
	port = flag.String("port", "8080", "The port to connect to")
)

var LetterRunes []rune = []rune("3ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing command: getprodtypes or getprods")
		os.Exit(1)
	}

	// creds, err := credentials.NewClientTLSFromFile("../cert/service.pem", "")
	// if err != nil {
	// 	log.Fatalf("could not process the credentials: %v", err)
	// }

	conn, err := grpc.Dial(net.JoinHostPort(*addr, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server:, %s", err)

	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()

	switch cmd := flag.Arg(0); cmd {
	case "getprodtypes":
		err = getprodtypes(ctx, client, flag.Arg(1))
	case "getprods":
		err = getprods(ctx, client, flag.Arg(1), flag.Arg(2))
	case "setprods":
		err = setprods(ctx, client, flag.Arg(1), flag.Arg(2))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func getprodtypes(ctx context.Context, client pb.ProductServiceClient, vendor string) error {

	log.Printf("requesting all product types from vendor: %s", vendor)

	if vendor == "" {
		return fmt.Errorf("Vendor arg is missing, select between available cloud vendors: google, aws, oracle")
	}

	requestProdType := pb.ClientRequestType{
		Vendor: vendor,
	}

	// ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	// defer cancel()

	response, err := client.GetVendorProductTypes(ctx, &requestProdType)
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			return status.Errorf(errStatus.Code(), "error while calling client.GetVendorProdTypes() method: %v ", errStatus.Message())
		}
		return fmt.Errorf("Could not get the products: %v", err)
	}

	fmt.Printf("%s cloud products type are: %s\n", vendor, response.GetProductType())

	return nil

}

func getprods(ctx context.Context, client pb.ProductServiceClient, vendor string, prodType string) error {

	log.Printf("requesting all %s products from %s", prodType, vendor)

	if vendor == "" || prodType == "" {
		return fmt.Errorf("You need both, vendor and prodType args. Example command: $client oracle storage")
	}

	requestProd := pb.ClientRequestProducts{
		Vendor:      vendor,
		ProductType: prodType,
	}

	stream, err := client.GetVendorProducts(ctx, &requestProd)
	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			return status.Errorf(errStatus.Code(), "error while calling client.GetVendorProds() method: %v ", errStatus.Message())
		}
		return fmt.Errorf("Could not get the stream of products : %v", err)
	}

	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if errStatus, ok := status.FromError(err); ok {
				return status.Errorf(errStatus.Code(), "error while receiving the stream for client.GetVendorProds: %v ", errStatus.Message())
			}
			return fmt.Errorf("error while receiving the stream for client.GetVendorProds: %v", err)
		}
		fmt.Printf("Title: %s, Url: %s,  ShortUrl: %s\n", product.GetProduct().GetTitle(), product.GetProduct().GetUrl(), product.GetProduct().GetShortUrl())
	}

	return nil
}

func setprods(ctx context.Context, client pb.ProductServiceClient, vendor string, prodType string) error {

	log.Printf("setting  %s products from %s", prodType, vendor)

	if vendor == "" || prodType == "" {
		return fmt.Errorf("You need both, vendor and prodType args. Example command: $client oracle storage")
	}

	stream, err := client.SetVendorProducts(ctx)
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}

	totalL := 0
	for {
		// fmt.Printf("\nEnter number of products to set: ")
		var cnt int = 12
		// fmt.Scanf("%d\n", &cnt)

		productNames := make([]string, 0, cnt)

		for i := 0; i < cnt; i++ {
			productNames = append(productNames, genRandomStr(3))
		}

		if len(productNames) == 0 || productNames[0] == "" {
			fmt.Println("Closing stream")
			if err := stream.CloseSend(); err != nil {
				log.Println(err)
				return err
			}
		}

		totalL += len(productNames)

		time.Sleep(time.Millisecond * 400)
		for _, prod := range productNames {
			fmt.Println("Setting product: ", prod)
			id := uuid.Must(uuid.NewRandom()).String()

			requestProd := pb.AdminClientRequestProducts{
				Product: &pb.ProdsPrep{
					Title:    prod,
					Url:      "sample Url",
					ShortUrl: "https://made-up-url.com/" + id[:6],
				},
				Vendor:      vendor,
				ProductType: prodType,
			}

			// fmt.Println(requestProd)
			// time.Sleep(time.Millisecond * 600)

			if err := stream.Send(&requestProd); err != nil {
				log.Println("Error while sending: %v", err)
				log.Println("Total products sent: ", totalL)
				return err
			}

			out, err := json.Marshal(requestProd)
			if err != nil {
				log.Println(err.Error())
			}

			log.Printf("%v Sent", string(out))
		}

	}

	return nil
}

func genRandomStr(length int) string {
	strArr := make([]rune, length)
	for i := range strArr {
		strArr[i] = LetterRunes[rand.Intn(len(LetterRunes))]
	}
	return string(strArr)
}
