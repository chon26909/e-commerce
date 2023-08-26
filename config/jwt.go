package config

type jwt struct {
	adminKey         string
	secertKey        string
	apiKey           string
	accessExpiresAt  int
	refreshExpiresAt int
}

type IJwtConfig interface {
	SecertKey() []byte
	AdminKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	SetJwtAccessExpiresAt(t int)
	RefreshExpiresAt() int
	SetJwtRefreshExpiresAt(t int)
}

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}

// SecertKey implements IJwtConfig.
func (j *jwt) SecertKey() []byte {
	return []byte(j.secertKey)
}

// AdminKey implements IJwtConfig.
func (j *jwt) AdminKey() []byte {
	return []byte(j.adminKey)
}

// ApiKey implements IJwtConfig.
func (j *jwt) ApiKey() []byte {
	return []byte(j.apiKey)
}

// AccessExpiresAt implements IJwtConfig.
func (j *jwt) AccessExpiresAt() int {
	return j.accessExpiresAt
}

// RefreshExpiresAt implements IJwtConfig.
func (j *jwt) RefreshExpiresAt() int {
	return j.refreshExpiresAt
}

// SetJwtAccessExpiresAt implements IJwtConfig.
func (j *jwt) SetJwtAccessExpiresAt(t int) {
	j.accessExpiresAt = t
}

// SetJwtRefreshExpiresAt implements IJwtConfig.
func (j *jwt) SetJwtRefreshExpiresAt(t int) {
	j.refreshExpiresAt = t
}
