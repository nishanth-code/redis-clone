package store

type KVStore interface {
    Set(key string, value string)
    Get(key string) (string, bool)
    Delete(key string)
    Exists(key string) bool
}