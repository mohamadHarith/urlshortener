package service

import "errors"

type URLShortener struct {
	urls map[string]string
}

func NewURLShortener() *URLShortener {
	return &URLShortener{}
}

func (u *URLShortener) Store(key, value string) {
	_, ok := u.urls[key]
	if !ok {
		u.urls[key] = value
	}
}

func (u *URLShortener) Get(key string) (string, error) {

	v, ok := u.urls[key]
	if !ok {
		return "", errors.New("not found")
	}

	return v, nil
}
