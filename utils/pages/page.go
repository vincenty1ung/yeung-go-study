package pages

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

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

func NewPageInfo[T any]() *PageInfo[T] {
	return &PageInfo[T]{}
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

func (m *PageInfo[T]) FindByPageArgAndQuery(
	ctx context.Context, pageNum, pageSize uint64, countBuilder, queryBuilder squirrel.SelectBuilder,
	conn sqlc.CachedConn,
) error {
	m.checkPagNumAndSize(&pageSize, &pageSize)
	var (
		resp  []T
		count uint64
	)
	countQuery, values, err := countBuilder.ToSql()
	if err != nil {
		return err
	}
	err = conn.QueryRowNoCacheCtx(ctx, &count, countQuery, values...)
	if err != nil {
		return err
	}
	offset := pageSize * (pageNum - 1)
	queryBuilder = queryBuilder.Limit(pageSize).Offset(offset)
	query, values, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}
	err = conn.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		return err
	}

	m.Pages = m.pagesNum(count, pageSize)
	m.Total = count
	m.PageNum = pageNum
	m.PageSize = pageSize
	m.DataList = resp
	return nil
}

func (m *PageInfo[T]) pagesNum(count uint64, pageSize uint64) uint64 {
	// 计算最大页
	var pages uint64
	if count%pageSize == 0 {
		pages = count / pageSize
	} else {
		pages = count/pageSize + 1
	}
	return pages
}
