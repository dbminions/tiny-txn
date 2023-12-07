package mvcc

type MVCC interface {
	Get(key string, ts uint64) ([]byte, uint64, error)
	Set(key string, ts uint64, val []byte) error
	Del(key string, ts uint64) error

	Close()
}
