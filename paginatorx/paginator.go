package paginatorx

func Paginator(pageNum, pageSize int64) (int64, int64) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 || pageSize >= 20 {
		pageSize = 10
	}
	return (pageNum - 1) * pageSize, pageSize
}

func OrderBy(orderBy string) string {
	if orderBy == "" {
		return " created_time desc "
	}
	return orderBy
}
