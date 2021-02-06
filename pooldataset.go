package gotruenas

import (
	"context"
	"fmt"
	"net/http"
)

const poolDatasetPath = basePath + "/pool/dataset"

type value struct {
	Parsed   *string `json:"parsed,omitempty"`
	RawValue *string `json:"rawvalue,omitempty"`
	Value    *string `json:"value,omitempty"`
	Source   *string `json:"source,omitempty"`
}

// PoolDataset defines model for a pool dataset.
type PoolDataset struct {
	ID                    *string        `json:"id"`
	Type                  *string        `json:"type"`
	Children              *[]PoolDataset `json:"children"`
	Name                  *string        `json:"name"`
	Pool                  *string        `json:"pool"`
	Encryption            *bool          `json:"encryption"`
	EncryptionRoot        *bool          `json:"encryption_root"`
	KeyLoaded             *bool          `json:"key_loaded"`
	Mountpoint            *string        `json:"mountpoint"`
	Deduplication         *value         `json:"deduplication"`
	ACLMode               *value         `json:"aclmode"`
	ACLType               *value         `json:"acltype"`
	Xattr                 *value         `json:"xattr"`
	Atime                 *value         `json:"atime"`
	CaseSensitivity       *value         `json:"casesensitivity"`
	Exec                  *value         `json:"exec"`
	Sync                  *value         `json:"sync"`
	Compression           *value         `json:"compression"`
	CompressRatio         *value         `json:"compressratio"`
	Origin                *value         `json:"origin"`
	Quota                 *value         `json:"quota"`
	RefQuota              *value         `json:"refquota"`
	Reservation           *value         `json:"reservation"`
	RefReservation        *value         `json:"refreservation"`
	Copies                *value         `json:"copies"`
	SnapDir               *value         `json:"snapdir"`
	ReadOnly              *value         `json:"readonly"`
	RecordSize            *value         `json:"recordsize"`
	KeyFormat             *value         `json:"key_format"`
	EncryptionAlgorithm   *value         `json:"encryption_algorithm"`
	Used                  *value         `json:"used"`
	Available             *value         `json:"available"`
	SpecialSmallBlockSize *value         `json:"special_small_block_size"`
	Pbkdf2iters           *value         `json:"pbkdf2iters"`
	Locked                *bool          `json:"locked"`
}

// PoolDatasetCreate defines model for creating a pool dataset.
type PoolDatasetCreate struct {
	Name                  *string `json:"name"`
	Type                  *string `json:"type"`
	VolSize               *int    `json:"volsize,omitempty"`
	VolBlockSize          *int    `json:"volblocksize,omitempty"`
	Sparse                *bool   `json:"sparse,omitempty"`
	ForceSize             *bool   `json:"force_size,omitempty"`
	Comments              *string `json:"comments,omitempty"`
	Sync                  *string `json:"sync,omitempty"`
	Compression           *string `json:"compression,omitempty"`
	ATime                 *string `json:"atime,omitempty"`
	Exec                  *string `json:"exec,omitempty"`
	ManagedBy             *string `json:"managedby,omitempty"`
	Quota                 *int    `json:"quota,omitempty"`
	QuotaWarning          *int    `json:"quota_warning,omitempty"`
	QuotaCritical         *int    `json:"quota_critical,omitempty"`
	RefQuota              *int    `json:"refquota,omitempty"`
	RefQuotaWarning       *int    `json:"refquota_warning,omitempty"`
	RefquotaCritical      *int    `json:"refquota_critical,omitempty"`
	Reservation           *int    `json:"reservation,omitempty"`
	RefReservation        *int    `json:"refreservation,omitempty"`
	SpecialSmallBlockSize *int    `json:"special_small_block_size,omitempty"`
	Copies                *int    `json:"copies,omitempty"`
	SnapDir               *string `json:"snapdir,omitempty"`
	Deduplication         *string `json:"deduplication,omitempty"`
	ReadOnly              *string `json:"readonly,omitempty"`
	RecordSize            *string `json:"recordsize,omitempty"`
	CaseSensitivity       *string `json:"casesensitivity,omitempty"`
	ACLMode               *string `json:"aclmode,omitempty"`
	ACLType               *string `json:"acltype,omitempty"`
	ShareType             *string `json:"share_type,omitempty"`
	Xattr                 *string `json:"xattr,omitempty"`
	Encryption            *bool   `json:"encryption,omitempty"`
	InheritEncryption     *bool   `json:"inherit_encryption,omitempty"`
	EncryptionOptions     struct {
		GenerateKey *bool   `json:"generate_key,omitempty"`
		Pbkdf2iters *int    `json:"pbkdf2iters,omitempty"`
		Algorithm   *string `json:"algorithm,omitempty"`
		Passphrase  *string `json:"passphrase,omitempty"`
		Key         *string `json:"key,omitempty"`
	} `json:"encryption_options,omitempty"`
}

// PoolDatasetsOp handles communication with the Pool Dataset related methods of the
// TrueNAS REST API
type PoolDatasetsOp struct {
	client *Client
}

// PoolDatasets is an interface for interacting with
// Pool Datasets
type PoolDatasets interface {
	Create(ctx context.Context, body PoolDatasetCreate) (*PoolDataset, *http.Response, error)
	Delete(ctx context.Context, id string) (*http.Response, error)
	Get(ctx context.Context, id string) (*PoolDataset, *http.Response, error)
	List(ctx context.Context, opt *ListOptions) (*[]PoolDataset, *http.Response, error)
}

// Create a new activation key
func (s *PoolDatasetsOp) Create(ctx context.Context, body PoolDatasetCreate) (*PoolDataset, *http.Response, error) {
	path := poolDatasetPath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, nil, err
	}

	dataset := new(PoolDataset)
	resp, err := s.client.Do(ctx, req, dataset)
	if err != nil {
		return nil, resp, err
	}

	return dataset, resp, nil
}

// Delete an activation key by its ID
func (s *PoolDatasetsOp) Delete(ctx context.Context, id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/id/%s", poolDatasetPath, id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Get a single pool dataset by its ID
func (s *PoolDatasetsOp) Get(ctx context.Context, id string) (*PoolDataset, *http.Response, error) {
	path := fmt.Sprintf("%s/id/%s", poolDatasetPath, id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	dataset := new(PoolDataset)
	resp, err := s.client.Do(ctx, req, dataset)
	if err != nil {
		return nil, resp, err
	}

	return dataset, resp, nil
}

// List all activation keys or a filtered list of activation keys
func (s *PoolDatasetsOp) List(ctx context.Context, opt *ListOptions) (*[]PoolDataset, *http.Response, error) {
	path := poolDatasetPath

	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new([]PoolDataset)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, err
}
