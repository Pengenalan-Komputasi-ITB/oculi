package storage

import (
	"sync"

	"github.com/Pengenalan-Komputasi-ITB/oculi/storage"
	"github.com/minio/minio-go/v7"
)

type (
	impl struct {
		cl      *minio.Client
		signer  *minio.Client
		region  string
		mu      sync.RWMutex
		buckets map[string]storage.Bucket
	}

	bucket struct {
		cl        *minio.Client
		signer    *minio.Client
		name      string
		parent    *impl
		isDeleted bool
		mu        sync.RWMutex
	}
)
