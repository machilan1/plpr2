package sqldb

// Config is the required properties to use the database.
type Config struct {
	User         string
	Password     string
	Host         string
	Name         string
	Schema       string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
	// TimeZone determines the timezone used in the session. Default value is "UTC".
	TimeZone string
	// CloudSQLInstanceConnName is the connection name for the CloudSQL instance.
	// For example, "project:region:instance".
	// When set, the CloudSQL connector will be used, and `Password`, `Host`, `DisableTLS`
	// will be ignored. The `User` is the IAM user account email address that has access to the CloudSQL instance.
	CloudSQLInstanceConnName string
}
