package entity

type BukalapakTokenInfo struct {
  Token             string                      `json:"token"`
  Scopes            string                      `json:"scopes"`
  RefreshToken      string                      `json:"refresh_token"`
  ResourceOwnerID   int                         `json:"resource_owner_id"`
  ResourceOwnerRole string                      `json:"resource_owner_role"`
  ResourceOwner     BukalapakTokenResourceOwner `json:"resource_owner"`
  ApplicationID     int                         `json:"application_id"`
  ApplicationName   string                      `json:"application_json"`
  ApplicationScopes string                      `json:"application_scopes"`
  ApplicationUID    string                      `json:"application_uid"`
}

type BukalapakTokenResourceOwner struct {
  ID       int    `json:"id"`
  Username string `json:"username"`
  Name     string `json:"name"`
  Email    string `json:"email"`
  Gender   string `json:"gender"`
  Phone    string `json:"phone"`
  Role     string `json:"role"`
}
