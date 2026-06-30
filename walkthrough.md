# Auth Module Refactor Walkthrough

I have successfully refactored the auth module to support refresh tokens and make user context more easily available across the app. Here is a summary of what changed:

## Changes Made

### 1. Token Generation Updates
- **File:** `internal/auth/constants.go` and `internal/auth/jwt.go`
- **Change:** I added a `TokenType` string to `JWTClaims` so we can differentiate between access and refresh tokens. The token generator was then updated to produce a **Token Pair**:
  - `access_token`: Expires in 15 minutes.
  - `refresh_token`: Expires in 7 days.

### 2. DTO and API Endpoint
- **File:** `internal/dto/authenticate.dto.go`, `internal/handlers/authenticate.handler.go`, and `internal/routes/authenticate.route.go`
- **Change:** 
  - Login now returns `access_token` and `refresh_token` inside the JSON response.
  - I created a new POST `/auth/refresh` endpoint. This route expects a JSON body with a `refresh_token` and uses it to issue a brand new token pair.

### 3. Service Layer Enhancements
- **File:** `internal/services/authenticate.service.go`
- **Change:** I implemented the `RefreshToken` business logic. It securely parses and validates the provided refresh token's signature, ensures its `TokenType` is indeed `"refresh"`, looks up the user in the database to verify they are still active, and finally generates a new pair.

### 4. Middleware and Easy Data Access
- **File:** `internal/middleware/authenticate.go`
- **Change:** 
  - The middleware now strictly rejects any tokens where `TokenType` is not `"access"`, preventing users from using a refresh token to access protected APIs.
  - **New Feature:** The `userID`, `officeID`, and `role` are now directly injected into the Gin context. This solves your requirement!

> [!TIP]
> **How to use `createdBy` easily now:**
> Inside any handler protected by the `Authenticate` middleware, you can simply do:
> ```go
> userID := c.GetInt64("user_id") // Very easy to use for 'createdBy'
> officeID := c.GetInt64("office_id")
> role := c.GetString("role")
> ```
> This avoids having to call `auth.GetClaims()` explicitly and extract the ID from the struct yourself.

## Validation Results
- Verified that all changes compile successfully (`go build ./...` returned 0).
