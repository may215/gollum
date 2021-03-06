// Copyright 2015-2017 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package producer

import (
	"github.com/trivago/gollum/core"
	"sync"
	"time"
)

// Null producer plugin
// This producer does nothing and provides only bare-bone configuration (i.e.
// enabled and streams). Use this producer to test consumer performance.
type Null struct {
	core.DirectProducer `gollumdoc:"embed_type"`
	control             chan core.PluginControl
	streams             []core.MessageStreamID
}

func init() {
	core.TypeRegistry.Register(Null{})
}

// Configure initializes the basic members
func (prod *Null) Configure(conf core.PluginConfigReader) error {
	return prod.DirectProducer.Configure(conf)
}

// Enqueue simply ignores the message
func (prod *Null) Enqueue(msg core.Message, timeout *time.Duration) {
}

// Produce starts a control loop only
func (prod *Null) Produce(threads *sync.WaitGroup) {
	prod.ControlLoop()
}
