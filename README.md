

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
- Data First
- Behavior First

