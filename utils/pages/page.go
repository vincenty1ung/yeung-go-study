package pages

const (
	defaultZero     = 0
	defaultPageNum  = 1
	defaultPageSize = 10
)

type PageInfo[T any] struct {
	source   []T
	DataList []T    `json:"data_list"`
	PageNum  uint64 `json:"page_num"`
	PageSize uint64 `json:"page_size"`
	Pages    uint64 `json:"pages"`
	Total    uint64 `json:"total"`
}

func NewPageInfoByDataList[T any](source []T) *PageInfo[T] {
	return &PageInfo[T]{
		source: source,
	}
}

// ConvertMemoryPaginatedDataByPagNumAndSize 转换内存分页数据(Convert paginated data)
func (m *PageInfo[T]) ConvertMemoryPaginatedDataByPagNumAndSize(pageNumArg, pageSizeArg uint64) {
	m.checkPagNumAndSize(&pageNumArg, &pageSizeArg)
	pageNum, pageSize, offset, total, pageMax, limit := m.paginatedData(pageNumArg, pageSizeArg)
	m.Pages = pageMax
	m.Total = total
	m.PageNum = pageNum
	m.PageSize = pageSize
	m.DataList = m.source[offset : offset+limit]
}

func (m *PageInfo[T]) checkPagNumAndSize(pageNumArg, pageSizeArg *uint64) {
	if *pageNumArg == defaultZero {
		*pageNumArg = defaultPageNum
	}
	if *pageSizeArg == defaultZero {
		*pageSizeArg = defaultPageSize
	}
}
func (m *PageInfo[T]) paginatedData(pageNumArg, pageSizeArg uint64) (pageNum, pageSize, offset, total, pageMax, limit uint64) {
	pageNum, pageSize = pageNumArg, pageSizeArg
	total = uint64(len(m.source))
	if total%pageSizeArg == 0 {
		pageMax = total / pageSizeArg
	} else {
		pageMax = total/pageSizeArg + 1
	}
	// 当前页合法性检查
	if pageNumArg > pageMax {
		pageNum, pageNumArg = pageMax, pageMax
	}
	// 越界超精度判断
	if offset = pageSizeArg * (pageNumArg - 1); offset >= total {
		offset = 0
		pageNum = 1
	}
	// 越界检测
	limit = pageSize
	if offset+pageSize > total {
		limit = total - offset
	}
	return
}
