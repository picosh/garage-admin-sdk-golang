# \BucketAPI

All URIs are relative to *http://localhost:3903/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowBucketKey**](BucketAPI.md#AllowBucketKey) | **Post** /bucket/allow | Allow key
[**CreateBucket**](BucketAPI.md#CreateBucket) | **Post** /bucket | Create a bucket
[**DeleteBucket**](BucketAPI.md#DeleteBucket) | **Delete** /bucket | Delete a bucket
[**DeleteBucketGlobalAlias**](BucketAPI.md#DeleteBucketGlobalAlias) | **Delete** /bucket/alias/global | Delete a global alias
[**DeleteBucketLocalAlias**](BucketAPI.md#DeleteBucketLocalAlias) | **Delete** /bucket/alias/local | Delete a local alias
[**DenyBucketKey**](BucketAPI.md#DenyBucketKey) | **Post** /bucket/deny | Deny key
[**GetBucketInfo**](BucketAPI.md#GetBucketInfo) | **Get** /bucket | Get a bucket
[**ListBuckets**](BucketAPI.md#ListBuckets) | **Get** /bucket?list | List all buckets
[**PutBucketGlobalAlias**](BucketAPI.md#PutBucketGlobalAlias) | **Put** /bucket/alias/global | Add a global alias
[**PutBucketLocalAlias**](BucketAPI.md#PutBucketLocalAlias) | **Put** /bucket/alias/local | Add a local alias
[**UpdateBucket**](BucketAPI.md#UpdateBucket) | **Put** /bucket | Update a bucket



## AllowBucketKey

> BucketInfo AllowBucketKey(ctx).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()

Allow key



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
	allowBucketKeyRequest := *openapiclient.NewAllowBucketKeyRequest("e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b", "GK31c2f218a2e44f485b94239e", *openapiclient.NewAllowBucketKeyRequestPermissions(true, true, true)) // AllowBucketKeyRequest | Aliases to put on the new bucket 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.AllowBucketKey(context.Background()).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.AllowBucketKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllowBucketKey`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.AllowBucketKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowBucketKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowBucketKeyRequest** | [**AllowBucketKeyRequest**](AllowBucketKeyRequest.md) | Aliases to put on the new bucket  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBucket

> BucketInfo CreateBucket(ctx).CreateBucketRequest(createBucketRequest).Execute()

Create a bucket



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
	createBucketRequest := *openapiclient.NewCreateBucketRequest() // CreateBucketRequest | Aliases to put on the new bucket 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.CreateBucket(context.Background()).CreateBucketRequest(createBucketRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.CreateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBucket`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.CreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBucketRequest** | [**CreateBucketRequest**](CreateBucketRequest.md) | Aliases to put on the new bucket  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> DeleteBucket(ctx).Id(id).Execute()

Delete a bucket



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
	id := "b4018dc61b27ccb5c64ec1b24f53454bbbd180697c758c4d47a22a8921864a87" // string | The exact bucket identifier, a 32 bytes hexadecimal string

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BucketAPI.DeleteBucket(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.DeleteBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The exact bucket identifier, a 32 bytes hexadecimal string | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucketGlobalAlias

> BucketInfo DeleteBucketGlobalAlias(ctx).Id(id).Alias(alias).Execute()

Delete a global alias



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
	id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
	alias := "my_documents" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.DeleteBucketGlobalAlias(context.Background()).Id(id).Alias(alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.DeleteBucketGlobalAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBucketGlobalAlias`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.DeleteBucketGlobalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketGlobalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucketLocalAlias

> BucketInfo DeleteBucketLocalAlias(ctx).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()

Delete a local alias



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
	id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
	accessKeyId := "GK31c2f218a2e44f485b94239e" // string | 
	alias := "my_documents" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.DeleteBucketLocalAlias(context.Background()).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.DeleteBucketLocalAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBucketLocalAlias`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.DeleteBucketLocalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketLocalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **accessKeyId** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DenyBucketKey

> BucketInfo DenyBucketKey(ctx).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()

Deny key



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
	allowBucketKeyRequest := *openapiclient.NewAllowBucketKeyRequest("e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b", "GK31c2f218a2e44f485b94239e", *openapiclient.NewAllowBucketKeyRequestPermissions(true, true, true)) // AllowBucketKeyRequest | Aliases to put on the new bucket 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.DenyBucketKey(context.Background()).AllowBucketKeyRequest(allowBucketKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.DenyBucketKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DenyBucketKey`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.DenyBucketKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDenyBucketKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowBucketKeyRequest** | [**AllowBucketKeyRequest**](AllowBucketKeyRequest.md) | Aliases to put on the new bucket  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucketInfo

> BucketInfo GetBucketInfo(ctx).Id(id).GlobalAlias(globalAlias).Execute()

Get a bucket



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
	id := "b4018dc61b27ccb5c64ec1b24f53454bbbd180697c758c4d47a22a8921864a87" // string | The exact bucket identifier, a 32 bytes hexadecimal string.  Incompatible with `alias`.  (optional)
	globalAlias := "my_documents" // string | The exact global alias of one of the existing buckets.  Incompatible with `id`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.GetBucketInfo(context.Background()).Id(id).GlobalAlias(globalAlias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.GetBucketInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBucketInfo`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.GetBucketInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The exact bucket identifier, a 32 bytes hexadecimal string.  Incompatible with &#x60;alias&#x60;.  | 
 **globalAlias** | **string** | The exact global alias of one of the existing buckets.  Incompatible with &#x60;id&#x60;.  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> []ListBuckets200ResponseInner ListBuckets(ctx).Execute()

List all buckets



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
	resp, r, err := apiClient.BucketAPI.ListBuckets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.ListBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuckets`: []ListBuckets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.ListBuckets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


### Return type

[**[]ListBuckets200ResponseInner**](ListBuckets200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketGlobalAlias

> BucketInfo PutBucketGlobalAlias(ctx).Id(id).Alias(alias).Execute()

Add a global alias



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
	id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
	alias := "my_documents" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.PutBucketGlobalAlias(context.Background()).Id(id).Alias(alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.PutBucketGlobalAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutBucketGlobalAlias`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.PutBucketGlobalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketGlobalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketLocalAlias

> BucketInfo PutBucketLocalAlias(ctx).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()

Add a local alias



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
	id := "e6a14cd6a27f48684579ec6b381c078ab11697e6bc8513b72b2f5307e25fff9b" // string | 
	accessKeyId := "GK31c2f218a2e44f485b94239e" // string | 
	alias := "my_documents" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.PutBucketLocalAlias(context.Background()).Id(id).AccessKeyId(accessKeyId).Alias(alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.PutBucketLocalAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutBucketLocalAlias`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.PutBucketLocalAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketLocalAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **accessKeyId** | **string** |  | 
 **alias** | **string** |  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> BucketInfo UpdateBucket(ctx).Id(id).UpdateBucketRequest(updateBucketRequest).Execute()

Update a bucket



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
	id := "b4018dc61b27ccb5c64ec1b24f53454bbbd180697c758c4d47a22a8921864a87" // string | The exact bucket identifier, a 32 bytes hexadecimal string
	updateBucketRequest := *openapiclient.NewUpdateBucketRequest() // UpdateBucketRequest | Requested changes on the bucket. Both root fields are optionals. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BucketAPI.UpdateBucket(context.Background()).Id(id).UpdateBucketRequest(updateBucketRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.UpdateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBucket`: BucketInfo
	fmt.Fprintf(os.Stdout, "Response from `BucketAPI.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The exact bucket identifier, a 32 bytes hexadecimal string | 
 **updateBucketRequest** | [**UpdateBucketRequest**](UpdateBucketRequest.md) | Requested changes on the bucket. Both root fields are optionals.  | 

### Return type

[**BucketInfo**](BucketInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

