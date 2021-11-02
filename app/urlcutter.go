package app

import (
	"urlcutter/domain"
	"urlcutter/infrastructure/localservices"
)

type urlcutterSrvice struct {
	keyGenerator localservices.KeyGenerator
	repo         UrlSaver
}

func NewUrlcutterSrvice(keyGenerator localservices.KeyGenerator, repo UrlSaver) *urlcutterSrvice {
	return &urlcutterSrvice{keyGenerator: keyGenerator, repo: repo}
}

func (u *urlcutterSrvice) MakeKey(s string) (string, error) {
	key := u.keyGenerator.Generate()
	url := domain.LongURL{LongURLData: s}
	err := u.repo.Save(url, key)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (u *urlcutterSrvice) GetURL(s string) (string, error) {
	url, err := u.repo.Get(s)
	if err != nil {
		return "", err
	}
	return url, nil
}
