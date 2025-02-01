package tenant

type CreateTenantRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTenantResponse struct {
	CreateTenantRequest
	ID string `json:"id"`
}
