package rebrandly

// OrderBy represent a given ordering
type OrderBy string

const (
	// OrderByCreatedAt represents the "createdAt" ordering for a OrderBy
	OrderByCreatedAt OrderBy = "createdAt"
	// OrderByUpdatedAt represents the "updatedAt" ordering for a OrderBy
	OrderByUpdatedAt OrderBy = "updatedAt"
	// OrderByFullName represents the "fullName" ordering for a OrderBy
	OrderByFullName OrderBy = "fullName"
)

// OrderDir represents the sorting direction
type OrderDir string

const (
	// OrderDirDesc represents the "desc" sorting direction for a OrderDir
	OrderDirDesc OrderDir = "desc"
	// OrderDirAsc represents the "asc" sorting direction for a OrderDir
	OrderDirAsc OrderDir = "asc"
)
