# Rosie Hamilton

I have written the client library in a separate Go package called `apiclient`

# Design

I have chosen to use functions to query the end points `create`, `fetch`, `list` and `delete` 
Each of these functions is in a separate Go file, this is to increase readability and aid debugging. 
The first solution I came up with used methods on a Client struct to query the endpoints, but I changed my design to using functions instead.
This allowed me to pass by value instead of passing by reference which in turn allowed me to keep the code cleaner and make it more readable.

The core Go team did not set any timeouts on the standard `net/http` client so I have configured the http client to use a sensible timeout of 10 seconds in `client.go`

To help ensure a valid request body is passed to the `create` endpoint I have included structs representing Account data in `models.go`. I use these structs to marshal values to and from JSON. Rather than returning raw JSON output `apiclient` unmarshalls the JSON responses into Go structs before returning them. I have chosen to expose these structs within the package so that other packages can reuse them for marshalling/unmarshalling.

For the optional parameters which can be passed to the `list` endpoint I chose to store these inside a struct using *int. I chose to use pointers to int as this type can differentiate between 0 and nil. 

I have used `gopkg.in/retry.v1` to implement the exponential back-off as recommended in the Account API documents. This will retry for a maximum of 60 seconds when responses with status codes 429, 500, 503 or 504 are received.

# Usage

Create a new instance of `apiclient` using the `New` function
Use this instance of `apiclient` to call the required endpoint function
Data is returned as a struct

# Testing

From the description of the task, I understand the job of `apiclient` is to issue certain calls to the Accounts API and return the responses. The tests for `apiclient` should therefore verify the calls it makes are correct. I have chosen to use a test server from  `net/http/httptest` to intercept the requests made by `apiclient` and return a variety of responses and status codes. This ensures that `apiclient` tests can run independely of the Accounts API that they interact with, and they will not fail if the Accounts API is unavailable. 

When unit tests trigger the exponential back off for retrying, a limit of 10ms is set for retrying. This triggers the retry a couple of times for a 500 response but still allows the tests to continue running without timing out.

For testing the Accounts API itself (the code pre-written by Form3) unit tests asserting on responses generated should exist in the AccountsAPI code base itself and be run by the AccountsAPI pipeline every time the AccountsAPI is built. 

# Docker

I built my Dockerfile using `docker build -t my-apiclient .`
I ran my Docker image using `docker run --rm my-apiclient:latest`
I have tagged my Docker image and uploaded it to DockerHub. It is available at `https://hub.docker.com/u/rosalita`