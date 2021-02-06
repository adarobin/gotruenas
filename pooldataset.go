package gotruenas

import (
	"context"
	"fmt"
	"net/http"
)

const poolDatasetPath = basePath + "/pool/dataset"

// PoolDataset defines model for a pool dataset.
type PoolDataset struct{}

// PoolDatasetList defines model for a pool dataset.
type PoolDatasetList struct{}

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

// PoolDatasetOp handles communication with the Pool Dataset related methods of the
// TrueNAS REST API
type PoolDatasetOp struct {
	client *Client
}

// Create a new activation key
func (s *PoolDatasetOp) Create(ctx context.Context, body PoolDatasetCreate) (*PoolDataset, *http.Response, error) {
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
func (s *PoolDatasetOp) Delete(ctx context.Context, id int) (*http.Response, error) {
	path := fmt.Sprintf("%s/id/%d", poolDatasetPath, id)

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
func (s *PoolDatasetOp) Get(ctx context.Context, id int) (*PoolDataset, *http.Response, error) {
	path := fmt.Sprintf("%s/id/%d", poolDatasetPath, id)

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
func (s *PoolDatasetOp) List(ctx context.Context, opt *ListOptions) (*PoolDatasetList, *http.Response, error) {
	path := poolDatasetPath

	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(PoolDatasetList)
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}

	return list, resp, err
}
