package types

type Distributor struct {
	Code        string                  `json:"code"`
	Permissions *DistributorPermissions `json:"permissions"`
	ParentCode  string                  `json:"parent_code"`
}

type DistributorPermissions struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

type IsServiceableRequest struct {
	Code   string `json:"code"`
	Region string `json:"region"`
}

type IsServiceableResponse struct {
	Code          string `json:"code"`
	Region        string `json:"region"`
	IsServiceable string `json:"is_serviceable"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
