// Copyright 2018 The Prometheus Authors
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

//go:build ignore
// +build ignore

package main

import (
	"log"
	"time"

	"github.com/shurcooL/vfsgen"

	"github.com/tyr1k/alertmanager/asset"
	"github.com/tyr1k/alertmanager/pkg/modtimevfs"
)

func main() {
	fs := modtimevfs.New(asset.Assets, time.Unix(1, 0))
	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "asset",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
