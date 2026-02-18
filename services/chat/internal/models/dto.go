package models

type CursorInfo struct {
	UUID     string
	IsCursor bool
}

type Pagination struct {
	PageSize int32
	Cursor   string
}

func ToCursorInfo(cursor string) CursorInfo {
	if cursor == "" {
		return CursorInfo{IsCursor: false}
	}
	return CursorInfo{UUID: cursor, IsCursor: true}
}
