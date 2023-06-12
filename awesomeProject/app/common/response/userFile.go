package response

type UserFileResponse struct {
	List  interface{} `json:"list"`
	Count int         `json:"count"`
}
