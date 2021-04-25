// Code generated by go generate internal/cmd; DO NOT EDIT.
package pairs

import (
	"context"

	"github.com/aos-dev/go-storage/v3/pkg/httpclient"
	. "github.com/aos-dev/go-storage/v3/types"
)

// WithContentMd5 will apply content_md5 value to Options.
//
// ContentMd5
func WithContentMd5(v string) Pair {
	return Pair{
		Key:   "content_md5",
		Value: v,
	}
}

// WithContentType will apply content_type value to Options.
//
// ContentType
func WithContentType(v string) Pair {
	return Pair{
		Key:   "content_type",
		Value: v,
	}
}

// WithContext will apply context value to Options.
//
// Context
func WithContext(v context.Context) Pair {
	return Pair{
		Key:   "context",
		Value: v,
	}
}

// WithContinuationToken will apply continuation_token value to Options.
//
// ContinuationToken specify the continuation token for list
func WithContinuationToken(v string) Pair {
	return Pair{
		Key:   "continuation_token",
		Value: v,
	}
}

// WithCredential will apply credential value to Options.
//
// Credential specify how to provide credential for service or storage
func WithCredential(v string) Pair {
	return Pair{
		Key:   "credential",
		Value: v,
	}
}

// WithEndpoint will apply endpoint value to Options.
//
// Endpoint specify how to provide endpoint for service or storage
func WithEndpoint(v string) Pair {
	return Pair{
		Key:   "endpoint",
		Value: v,
	}
}

// WithExpire will apply expire value to Options.
//
// Expire specify when the url returned by reach will expire
func WithExpire(v int) Pair {
	return Pair{
		Key:   "expire",
		Value: v,
	}
}

// WithHTTPClientOptions will apply http_client_options value to Options.
//
// HTTPClientOptions
func WithHTTPClientOptions(v *httpclient.Options) Pair {
	return Pair{
		Key:   "http_client_options",
		Value: v,
	}
}

// WithInterceptor will apply interceptor value to Options.
//
// Interceptor
func WithInterceptor(v Interceptor) Pair {
	return Pair{
		Key:   "interceptor",
		Value: v,
	}
}

// WithIoCallback will apply io_callback value to Options.
//
// IoCallback specify what todo every time we read data from source
func WithIoCallback(v func([]byte)) Pair {
	return Pair{
		Key:   "io_callback",
		Value: v,
	}
}

// WithListMode will apply list_mode value to Options.
//
// ListMode
func WithListMode(v ListMode) Pair {
	return Pair{
		Key:   "list_mode",
		Value: v,
	}
}

// WithLocation will apply location value to Options.
//
// Location specify the location for service or storage
func WithLocation(v string) Pair {
	return Pair{
		Key:   "location",
		Value: v,
	}
}

// WithMultipartID will apply multipart_id value to Options.
//
// MultipartID
func WithMultipartID(v string) Pair {
	return Pair{
		Key:   "multipart_id",
		Value: v,
	}
}

// WithName will apply name value to Options.
//
// Name specify the storage name
func WithName(v string) Pair {
	return Pair{
		Key:   "name",
		Value: v,
	}
}

// WithOffset will apply offset value to Options.
//
// Offset specify offset for this request, storage will seek to this offset before read
func WithOffset(v int64) Pair {
	return Pair{
		Key:   "offset",
		Value: v,
	}
}

// WithPairPolicy will apply pair_policy value to Options.
//
// PairPolicy
func WithPairPolicy(v PairPolicy) Pair {
	return Pair{
		Key:   "pair_policy",
		Value: v,
	}
}

// WithSize will apply size value to Options.
//
// Size specify size for this request, storage will only read limited content data
func WithSize(v int64) Pair {
	return Pair{
		Key:   "size",
		Value: v,
	}
}

// WithWorkDir will apply work_dir value to Options.
//
// WorkDir specify the work dir for service or storage, every operation will be relative to this dir.
// work_dir MUST start with / for every storage services.
// work_dir will be default to / if not set.
// For fs storage service on windows platform, the behavior is undefined.
func WithWorkDir(v string) Pair {
	return Pair{
		Key:   "work_dir",
		Value: v,
	}
}
