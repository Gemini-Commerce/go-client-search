# Go API client for search

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import search "github.com/Gemini-Commerce/go-client-search"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `search.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), search.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `search.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), search.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `search.ContextOperationServerIndices` and `search.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), search.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), search.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://search.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SearchAPI* | [**SearchConfiguration**](docs/SearchAPI.md#searchconfiguration) | **Post** /search.Search/Configuration | Indexes configuration
*SearchAPI* | [**SearchDelete**](docs/SearchAPI.md#searchdelete) | **Post** /search.Search/Delete | Delete Indexes
*SearchAPI* | [**SearchInsertOrReplace**](docs/SearchAPI.md#searchinsertorreplace) | **Post** /search.Search/InsertOrReplace | Insert or replace documents to different elasticsearch indexes
*SearchAPI* | [**SearchMSearch**](docs/SearchAPI.md#searchmsearch) | **Post** /search.Search/MSearch | Send queries to different elasticsearch indexes
*SearchAPI* | [**SearchUpdate**](docs/SearchAPI.md#searchupdate) | **Post** /search.Search/Update | Update documents to different elasticsearch indexes


## Documentation For Models

 - [ParamSortOrder](docs/ParamSortOrder.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [SearchAggrMap](docs/SearchAggrMap.md)
 - [SearchConfigRequest](docs/SearchConfigRequest.md)
 - [SearchConfigSchema](docs/SearchConfigSchema.md)
 - [SearchConfigSchemaAttribute](docs/SearchConfigSchemaAttribute.md)
 - [SearchConfigSchemaAttributeType](docs/SearchConfigSchemaAttributeType.md)
 - [SearchConfigSchemaFacet](docs/SearchConfigSchemaFacet.md)
 - [SearchConfigSchemaFilter](docs/SearchConfigSchemaFilter.md)
 - [SearchConfigSchemaSearchable](docs/SearchConfigSchemaSearchable.md)
 - [SearchConfigSchemaSortable](docs/SearchConfigSchemaSortable.md)
 - [SearchDeleteConstraints](docs/SearchDeleteConstraints.md)
 - [SearchDeleteError](docs/SearchDeleteError.md)
 - [SearchDeleteRequest](docs/SearchDeleteRequest.md)
 - [SearchDeleteResponse](docs/SearchDeleteResponse.md)
 - [SearchParamAttribute](docs/SearchParamAttribute.md)
 - [SearchParamFacet](docs/SearchParamFacet.md)
 - [SearchParamFacetType](docs/SearchParamFacetType.md)
 - [SearchParamFilter](docs/SearchParamFilter.md)
 - [SearchParamFilterType](docs/SearchParamFilterType.md)
 - [SearchParamSearchable](docs/SearchParamSearchable.md)
 - [SearchParamSort](docs/SearchParamSort.md)
 - [SearchParams](docs/SearchParams.md)
 - [SearchPayload](docs/SearchPayload.md)
 - [SearchQueryError](docs/SearchQueryError.md)
 - [SearchQueryRequest](docs/SearchQueryRequest.md)
 - [SearchQueryRequestType](docs/SearchQueryRequestType.md)
 - [SearchQueryResponse](docs/SearchQueryResponse.md)
 - [SearchResult](docs/SearchResult.md)
 - [SearchWriteError](docs/SearchWriteError.md)
 - [SearchWriteRequest](docs/SearchWriteRequest.md)
 - [SearchWriteResponse](docs/SearchWriteResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### geminiAuthorization

- **Type**: API key
- **API key parameter name**: X-Gem-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: geminiAuthorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		search.ContextAPIKeys,
		map[string]search.APIKey{
			"geminiAuthorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### standardAuthorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: standardAuthorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		search.ContextAPIKeys,
		map[string]search.APIKey{
			"standardAuthorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

