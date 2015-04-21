package internet

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type ASN2ASDescClient struct {
	conn redis.Conn
}

//NewASN2ASDescClient .
func NewASN2ASDescClient(conn redis.Conn) *ASN2ASDescClient {
	asnasd := ASN2ASDescClient{
		conn: conn,
	}
	return &asnasd
}

// importedDates fetches all imported dates from redis
func (i *ASN2ASDescClient) importedDates() ([]string, error) {
	return redis.Strings(i.conn.Do("SMEMBERS", "asd:imported_dates"))
}

// Current returns the latest known result for an IP2ASN lookup.
func (i *ASN2ASDescClient) Current(ASN int) string {
	allDates, err := i.importedDates()
	if err != nil {
		return ""
	}
	if len(allDates) < 0 {
		return ""
	}
	current := allDates[len(allDates)-1]
	result, err := redis.String(i.conn.Do("HGET", fmt.Sprintf("asd:%d", ASN), current))
	if err != nil {
		return ""
	}

	return result
}