package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	pb "github.com/bharat-rajani/grpc-products-demo/gen/proto/products"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var allProducts map[string]map[string][]string = map[string]map[string][]string{
	"google": {
		"compute": []string{"App Engine", "Cloud Run", "App Engine"},
		"storage": []string{"Cloud Storage", "Filestore"},
	},
	"aws": {
		"compute": []string{"ECS", "EKR", "AWS Fargate"},
		"storage": []string{"Amazon Aurora", "Amazon RDS", "Amazon Redshift"},
	},
	"oracle": {
		"compute": []string{"VM", "Bare Metal"},
		"storage": []string{"Oracle ZFS", "Oracle StorageTek"},
	},
}

var commonProdChan chan *pb.AdminClientRequestProducts = make(chan *pb.AdminClientRequestProducts, 10)

type ProductServer struct {
	ProducTypes map[string][]string
	pb.UnimplementedProductServiceServer
}

func (pserv *ProductServer) GetVendorProductTypes(ctx context.Context, req *pb.ClientRequestType) (*pb.ClientResponseType, error) {

	log.Printf("have received a request for -> %s <- as vendor", req.GetVendor())

	time.Sleep(5 * time.Second)
	var prodTypes []string
	if ctx.Err() == context.Canceled {
		log.Print("the user has canceled the request, stoping server side operation")
		return nil, status.Error(codes.Canceled, "the user has canceled the request, stoping server side operation")
	}

	log.Println("Prepairing reponse...")
	if vendorProductTypes, found := pserv.ProducTypes[req.GetVendor()]; found {

		for _, prodType := range vendorProductTypes {
			prodTypes = append(prodTypes, req.GetVendor()+" "+prodType)
		}

	} else {
		return nil, status.Error(codes.InvalidArgument, "Wrong vendor, select between google, aws, oracle")
	}

	clientResponse := pb.ClientResponseType{
		ProductType: strings.Join(prodTypes, ","),
	}

	// if ctx.Err() == context.DeadlineExceeded {
	// 	//	log.Printf("dealine has exceeded, stoping server side operation")
	// 	return nil, status.Error(codes.DeadlineExceeded, "dealine has exceeded, stoping server side operation")
	// }
	log.Println("Sending reponse...")

	return &clientResponse, nil
}

func (pserv *ProductServer) GetVendorProducts(req *pb.ClientRequestProducts, stream pb.ProductService_GetVendorProductsServer) error {

	log.Printf("have received a request for -> %s <- product type from -> %s <- vendor", req.GetProductType(), req.GetVendor())

	// log.Printf("fetch response for id : %d", in.Id)
	ctx := stream.Context()

	productChan := make(chan string, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go readProducts(ctx, wg, productChan, req.GetVendor(), req.GetProductType())
	for {

		product := <-productChan
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("dealine has exceeded, stoping server side operation")
			return status.Error(codes.DeadlineExceeded, "Deadline execeeded, stopping..")
		}

		if ctx.Err() == context.Canceled {
			log.Print("the user has canceled the request, stoping server side operation")
			return status.Error(codes.Canceled, "User cancelled, stopping...")
		}

		// time.Sleep(time.Duration(1) * time.Second)

		id := uuid.Must(uuid.NewRandom()).String()

		if err := stream.Send(&pb.ClientResponseProducts{
			Product: &pb.ProdsPrep{
				Title:    product,
				Url:      "sampleUrl",
				ShortUrl: "https://made-up-url.com/" + id[:6],
			},
		}); err != nil {
			return err
		}

	}
	wg.Wait()
	log.Printf("the response was sent to client")
	return nil
}

func (pserv *ProductServer) SetVendorProducts(stream pb.ProductService_SetVendorProductsServer) error {

	// log.Printf("have received a request for -> %s <- product type from -> %s <- vendor", req.GetProductType(), req.GetVendor())

	// log.Printf("fetch response for id : %d", in.Id)
	ctx := stream.Context()
	//use wait group to allow process to be concurrent
	// var wg sync.WaitGroup

	startTime := time.Now()

	var productCnt int32

	for {
		product, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			log.Printf("Elapsed Time (sec): %d", int32(endTime.Sub(startTime).Seconds()))
			return stream.SendAndClose(&pb.ProductCount{
				Count: productCnt,
			})
		}
		if err != nil {
			return err
		}

		if ctx.Err() == context.Canceled {
			log.Print("the user has canceled the request, stoping server side operation")
			return status.Error(codes.Canceled, "Client cancelled connection.")
		}

		productCnt++
		saveProduct(product)
	}
}

func (s *ProductServer) ChatVendorSales(stream pb.ProductService_ChatVendorSalesServer) error {
	ctx := stream.Context()
	startTime := time.Now()
	for {
		msg, err := stream.Recv()
		fmt.Println("Received: ", msg)
		if err == io.EOF {
			endTime := time.Now()
			log.Printf("Elapsed Time (sec): %d", int32(endTime.Sub(startTime).Seconds()))
			return stream.Send(&pb.ChatMessage{MessageContent: "goodbye"})
		}
		if err != nil {
			return err
		}

		if ctx.Err() == context.Canceled {
			log.Print("the user has canceled the request, stoping server side operation")
			return status.Error(codes.Canceled, "Client cancelled connection.")
		}

	}
}

func NewProductServer(productTypes map[string][]string) *ProductServer {
	return &ProductServer{ProducTypes: productTypes}
}

func saveProduct(product *pb.AdminClientRequestProducts) {
	log.Printf("Saving prdouct %v\n", product)
	commonProdChan <- product
}

func readProducts(ctx context.Context, wg sync.WaitGroup, productChan chan<- string, vendor string, productType string) {

	for _, product := range allProducts[vendor][productType] {
		productChan <- product
	}

	for {
		select {
		case c := <-commonProdChan:
			if c.Vendor == vendor && c.ProductType == productType {
				productChan <- c.Product.Title
			}

			// case <-ctx.Done():
			// 	log.Println("Request done/cancelled.")
			// 	wg.Done()
		}
	}
	wg.Done()

}
