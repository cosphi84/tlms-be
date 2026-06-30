# Office Module Implementation Plan

The goal is to implement a complete set of REST APIs for the Office module with full CRUD operations, pagination, specific data mapping for frontend Select components, and secure Role-Based Access Control via Casbin.

## Open Questions
- Should the `Delete` endpoint be a hard delete or a soft delete? (I will implement **soft delete** using GORM's `DeletedAt` feature as requested, but please let me know if this needs to change).
- For Casbin permissions, the plan maps `/api/offices` routes to roles. Is it acceptable to use the user's **Role** (e.g., `superadmin`) as the Casbin subject rather than their individual User ID? This is significantly faster and doesn't require adding user-role mapping rows in the database for every new user.

## Proposed Changes

### 1. Security & Casbin Integration
#### [MODIFY] `internal/middleware/authorize.go`
- Update the Casbin subject (`sub`) from `fmt.Sprintf("user: %d", claims.UserID)` to `string(claims.Role)`. This allows Casbin to resolve permissions instantly based on the role embedded in the JWT.
#### [MODIFY] `internal/seeder/casbin.seeder.go`
- Update the seeder to explicitly grant:
  - `superadmin` and `admin_hq`: `GET`, `POST`, `PUT`, `DELETE` on `/offices` and `/offices/*`.
  - All other roles: `GET` on `/offices` and `/offices/*`.

### 2. DTOs and Data Mapping
#### [MODIFY] `internal/dto/office.dto.go`
- Add `UpdateOfficeRequest` (Code, Name, Type).
- Ensure `OfficeOptionResponse` correctly maps `ID` and `Label` (formatted as `"name - code"`).

### 3. Database Repository
#### [MODIFY] `internal/repositories/office.repository.go`
- Add `FindAll(pagination *dto.PaginationRequest)` to list all active offices using your existing `Paginate` helper.
- Add `FindOptions()` to query just the IDs and concatenated labels for the Shadcn UI dynamic select.
- Add `Update(office *models.Office)` for edits.
- Add `Delete(id int64)` to trigger a GORM soft delete.

### 4. Business Logic (Service Layer)
#### [MODIFY] `internal/services/office.service.go`
- Add `GetOffices`, `GetOfficeOptions`, `UpdateOffice`, and `DeleteOffice` methods. 
- Ensure `CreateOffice` properly handles the parent-child relationship (this logic is mostly there already).

### 5. API Handlers and Routing
#### [NEW] `internal/handlers/office.handler.go`
- Create the HTTP handlers: `Create`, `FindAll`, `FindOptions`, `Update`, `Delete`.
#### [NEW] `internal/routes/office.route.go`
- Register `POST /offices`, `GET /offices`, `GET /offices/options`, `PUT /offices/:id`, and `DELETE /offices/:id`.
- Protect all routes with `middleware.Authenticate()` and `middleware.Authorize(authz)`.
#### [MODIFY] `internal/routes/route.go`
- Mount the new Office routes under the main API group.

### 6. Bootstrapping
#### [NEW] `internal/bootstraps/office.app.go`
- Create `InitOfficeModule` to wire the repository, service, and handler together (Dependency Injection).
#### [MODIFY] `internal/bootstraps/app.go`
- Add `OfficeHandler` to the `App` struct and initialize it.

## Verification Plan
1. Re-run `go run cmd/seed/main.go` to update Casbin policies.
2. Build the app using `go build ./...` to ensure no syntax errors.
