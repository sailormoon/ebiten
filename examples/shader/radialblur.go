// Copyright 2020 The Ebiten Authors
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

// +build ignore

package main

var Time float
var Cursor vec2
var ScreenSize vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	dir := normalize(position.xy - Cursor)
	clr := image2TextureAt(texCoord)

	samples := [10]float{
		-22, -14, -8, -4, -2, 2, 4, 8, 14, 22,
	}
	// TODO: Add len(samples)
	sum := clr
	for i := 0; i < 10; i++ {
		// TODO: Consider the source region not to violate the region.
		sum += image2TextureAt(texCoord + dir*samples[i]/image2TextureSize())
	}
	sum /= 10 + 1

	dist := distance(position.xy, Cursor)
	t := clamp(dist/256, 0, 1)
	return mix(clr, sum, t)
}