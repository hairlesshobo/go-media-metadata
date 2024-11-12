// =================================================================================
//
//		go-media-metadata - https://www.github.com/hairlesshobo/go-media-metadata
//
//	 go-xlive currently does nothing - but will hoepfully one day allow for
//   easy extraction of metadata from mov files and maybe more.
//
//		Copyright (c) 2024 Steve Cross <flip@foxhollow.cc>
//
//		Licensed under the Apache License, Version 2.0 (the "License");
//		you may not use this file except in compliance with the License.
//		You may obtain a copy of the License at
//
//		     http://www.apache.org/licenses/LICENSE-2.0
//
//		Unless required by applicable law or agreed to in writing, software
//		distributed under the License is distributed on an "AS IS" BASIS,
//		WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//		See the License for the specific language governing permissions and
//		limitations under the License.
//
// =================================================================================

package main

import (
	"fmt"
	// "go-media-metadata/mediameta"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no MOV file wad provided")
		return
	}

	// sessionInfo := xlive.ReadSessionFile(os.Args[1])

	// fmt.Printf("%+v\n", sessionInfo)
}
