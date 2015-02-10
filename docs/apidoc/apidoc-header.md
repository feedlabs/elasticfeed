# Principles

### HTTP Methods

- **GET** - Retrieve the JSON representation of a resource.
- **POST** - Create a new resource using the provided JSON representation.
- **PUT** - Replace a resource with the provided JSON representation.
- **DELETE** - Remove a resource.

### JSON

All entities are represented in JSON. The following rules and conventions apply:

- When sending JSON to the server via **POST** or **PUT**, make sure to specify the correct content type request header: **Content-Type: application/json**
- Invalid fields will be rejected rather than ignored. If, for example, you attempt to create a new entity and misspell one of the fields, or if you attempt to update an existing entity and include a field that cannot be modified, the server will respond with a 400 status code and an error message stating which field was invalid.
- The fields in the JSON documents returned by the server are in no particular order, and it may change. Do not depend on the order of the fields.

### Pretty Printing

By default, extraneous whitespace is stripped from the JSON returned by the server. To ask for pretty-printed JSON, simply append the **pretty=true** query parameter to any request. Note that all the examples in this document show pretty-printed JavaScript for clarity, although the example URLs do not contain this additional query parameter.

### Response Codes

Responses utilize the standard HTTP response codes, including:

Code | Meaning               | Notes
---- | --------------------- | -------------------------------------------------------------------------------------------------------
200  | OK                    | The request was successful. This is typically the response to a successful **GET** request.
201  | Created               | A new resource was created. This is typically the response to a succcessful **POST** request.
400  | Bad Request           | Something was wrong with the client request.
401  | Unauthorized          | Authentication is required but was not present in the request. Typically this means that the digest authentication information was omitted from the request.
403  | Forbidden             | Access to the specified resource is not permitted. Usually means that the user associated with the given API Key is not allowed to access the requested resource.
404  | Not Found             | The requested resource does not exist.
405  | Method Not Allowed    | The HTTP method is not supported for the specified resource. Keep in mind that each resource may only support a subset of HTTP methods. For example, you are not allowed to DELETE the root resource.
409  | Conflict              | This is typically the response to a request to create or modify a property of an entity that is unique when an existing entity already exists with the same value for that property. For example, attempting to create a group with the same name as an existing group is not allowed.
429  | Too Many Requests     | You have exceeded the rate limit. See the section on Rate Limiting for more information.
5xx  | Various server errors | Something unexpected went wrong. Try again later and consider notifying MMS Support.
