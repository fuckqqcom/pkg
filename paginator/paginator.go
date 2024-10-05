package paginator

func Paginator(pageNum, pageSize int64) (int64, int64, int64) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 || pageSize >= 20 {
		pageSize = 10
	}
	// num limit,offset
	return pageNum, pageSize, (pageNum - 1) * pageSize
}

func OrderBy(orderBy string) string {
	if orderBy == "" {
		return " created_time desc "
	}
	return orderBy
}
