package main

import (
	"context"
	"flag"
	"fmt"
	v1 "github.com/zcking/steggy/gen/proto/go/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {
	host := flag.String("host", "127.0.0.1", "host of steggy server")
	port := flag.Int("port", 8080, "port on steggy server")
	msg := flag.String("msg", "", "message to encode")
	file := flag.String("file", "", "image file to encode/decode")
	decode := flag.Bool("decode", false, "decode instead of encode")
	flag.Parse()

	// Read the image file bytes
	img, err := os.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	// connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("connecting to steggy at %s\n", addr)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewSteggyServiceClient(conn)

	if !*decode {
		log.Printf("encoding message into %s ...\n", *file)
		resp, err := client.Encode(ctx, &v1.EncodeRequest{
			Message: *msg,
			Image:   img,
		})
		if err != nil {
			log.Fatal(err)
		}

		// Write out the encoded image to file system
		outFileName := fmt.Sprintf("encoded_%s", *file)
		if err = os.WriteFile(outFileName, resp.GetEncodedImage(), 0666); err != nil {
			log.Fatal(err)
		}

		log.Printf("encoded image written: %s\n", outFileName)
	} else {
		log.Printf("decoding message from %s ...\n", *file)
		resp, err := client.Decode(ctx, &v1.DecodeRequest{
			Image: img,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("decoded: %s\n", resp.GetDecodedMessage())
	}
}
