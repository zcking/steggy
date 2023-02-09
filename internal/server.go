package internal

import (
	"bytes"
	"context"
	"github.com/auyer/steganography"
	v1 "github.com/zcking/steggy/gen/proto/go/api/v1"
	"image/png"
	"log"
)

type Server struct {
	v1.UnimplementedSteggyServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) Encode(_ context.Context, req *v1.EncodeRequest) (*v1.EncodeResponse, error) {
	reader := bytes.NewReader(req.GetImage())
	img, err := png.Decode(reader)
	if err != nil {
		log.Printf("ERR: %v\n", err)
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = steganography.Encode(buf, img, []byte(req.GetMessage())); err != nil {
		log.Printf("ERR: %v\n", err)
		return nil, err
	}

	return &v1.EncodeResponse{
		EncodedImage: buf.Bytes(),
	}, nil
}

func (s *Server) Decode(_ context.Context, req *v1.DecodeRequest) (*v1.DecodeResponse, error) {
	reader := bytes.NewReader(req.GetImage())
	img, err := png.Decode(reader)
	if err != nil {
		log.Printf("ERR: %v\n", err)
		return nil, err
	}

	// retrieving message size to decode in the next line
	sizeOfMessage := steganography.GetMessageSizeFromImage(img)

	// decoding the message from the file
	msg := steganography.Decode(sizeOfMessage, img)
	return &v1.DecodeResponse{
		DecodedMessage: string(msg),
	}, nil
}
