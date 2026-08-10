

## Define the requirements

1. Functional Requirements 
    - User signup, signin
    - Listing create, update, list, delete
    - Product image uploads
    - Category create, update, list, delete
    - Only owners can edit / delete the listing

2. Non Functional Requirements
    - Security - secure password, secure upload
    - Privacy - strip the GPS
    - Reliability - worker crash - job reclaim, exponential backoff
    - Availability 
    - Performance - upload, download, caching
    - Scalability - worker horizontally scale
    - Portability - Object store - cloudflare R2, could be switch without changing much code just need few configuration changes
    - Maintainability - Code structure, versioned migrations
    - Observability - Structured Logging, Request correlations (users error with ID could be shared or pin pointed)


## High Level Design


Browser  -----(req)---->  Api ( golang ) --------> database


## Software Design Approaches
- Data First (Find the things (entities)) make a table for each ---> `Data Driven Design` Simple apps (CRUD), Inventories, blogs --> Mostly showing list and editing lists - Data-first
- Behavior First
(Find what app does - place order, cancel booking, invoice)
Build the logic first and then figure out how to store that data
`Domain Driven Design` --> Full of rules, complex logic - Behavior first


## Identify Entities

1. User
    - id
    - name
    - email
    - password

2. Listing
    - id
    - title
    - description
    - price
    - city
    - status
    - user_id
    - category_id

3. Image ( not sure initially )
    - id
    - listing_id
    - object_key

4. Category
    - id
    - name


`go run main.go` - 

`go build -o bin/main main.go`



