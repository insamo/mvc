package controllers

import (
	"errors"

	"github.com/insamo/mvc/utils/crypto"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var (
	// ErrUnexpected represents an unexpected error
	ErrUnexpected = errors.New("An unexpected error as occur")

	// ErrNotFoundOnCache represents an error when the key was not found on cache
	ErrNotFoundOnCache = errors.New("Not found on cache")

	ErrNotFound = errors.New("Not found")

	// ErrGetCacheValue represents an error when an error occur when get cache value
	ErrGetCacheValue = errors.New("Not found on cache")
)

func SetEtag(ctx iris.Context, r interface{}) {
	etag := crypto.GenerateSha256Hash(r)
	ctx.Header(context.ETagHeaderKey, etag)
	ctx.Values().Set(context.ETagHeaderKey, etag)
}
