// Code generated by Ice-cream-maker DO NOT EDIT.
package v3
 
func (ClusterToken) ColumnNames() []string {
	return []string{
 		"id",
 		"uuid",
 		"name",
 		"summary",
 		"cluster_uuid",
 		"token",
 		"issued_at_time",
 		"expiration_time",
 		"created",
 		"updated",
 		"deleted",
	}
}
 
func (ClusterToken) ColumnNamesWithAlias() []string {
	return []string{
 		"id",
 		"uuid",
 		"name",
 		"summary",
 		"cluster_uuid",
 		"token",
 		"issued_at_time",
 		"expiration_time",
 		"created",
 		"updated",
 		"deleted",
	}
}
 
func (row ClusterToken) Values() []interface{} {
	return []interface{}{
		row.ID,
		row.Uuid,
		row.Name,
		row.Summary,
		row.ClusterUuid,
		row.Token,
		row.IssuedAtTime,
		row.ExpirationTime,
		row.Created,
		row.Updated,
		row.Deleted,
	}
}

type Scanner interface {
	Scan(dest ...interface{}) error
}
 
func (row *ClusterToken) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.ID,
		&row.Uuid,
		&row.Name,
		&row.Summary,
		&row.ClusterUuid,
		&row.Token,
		&row.IssuedAtTime,
		&row.ExpirationTime,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	)
}
 
func (row *ClusterToken) Ptrs() []interface{} {
	return []interface{}{
		&row.ID,
		&row.Uuid,
		&row.Name,
		&row.Summary,
		&row.ClusterUuid,
		&row.Token,
		&row.IssuedAtTime,
		&row.ExpirationTime,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	}
}
