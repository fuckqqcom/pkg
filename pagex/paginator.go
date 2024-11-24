package pagex

func Paginator(pageNum, pageSize int64, isLimit bool) (int64, int64, int64) {
	/*
		isLimit:true表示不限制分页参数
	*/
	if isLimit {
		return pageNum, pageSize, pageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 || pageSize >= 20 {
		pageSize = 10
	}
	// num limit,offset
	return pageNum, pageSize, (pageNum - 1) * pageSize
}

func OrderBy(orderBy map[string]any) map[string]any {
	if len(orderBy) == 0 {
		return map[string]any{"created_time": "desc"}
	}
	return orderBy
}
