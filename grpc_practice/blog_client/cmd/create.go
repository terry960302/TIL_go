/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	blogpb "first_golang/grpc_practice/blog"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new blog post",
	Long: `Create a new blogpost on the server through gRPC. 
	
	A blog post requires an AuthorId, Title and Content.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// flags로부터 데이터를 불러옴.
		author, err := cmd.Flags().GetString("author")
		title, err := cmd.Flags().GetString("title")
		content, err := cmd.Flags().GetString("content")
		if err != nil {
			return err
		}
		// Create a blog protobuffer message
		blog := &blogpb.Blog{
			AuthorId: author,
			Title:    title,
			Content:  content,
		}
		// RPC call
		res, err := client.CreateBlog(
			context.TODO(),
			// 위에서 만든 blog message를 CreateBlog request message 안에 감싸줌.
			&blogpb.CreateBlogReq{
				Blog: blog,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Blog가 생성되었습니다. : %s\n", res.Blog.Id)
		return nil
	},
}

func init() {

	createCmd.Flags().StringP("author", "a", "", "Add an author")
	createCmd.Flags().StringP("title", "t", "", "A title for the blog")
	createCmd.Flags().StringP("content", "c", "", "The content for the blog")
	createCmd.MarkFlagRequired("author")
	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("content")

	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
