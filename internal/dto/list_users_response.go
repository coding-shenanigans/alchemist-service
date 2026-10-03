package dto

type ListUsersResponse struct {
	Users         []*UserResult `json:"users"`
	NextPageToken string        `json:"nextPageToken,omitempty"`
}
