// Copyright 2015 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/prometheus/common/model"

	"github.com/prometheus/alertmanager/types"
)

type Iterator interface {
	Err() error
	Close()
}

type AlertIterator interface {
	Iterator
	Next() <-chan *types.Alert
}

// Alerts gives access to a set of alerts.
type Alerts interface {
	// Subscribe returns an iterator over active alerts that have not been
	// resolved and successfully notified about.
	// They are not guaranteed to be in chronological order.
	Subscribe() AlertIterator
	// GetPending returns an iterator over all alerts that have
	// pending notifications.
	GetPending() AlertIterator
	// Get returns the alert for a given fingerprint.
	Get(model.Fingerprint) (*types.Alert, error)
	// Put adds the given alert to the set.
	Put(...*types.Alert) error
}

// Silences gives access to silences.
type Silences interface {
	// The Silences provider must implement the Muter interface
	// for all its silences. The data provider may have access to an
	// optimized view of the data to perform this evaluation.
	types.Muter

	// All returns all existing silences.
	All() ([]*types.Silence, error)
	// Set a new silence.
	Set(*types.Silence) (uint64, error)
	// Del removes a silence.
	Del(uint64) error
	// Get a silence associated with a fingerprint.
	Get(uint64) (*types.Silence, error)
}

// Notifies provides information about pending and successful
// notifications.
type Notifies interface {
	Get(dest string, fps ...model.Fingerprint) ([]*types.Notify, error)
	// Set several notifies at once. All or none must succeed.
	Set(ns ...*types.Notify) error
}

type Config interface {
	// Reload initiates a configuration reload.
	Reload(...types.Reloadable) error
}