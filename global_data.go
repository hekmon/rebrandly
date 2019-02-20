package rebrandly

// OrderBy represent a given ordering
type OrderBy string

const (
	// OrderByCreatedAt represents the "createdAt" ordering for an OrderBy. Valid for Domains, Links.
	OrderByCreatedAt OrderBy = "createdAt"
	// OrderByUpdatedAt represents the "updatedAt" ordering for an OrderBy. Valid for Domains.
	OrderByUpdatedAt OrderBy = "updatedAt"
	// OrderByFullName represents the "fullName" ordering for an OrderBy. Valid for Domains.
	OrderByFullName OrderBy = "fullName"
	// OrderByTitle represents the "title" ordering for an OrderBy. Valid for Links.
	OrderByTitle OrderBy = "title"
	// OrderBySlashTag represents the "slashtag" ordering for an OrderBy. Valid for Links.
	OrderBySlashTag OrderBy = "slashtag"
)

// OrderDir represents the sorting direction
type OrderDir string

const (
	// OrderDirDesc represents the "desc" sorting direction for a OrderDir.
	OrderDirDesc OrderDir = "desc"
	// OrderDirAsc represents the "asc" sorting direction for a OrderDir.
	OrderDirAsc OrderDir = "asc"
)
