package cache

type RedisStruct struct {
	Name   string
	Expiry int
}
type RedisTimer struct {
	Time int
}

var (
	RedisKey = map[string]RedisStruct{
		"RESERVED_SEAT": {Name: "reserved_seat", Expiry: 15 * 60},
	}
)
