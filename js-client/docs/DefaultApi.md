# ProductClient.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createProduct**](DefaultApi.md#createProduct) | **POST** /product/create | Create a new product
[**getAll**](DefaultApi.md#getAll) | **GET** /product/list | Get a list of all products
[**getProductById**](DefaultApi.md#getProductById) | **GET** /product/get/{id} | Get a product by ID
[**updateProduct**](DefaultApi.md#updateProduct) | **PUT** /product/update | Update an existing product



## createProduct

> ProductResponse createProduct(createProductRequest)

Create a new product

### Example

```javascript
import ProductClient from 'product-client';

let apiInstance = new ProductClient.DefaultApi();
let createProductRequest = new ProductClient.CreateProductRequest(); // CreateProductRequest | 
apiInstance.createProduct(createProductRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProductRequest** | [**CreateProductRequest**](CreateProductRequest.md)|  | 

### Return type

[**ProductResponse**](ProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getAll

> [ProductResponse] getAll()

Get a list of all products

### Example

```javascript
import ProductClient from 'product-client';

let apiInstance = new ProductClient.DefaultApi();
apiInstance.getAll().then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters

This endpoint does not need any parameter.

### Return type

[**[ProductResponse]**](ProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getProductById

> ProductResponse getProductById(id)

Get a product by ID

### Example

```javascript
import ProductClient from 'product-client';

let apiInstance = new ProductClient.DefaultApi();
let id = "id_example"; // String | 
apiInstance.getProductById(id).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**ProductResponse**](ProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateProduct

> ProductResponse updateProduct(updateProductRequest)

Update an existing product

### Example

```javascript
import ProductClient from 'product-client';

let apiInstance = new ProductClient.DefaultApi();
let updateProductRequest = new ProductClient.UpdateProductRequest(); // UpdateProductRequest | 
apiInstance.updateProduct(updateProductRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProductRequest** | [**UpdateProductRequest**](UpdateProductRequest.md)|  | 

### Return type

[**ProductResponse**](ProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

