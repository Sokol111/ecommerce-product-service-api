# DefaultApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createProduct**](#createproduct) | **POST** /product/create | Create a new product|
|[**getAll**](#getall) | **GET** /product/list | Get a list of all products|
|[**getProductById**](#getproductbyid) | **GET** /product/get/{id} | Get a product by ID|
|[**updateProduct**](#updateproduct) | **PUT** /product/update | Update an existing product|

# **createProduct**
> ProductResponse createProduct(createProductRequest)


### Example

```typescript
import {
    DefaultApi,
    Configuration,
    CreateProductRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let createProductRequest: CreateProductRequest; //

const { status, data } = await apiInstance.createProduct(
    createProductRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createProductRequest** | **CreateProductRequest**|  | |


### Return type

**ProductResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Product created successfully |  -  |
|**400** | Validation error |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAll**
> Array<ProductResponse> getAll()


### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.getAll();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<ProductResponse>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of products retrieved successfully |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getProductById**
> ProductResponse getProductById()


### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let id: string; // (default to undefined)

const { status, data } = await apiInstance.getProductById(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] |  | defaults to undefined|


### Return type

**ProductResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Product retrieved successfully |  -  |
|**404** | Product not found |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateProduct**
> ProductResponse updateProduct(updateProductRequest)


### Example

```typescript
import {
    DefaultApi,
    Configuration,
    UpdateProductRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let updateProductRequest: UpdateProductRequest; //

const { status, data } = await apiInstance.updateProduct(
    updateProductRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateProductRequest** | **UpdateProductRequest**|  | |


### Return type

**ProductResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Product updated successfully |  -  |
|**412** | Versions mismatch |  -  |
|**400** | Validation error |  -  |
|**500** | Standard internal error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

