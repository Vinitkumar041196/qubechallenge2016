package types

type Distributor struct {
	Code        string                  `json:"code" example:"DISTRIBUTOR1"`
	Permissions *DistributorPermissions `json:"permissions"`
	ParentCode  string                  `json:"parent_code" example:"DISTRIBUTOR2"`
} //@name Distributor

type DistributorPermissions struct {
	Include []string `json:"include" example:"US,KA-IN,CENAI-TN-IN"`
	Exclude []string `json:"exclude" example:"YELUR-KA-IN"`
} //@name DistributorPermissions

type IsServiceableRequest struct {
	Code   string `json:"code" example:"DISTRIBUTOR1"`
	Region string `json:"region" example:"YELUR-KA-IN"`
} //@name IsServiceableRequest

type IsServiceableResponse struct {
	Code          string `json:"code" example:"DISTRIBUTOR1"`
	Region        string `json:"region" example:"YELUR-KA-IN"`
	IsServiceable string `json:"is_serviceable" enums:"YES,NO"`
} //@name IsServiceableResponse

type ErrorResponse struct {
	Error string `json:"error" example:"invalid parent code: DISTRIBUTOR2"`
} //@name ErrorResponse

type SuccessResponse struct {
	Code    string `json:"code" example:"DISTRIBUTOR1"`
	Message string `json:"message" example:"success"`
} //@name SuccessResponse
