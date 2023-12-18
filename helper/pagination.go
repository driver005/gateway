package helper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

// Parse parses limit and offset from *http.Request with given limits and defaults.
func Parse(r *http.Request, defaultLimit, defaultOffset, maxLimit int) (int, int) {
	var offset, limit int

	if offsetParam := r.URL.Query().Get("offset"); offsetParam == "" {
		offset = defaultOffset
	} else {
		if offset64, err := strconv.ParseInt(offsetParam, 10, 64); err != nil {
			offset = defaultOffset
		} else {
			offset = int(offset64)
		}
	}

	if limitParam := r.URL.Query().Get("limit"); limitParam == "" {
		limit = defaultLimit
	} else {
		if limit64, err := strconv.ParseInt(limitParam, 10, 64); err != nil {
			limit = defaultLimit
		} else {
			limit = int(limit64)
		}
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	if limit < 0 {
		limit = 0
	}

	if offset < 0 {
		offset = 0
	}

	return limit, offset
}

// Index uses limit, offset, and a slice's length to compute start and end indices for said slice.
func Index(limit, offset, length int) (start, end int) {
	if offset > length {
		return length, length
	} else if limit+offset > length {
		return offset, length
	}

	return offset, offset + limit
}

// MaxItemsPerPage is used to prevent DoS attacks against large lists by limiting the items per page to 500.
func MaxItemsPerPage(max, is int) int {
	if is > max {
		return max
	}
	return is
}

func header(u *fasthttp.URI, rel string, limit, offset int) string {
	q := u.QueryArgs()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(offset))
	// u.RawQuery = q.String()
	return fmt.Sprintf("<%s>; rel=\"%s\"", q.String(), rel)
}

// Header adds an http header for pagination using a responsewriter where backwards compatibility is required.
// The header will contain links any combination of the first, last, next, or previous (prev) pages in a paginated list (given a limit and an offset, and optionally a total).
// If total is not set, then no "last" page will be calculated.
// If no limit is provided, then it will default to 1.
func Header(c *fiber.Ctx, u *fasthttp.URI, pageSize int64, limit, offset int) {
	total := int(pageSize)
	if offset < 0 {
		offset = 0
	}

	if limit <= 0 {
		limit = 1
	}

	c.Set("X-Total-Count", strconv.Itoa(total))

	// lastOffset will either equal the offset required to contain the remainder,
	// or the limit.
	var lastOffset int
	if total%limit == 0 {
		lastOffset = total - limit
	} else {
		lastOffset = (total / limit) * limit
	}

	// Check for last page
	if offset >= lastOffset {
		if total == 0 {
			c.Set("Link", strings.Join(pq.StringArray{
				header(u, "first", limit, 0),
				header(u, "next", limit, ((offset/limit)+1)*limit),
				header(u, "prev", limit, ((offset/limit)-1)*limit),
			}, ","))
			return
		}

		if total < limit {
			c.Set("link", header(u, "first", total, 0))
			return
		}

		c.Set("Link", strings.Join(pq.StringArray{
			header(u, "first", limit, 0),
			header(u, "prev", limit, lastOffset-limit),
		}, ","))
		return
	}

	if offset < limit {
		c.Set("Link", strings.Join(pq.StringArray{
			header(u, "next", limit, limit),
			header(u, "last", limit, lastOffset),
		}, ","))
		return
	}

	c.Set("Link", strings.Join(pq.StringArray{
		header(u, "first", limit, 0),
		header(u, "next", limit, ((offset/limit)+1)*limit),
		header(u, "prev", limit, ((offset/limit)-1)*limit),
		header(u, "last", limit, lastOffset),
	}, ","))
	return
}
