# \NodesAPI

All URIs are relative to *http://localhost:3903/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNode**](NodesAPI.md#AddNode) | **Post** /connect | Connect a new node
[**GetHealth**](NodesAPI.md#GetHealth) | **Get** /health | Cluster health report
[**GetNodes**](NodesAPI.md#GetNodes) | **Get** /status | Describe cluster



## AddNode

> []AddNode200ResponseInner AddNode(ctx).RequestBody(requestBody).Execute()

Connect a new node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "git.deuxfleurs.fr/garage-sdk/garage-admin-sdk-golang"
)

func main() {
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.AddNode(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.AddNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNode`: []AddNode200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.AddNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**[]AddNode200ResponseInner**](AddNode200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealth

> GetHealth200Response GetHealth(ctx).Execute()

Cluster health report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "git.deuxfleurs.fr/garage-sdk/garage-admin-sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.GetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealth`: GetHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


### Return type

[**GetHealth200Response**](GetHealth200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodes

> GetNodes200Response GetNodes(ctx).Execute()

Describe cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "git.deuxfleurs.fr/garage-sdk/garage-admin-sdk-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.GetNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodes`: GetNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesRequest struct via the builder pattern


### Return type

[**GetNodes200Response**](GetNodes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

