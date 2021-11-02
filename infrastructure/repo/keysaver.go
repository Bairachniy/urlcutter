package repo

import (
	"errors"
	"sync"
	"urlcutter/domain"
)

type keysaver struct {
	repo123 TemporaryRepo
	mu      sync.RWMutex
}

func NewKeysaver() *keysaver {
	repo := make(map[string]string, 0)
	return &keysaver{repo123: TemporaryRepo{repo}}
}

func (k *keysaver) Get(s string) (string, error) {
	url, ok := k.repo123.Repo[s]
	if !ok {
		return "", errors.New("no such url")
	}
	return url, nil
}

func (k *keysaver) Save(url domain.LongURL, key string) error {
	k.mu.RLock()
	k.repo123.Repo[key] = url.LongURLData
	k.mu.RUnlock()
	return nil
}

type TemporaryRepo struct {
	Repo map[string]string
}
