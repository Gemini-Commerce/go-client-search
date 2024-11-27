# \SearchAPI

All URIs are relative to *https://search.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchConfiguration**](SearchAPI.md#SearchConfiguration) | **Post** /search.Search/Configuration | Indexes configuration
[**SearchDelete**](SearchAPI.md#SearchDelete) | **Post** /search.Search/Delete | Delete Indexes
[**SearchInsertOrReplace**](SearchAPI.md#SearchInsertOrReplace) | **Post** /search.Search/InsertOrReplace | Insert or replace documents to different elasticsearch indexes
[**SearchMSearch**](SearchAPI.md#SearchMSearch) | **Post** /search.Search/MSearch | Send queries to different elasticsearch indexes
[**SearchUpdate**](SearchAPI.md#SearchUpdate) | **Post** /search.Search/Update | Update documents to different elasticsearch indexes



## SearchConfiguration

> map[string]interface{} SearchConfiguration(ctx).Body(body).Execute()

Indexes configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-search"
)

func main() {
	body := *openapiclient.NewSearchConfigRequest() // SearchConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchConfiguration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchConfigRequest**](SearchConfigRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [geminiAuthorization](../README.md#geminiAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDelete

> SearchDeleteResponse SearchDelete(ctx).Body(body).Execute()

Delete Indexes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-search"
)

func main() {
	body := *openapiclient.NewSearchDeleteRequest() // SearchDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDelete`: SearchDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchDeleteRequest**](SearchDeleteRequest.md) |  | 

### Return type

[**SearchDeleteResponse**](SearchDeleteResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [geminiAuthorization](../README.md#geminiAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchInsertOrReplace

> SearchWriteResponse SearchInsertOrReplace(ctx).Body(body).Execute()

Insert or replace documents to different elasticsearch indexes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-search"
)

func main() {
	body := *openapiclient.NewSearchWriteRequest("Index_example", []string{"Documents_example"}) // SearchWriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchInsertOrReplace(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchInsertOrReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchInsertOrReplace`: SearchWriteResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchInsertOrReplace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchInsertOrReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchWriteRequest**](SearchWriteRequest.md) |  | 

### Return type

[**SearchWriteResponse**](SearchWriteResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [geminiAuthorization](../README.md#geminiAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMSearch

> SearchQueryResponse SearchMSearch(ctx).Body(body).Execute()

Send queries to different elasticsearch indexes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-search"
)

func main() {
	body := *openapiclient.NewSearchQueryRequest([]openapiclient.SearchPayload{*openapiclient.NewSearchPayload("Index_example", *openapiclient.NewSearchParams())}) // SearchQueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchMSearch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchMSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMSearch`: SearchQueryResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchMSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchQueryRequest**](SearchQueryRequest.md) |  | 

### Return type

[**SearchQueryResponse**](SearchQueryResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [geminiAuthorization](../README.md#geminiAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdate

> SearchWriteResponse SearchUpdate(ctx).Body(body).Execute()

Update documents to different elasticsearch indexes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-search"
)

func main() {
	body := *openapiclient.NewSearchWriteRequest("Index_example", []string{"Documents_example"}) // SearchWriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchUpdate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdate`: SearchWriteResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchWriteRequest**](SearchWriteRequest.md) |  | 

### Return type

[**SearchWriteResponse**](SearchWriteResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [geminiAuthorization](../README.md#geminiAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

