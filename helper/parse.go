package helper

import (
	"fmt"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

// ParseConnectionOptions parses values for max_conns, max_idle_conns, max_conn_lifetime from DSNs.
// It also returns the URI without those query parameters.
func ParseConnectionOptions(dsn string) (maxConns int, maxIdleConns int, maxConnLifetime, maxIdleConnTime time.Duration, cleanedDSN string) {
	maxConns = maxParallelism() * 2
	maxIdleConns = maxParallelism()
	maxConnLifetime = time.Duration(0)
	maxIdleConnTime = time.Duration(0)
	cleanedDSN = dsn

	parts := strings.Split(dsn, "?")
	if len(parts) != 2 {
		fmt.Println("No SQL connection options have been defined, falling back to default connection options.")
		return
	}

	query, err := url.ParseQuery(parts[1])
	if err != nil {
		fmt.Println("Unable to parse SQL DSN query, falling back to default connection options.")
		return
	}

	if v := query.Get("max_conns"); v != "" {
		s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Printf(`SQL DSN query parameter "max_conns" value %v could not be parsed to int, falling back to default value %d`, v, maxConns)
		} else {
			maxConns = int(s)
		}
		query.Del("max_conns")
	}

	if v := query.Get("max_idle_conns"); v != "" {
		s, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Printf(`SQL DSN query parameter "max_idle_conns" value %v could not be parsed to int, falling back to default value %d`, v, maxIdleConns)
		} else {
			maxIdleConns = int(s)
		}
		query.Del("max_idle_conns")
	}

	if v := query.Get("max_conn_lifetime"); v != "" {
		s, err := time.ParseDuration(v)
		if err != nil {
			fmt.Printf(`SQL DSN query parameter "max_conn_lifetime" value %v could not be parsed to duration, falling back to default value %d`, v, maxConnLifetime)
		} else {
			maxConnLifetime = s
		}
		query.Del("max_conn_lifetime")
	}

	if v := query.Get("max_conn_idle_time"); v != "" {
		s, err := time.ParseDuration(v)
		if err != nil {
			fmt.Printf(`SQL DSN query parameter "max_conn_idle_time" value %v could not be parsed to duration, falling back to default value %d`, v, maxIdleConnTime)
		} else {
			maxIdleConnTime = s
		}
		query.Del("max_conn_idle_time")
	}
	cleanedDSN = fmt.Sprintf("%s?%s", parts[0], query.Encode())

	return
}

// FinalizeDSN will return a finalized DSN URI.
func FinalizeDSN(dsn string) string {
	if strings.HasPrefix(dsn, "mysql://") {
		var q url.Values
		parts := strings.SplitN(dsn, "?", 2)

		if len(parts) == 1 {
			q = make(url.Values)
		} else {
			var err error
			q, err = url.ParseQuery(parts[1])
			if err != nil {
				fmt.Println("Unable to parse SQL DSN query, could not finalize the DSN URI.")
				return dsn
			}
		}

		q.Set("multiStatements", "true")
		q.Set("parseTime", "true")

		return fmt.Sprintf("%s?%s", parts[0], q.Encode())
	}

	return dsn
}
