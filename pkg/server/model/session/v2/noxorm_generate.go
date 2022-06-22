package v2

func (Session_essential) ColumnNames() []string {
	return []string{"name", "summary", "cluster_uuid", "token", "issued_at_time", "expiration_time"}
}

func (Session) ColumnNames() []string {
	return []string{"id", "uuid", "name", "summary", "cluster_uuid", "token", "issued_at_time", "expiration_time", "created", "updated", "deleted"}
}

func (row Session_essential) Values() []interface{} {
	return []interface{}{row.Name, row.Summary, row.ClusterUuid, row.Token, row.IssuedAtTime, row.ExpirationTime}
}

func (row Session) Values() []interface{} {
	return []interface{}{row.ID, row.Uuid, row.Session_essential.Name, row.Session_essential.Summary, row.Session_essential.ClusterUuid, row.Session_essential.Token, row.Session_essential.IssuedAtTime, row.Session_essential.ExpirationTime, row.Created, row.Updated, row.Deleted}
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func (row *Session_essential) Scan(scanner Scanner) error {
	return scanner.Scan(&row.Name, &row.Summary, &row.ClusterUuid, &row.Token, &row.IssuedAtTime, &row.ExpirationTime)
}

func (row *Session) Scan(scanner Scanner) error {
	return scanner.Scan(&row.ID, &row.Uuid, &row.Session_essential.Name, &row.Session_essential.Summary, &row.Session_essential.ClusterUuid, &row.Session_essential.Token, &row.Session_essential.IssuedAtTime, &row.Session_essential.ExpirationTime, &row.Created, &row.Updated, &row.Deleted)
}
