# YAML structure for commands

## GET

```yaml
api_requests:
  - method: GET
    url: "/api/v1/resources"
    description: "Fetch all resources"
    headers:
      Content-Type: "application/json"
    params:
      page: 1
      per_page: 20
    response:
      status: 200
      body:
        - id: 1
          name: "Resource 1"
          description: "Description of Resource 1"
        - id: 2
          name: "Resource 2"
          description: "Description of Resource 2"
          
  - method: POST
    url: "/api/v1/resources"
    description: "Create a new resource"
    headers:
      Content-Type: "application/json"
    body:
      name: "New Resource"
      description: "This is a new resource"
    response:
      status: 201
      body:
        id: 3
        name: "New Resource"
        description: "This is a new resource"
        
  - method: PUT
    url: "/api/v1/resources/{id}"
    description: "Update a resource"
    headers:
      Content-Type: "application/json"
    body:
      name: "Updated Resource"
      description: "This is an updated resource"
    response:
      status: 200
      body:
        id: "{id}"
        name: "Updated Resource"
        description: "This is an updated resource"

  - method: PATCH
    url: "/api/v1/resources/{id}"
    description: "Partially update a resource"
    headers:
      Content-Type: "application/json"
    body:
      description: "Updated partial description"
    response:
      status: 200
      body:
        id: "{id}"
        name: "Resource 1"
        description: "Updated partial description"

  - method: DELETE
    url: "/api/v1/resources/{id}"
    description: "Delete a resource"
    headers:
      Content-Type: "application/json"
    response:
      status: 204
      body: null
```

## Notes

- use whitespaces over tabs for easy parsing -> 2 whitespace indentation
