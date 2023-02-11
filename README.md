# Steggy

![Stegosaurus](./steggy.png)

Steggy is a gRPC server for using steganography to encode hidden messages into PNG images!

> [Steganography](https://en.wikipedia.org/wiki/Steganography) is the practice of 
> representing information within another message or physical object, in such a manner 
> that the presence of the information is not evident to human inspection. 
> In computing/electronic contexts, a computer file, message, image, or video is 
> concealed within another file, message, image, or video.  
> [[Wiki Source]](https://en.wikipedia.org/wiki/Steganography)

This project is a simple demonstration of gRPC, using a topic that I thought was fun and 
interesting. That's it. No "production" documentation and no long-term maintenance of this repo 
guaranteed. 

If you think this project is neat and want to fork it and make your own twists, here are 
a few ideas to recommend:  

* Create a standalone CLI instead of the server/client pattern
* Use gRPC bi-directional streaming to support HUGE images
* Support for video
* Support for ANY file type

> **Fun Fact:** the image provided in this repo and shown above, `steggy.png` is in fact already
> encoded using steggy. Go ahead, try decoding it and find the hidden message for yourself!

---

## Build

To build steggy from source:  

```shell
go build -o build/ ./...
```

The final executables can now be found in the `build/` directory.

## Usage - Server

To run the steggy server:  

```shell
./build/server
```

For example:  

```
❯ ./build/server 
2023/02/10 22:19:38 starting server on port 8080...
2023/02/10 22:19:38 server started on :8080
```

## Usage - Client

Once the server is running, you may use it to encode messages 
into a small PNG image, or decode a PNG image to extract the 
hidden message.

### Encoding a message

To encode a message into an image, use `./build/client -file <file> -msg <msg>`:  

```shell
❯ ./build/client -file steggy.png -msg 'Hello World!'
2023/02/10 22:21:37 connecting to steggy at 127.0.0.1:8080
2023/02/10 22:21:37 encoding message into steggy.png ...
2023/02/10 22:21:38 encoded image written: encoded_steggy.png
```

### Decoding an image

To decode an image and get its hidden message, use `./build/client -file <file> -decode>`:

```shell
❯ ./build/client -file encoded_steggy.png -decode
2023/02/10 22:20:14 connecting to steggy at 127.0.0.1:8080
2023/02/10 22:20:14 decoding message from encoded_steggy.png ...
2023/02/10 22:20:14 decoded: Hello World!
```

