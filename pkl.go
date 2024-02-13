// Copyright 2015 Matthew Holt and The Caddy Authors
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

package pkladapter

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/caddyserver/caddy/v2/caddyconfig"
)

func init() {
	caddyconfig.RegisterAdapter("pkl", Adapter{})
}

// Adapter adapts Pkl to Caddy JSON.
type Adapter struct{}

// Adapt converts the Pkl config in body to Caddy JSON.
func (a Adapter) Adapt(body []byte, options map[string]interface{}) (result []byte, warnings []caddyconfig.Warning, err error) {
	ctx := context.TODO()

	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions, func(opts *pkl.EvaluatorOptions) {
		pkl.PreconfiguredOptions(opts)
		opts.OutputFormat = "json"
	})
	if err != nil {
		return nil, nil, err
	}
	defer evaluator.Close()

	resultStr, err := evaluator.EvaluateOutputText(ctx, pkl.TextSource(string(body)))
	if err != nil {
		return nil, nil, err
	}

	return []byte(resultStr), nil, nil
}

// Interface guard
var _ caddyconfig.Adapter = (*Adapter)(nil)
