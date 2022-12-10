package utils

import "database/sql"

func NullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}

func NullFloat64(s float64) (ns sql.NullFloat64) {
	if s != 0 {
		ns.Float64 = s
		ns.Valid = true
	}
	return ns
}

func FormatNullTime(nt sql.NullTime, format string) string {
	if nt.Valid {
		return nt.Time.Format(format)
	}
	return ""
}
