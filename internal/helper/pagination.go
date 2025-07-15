package helper

type Pagination struct {
	Total int `json:"total"`
	Size  int `json:"size"`
	Page  int `json:"page"`
}

// return limit, offset
func GetLimitOffsett(page, size int) (int, int) {
	return size, (page - 1) * size
}
