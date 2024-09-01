package common

// PageData 是一个泛型结构体，T 代表实际的数据类型
type PageData[T any] struct {
	Page       int // 当前页码
	PageSize   int // 每页显示的记录数
	TotalPages int // 总页数
	TotalCount int // 总记录数
	Data       []T // 实际数据列表
}

// NewPageData 创建一个新的 PageData 实例
func NewPageData[T any](page, pageSize, totalCount int, data []T) *PageData[T] {
	println(page, pageSize)
	totalPages := 0
	if totalCount > 0 {
		totalPages = (totalCount + pageSize - 1) / pageSize // 计算总页数
	}

	return &PageData[T]{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		TotalCount: totalCount,
		Data:       data,
	}
}
