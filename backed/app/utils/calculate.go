package utils

// CalculatePaginationIndexes 计算分页的开始和结束索引
func CalculatePaginationIndexes(total, page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	page = page - 1
	start := total - (page * pageSize) - 1
	end := start - pageSize + 1
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	return start, end
}
