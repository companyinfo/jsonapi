// Copyright 2024 Company.info B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "time"

func fixtureBlogCreate(i int) *Blog {
	return &Blog{
		ID:        1 * i,
		Title:     "Title 1",
		CreatedAt: time.Now(),
		Posts: []*Post{
			{
				ID:    1 * i,
				Title: "Foo",
				Body:  "Bar",
				Comments: []*Comment{
					{
						ID:   1 * i,
						Body: "foo",
					},
					{
						ID:   2 * i,
						Body: "bar",
					},
				},
			},
			{
				ID:    2 * i,
				Title: "Fuubar",
				Body:  "Bas",
				Comments: []*Comment{
					{
						ID:   1 * i,
						Body: "foo",
					},
					{
						ID:   3 * i,
						Body: "bas",
					},
				},
			},
		},
		CurrentPost: &Post{
			ID:    1 * i,
			Title: "Foo",
			Body:  "Bar",
			Comments: []*Comment{
				{
					ID:   1 * i,
					Body: "foo",
				},
				{
					ID:   2 * i,
					Body: "bar",
				},
			},
		},
	}
}

func fixtureBlogsList() (blogs []interface{}) {
	for i := 0; i < 10; i++ {
		blogs = append(blogs, fixtureBlogCreate(i))
	}

	return blogs
}
