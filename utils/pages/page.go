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

func NewPageInfoForSource[T any](source []T) *PageInfo[T] {
	return &PageInfo[T]{
		source: source,
	}
}
func (m *PageInfo[T]) CopyDataList2Source() {
	if m != nil && m.source != nil {
		m.DataList = m.source[:len(m.source)]
		m.Total = uint64(len(m.source))
	}
}

// BuildMemoryPageDataListForSource 转换内存分页数据(Convert paginated data)
func (m *PageInfo[T]) BuildMemoryPageDataListForSource(pageNumArg, pageSizeArg uint64) {
	m.checkPagNumAndSize(&pageNumArg, &pageSizeArg)
	var offset, limit uint64
	m.PageNum, m.PageSize, offset, m.Total, m.Pages, limit = m.paginatedData(pageNumArg, pageSizeArg)
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
	limit = pageSize
	if total%pageSizeArg == 0 {
		pageMax = total / pageSizeArg
	} else {
		pageMax = total/pageSizeArg + 1
	}
	if defaultZero < pageMax && pageNum > pageMax {
		pageNum = pageMax
	}
	if offset = pageSize * (pageNum - 1); offset+pageSize > total {
		limit = total - offset
	}
	return
}
