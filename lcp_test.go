// Copyright © 2016 Abcum Ltd
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

package lcp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("Find lowest common prefix of -", t, func() {
		res := LCP()
		So(res, ShouldBeNil)
	})

	Convey("Find lowest common prefix of nil", t, func() {
		res := LCP(nil)
		So(res, ShouldBeNil)
	})

	Convey("Find lowest common prefix of 1 item", t, func() {
		res := LCP([]byte("tree"))
		So(res, ShouldResemble, []byte("tree"))
	})

	Convey("Find lowest common prefix of 2 identical items", t, func() {
		res := LCP([]byte("tree"), []byte("tree"))
		So(res, ShouldResemble, []byte("tree"))
	})

	Convey("Find lowest common prefix of 2 extendable items", t, func() {
		res := LCP([]byte("tree"), []byte("treetrunk"))
		So(res, ShouldResemble, []byte("tree"))
	})

	Convey("Find lowest common prefix of 2 similar items", t, func() {
		res := LCP([]byte("tree"), []byte("trie"))
		So(res, ShouldResemble, []byte("tr"))
	})

	Convey("Find lowest common prefix of 2 unmatchable items", t, func() {
		res := LCP([]byte("ptree"), []byte("btree"))
		So(res, ShouldBeNil)
	})

	Convey("Find lowest common prefix of 3 similar items", t, func() {
		res := LCP([]byte("internet"), []byte("interweb"), []byte("interstellar"))
		So(res, ShouldResemble, []byte("inter"))
	})

	Convey("Find lowest common prefix of 3 unmatchable items", t, func() {
		res := LCP([]byte("internet"), []byte(""), []byte("interstellar"))
		So(res, ShouldBeNil)
	})

}
