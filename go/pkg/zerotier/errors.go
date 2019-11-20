/*
 * Copyright (c)2019 ZeroTier, Inc.
 *
 * Use of this software is governed by the Business Source License included
 * in the LICENSE.TXT file in the project's root directory.
 *
 * Change Date: 2023-01-01
 *
 * On the date above, in accordance with the Business Source License, use
 * of this software will be governed by version 2.0 of the Apache License.
 */
/****/

package zerotier

// Err is a basic string error type for ZeroTier
type Err string

func (e Err) Error() string { return (string)(e) }

// Simple ZeroTier Errors
const (
	ErrInternal                 Err = "internal error"
	ErrNodeInitFailed           Err = "unable to initialize core Node instance"
	ErrInvalidMACAddress        Err = "invalid MAC address"
	ErrInvalidZeroTierAddress   Err = "invalid ZeroTier address"
	ErrInvalidNetworkID         Err = "invalid network ID"
	ErrInvalidParameter         Err = "invalid parameter"
	ErrTapInitFailed            Err = "unable to create native Tap instance"
	ErrUnrecognizedIdentityType Err = "unrecognized identity type"
	ErrInvalidKey               Err = "invalid key data"
	ErrInvalidSignature         Err = "invalid signature"
	ErrSecretKeyRequired        Err = "secret key required"
)

// APIErr is returned by the JSON API when a call fails
type APIErr struct {
	Reason string
}

func (e *APIErr) Error() string { return e.Reason }