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

import (
	"fmt"
	"time"

	"github.com/companyinfo/jsonapi"
)

// Blog is a model representing a blog site
type Blog struct {
	ID            int       `jsonapi:"primary,blogs"`
	Title         string    `jsonapi:"attr,title"`
	Posts         []*Post   `jsonapi:"relation,posts"`
	CurrentPost   *Post     `jsonapi:"relation,current_post"`
	CurrentPostID int       `jsonapi:"attr,current_post_id"`
	CreatedAt     time.Time `jsonapi:"attr,created_at"`
	ViewCount     int       `jsonapi:"attr,view_count"`
}

// Post is a model representing a post on a blog
type Post struct {
	ID       int        `jsonapi:"primary,posts"`
	BlogID   int        `jsonapi:"attr,blog_id"`
	Title    string     `jsonapi:"attr,title"`
	Body     string     `jsonapi:"attr,body"`
	Comments []*Comment `jsonapi:"relation,comments"`
}

// Comment is a model representing a user submitted comment
type Comment struct {
	ID     int    `jsonapi:"primary,comments"`
	PostID int    `jsonapi:"attr,post_id"`
	Body   string `jsonapi:"attr,body"`
}

// JSONAPILinks implements the Linkable interface for a blog
func (blog Blog) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("https://example.com/blogs/%d", blog.ID),
	}
}

// JSONAPIRelationshipLinks implements the RelationshipLinkable interface for a blog
func (blog Blog) JSONAPIRelationshipLinks(relation string) *jsonapi.Links {
	if relation == "posts" {
		return &jsonapi.Links{
			"related": fmt.Sprintf("https://example.com/blogs/%d/posts", blog.ID),
		}
	}
	if relation == "current_post" {
		return &jsonapi.Links{
			"related": fmt.Sprintf("https://example.com/blogs/%d/current_post", blog.ID),
		}
	}
	return nil
}

// JSONAPIMeta implements the Metable interface for a blog
func (blog Blog) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"detail": "extra details regarding the blog",
	}
}

// JSONAPIRelationshipMeta implements the RelationshipMetable interface for a blog
func (blog Blog) JSONAPIRelationshipMeta(relation string) *jsonapi.Meta {
	if relation == "posts" {
		return &jsonapi.Meta{
			"detail": "posts meta information",
		}
	}
	if relation == "current_post" {
		return &jsonapi.Meta{
			"detail": "current post meta information",
		}
	}
	return nil
}
