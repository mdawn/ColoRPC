# ColoRPC 

2 gRPC services utilizing channels, routines, and protobufs. Used with Evans.

## What it Does

- The unary RPC takes an adjective and a color, returning the result of your dazzling descriptive input

- The bidirectional RPC pushes out a stream of various shades of green. I'm making the argument that this RPC can bring you luck. Otherwise it's useless unless you _REALLY_ like green. 

## Simple Setup

**STEP 1**: Install Evans using homebrew</br>
`brew tap ktr0731/evans`</br>
`brew install evans`

**STEP 2**: Clone this repo. Cd in there. </br> 
Run `go run server.go`

**STEP 3**: Connect the server to Evans </br>
`evans -p 50051 -r`

## Try the Unary with Evans!

Once in Evans (having used `evans -p 50051 -r above`), I show available services:

- `show service`

I view the service I want:

- `service ColorService`

I call the RPC I want (in this case, I chose the unary):

- `call Color`

And enter my data:

- colors::adjective (TYPE_STRING) => `mellow ` 
- colors::base_color (TYPE_STRING) => `yellow`

Your object will return like so:


```
{
  "result": "mellow yellow"`
} 
```

(The Bidirectional works best with the client. For extra fun, run `go run client.go` in another terminal to get a stream of green.)