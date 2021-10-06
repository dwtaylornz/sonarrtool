/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

type Library struct {
	XMLName         xml.Name `xml:"MediaContainer"`
	Text            string   `xml:",chardata"`
	Size            string   `xml:"size,attr"`
	AllowSync       string   `xml:"allowSync,attr"`
	Identifier      string   `xml:"identifier,attr"`
	MediaTagPrefix  string   `xml:"mediaTagPrefix,attr"`
	MediaTagVersion string   `xml:"mediaTagVersion,attr"`
	Title1          string   `xml:"title1,attr"`
	Directory       []struct {
		Text             string `xml:",chardata"`
		AllowSync        string `xml:"allowSync,attr"`
		Art              string `xml:"art,attr"`
		Composite        string `xml:"composite,attr"`
		Filters          string `xml:"filters,attr"`
		Refreshing       string `xml:"refreshing,attr"`
		Thumb            string `xml:"thumb,attr"`
		Key              string `xml:"key,attr"`
		Type             string `xml:"type,attr"`
		Title            string `xml:"title,attr"`
		Agent            string `xml:"agent,attr"`
		Scanner          string `xml:"scanner,attr"`
		Language         string `xml:"language,attr"`
		Uuid             string `xml:"uuid,attr"`
		UpdatedAt        string `xml:"updatedAt,attr"`
		CreatedAt        string `xml:"createdAt,attr"`
		ScannedAt        string `xml:"scannedAt,attr"`
		Content          string `xml:"content,attr"`
		Directory        string `xml:"directory,attr"`
		ContentChangedAt string `xml:"contentChangedAt,attr"`
		Hidden           string `xml:"hidden,attr"`
		Location         struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Path string `xml:"path,attr"`
		} `xml:"Location"`
	} `xml:"Directory"`
}

type Media struct {
	XMLName             xml.Name `xml:"MediaContainer"`
	Text                string   `xml:",chardata"`
	Size                string   `xml:"size,attr"`
	AllowSync           string   `xml:"allowSync,attr"`
	Art                 string   `xml:"art,attr"`
	Identifier          string   `xml:"identifier,attr"`
	LibrarySectionID    string   `xml:"librarySectionID,attr"`
	LibrarySectionTitle string   `xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string   `xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string   `xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string   `xml:"mediaTagVersion,attr"`
	Thumb               string   `xml:"thumb,attr"`
	Title1              string   `xml:"title1,attr"`
	Title2              string   `xml:"title2,attr"`
	ViewGroup           string   `xml:"viewGroup,attr"`
	ViewMode            string   `xml:"viewMode,attr"`
	Video               []struct {
		Text                  string `xml:",chardata"`
		RatingKey             string `xml:"ratingKey,attr"`
		Key                   string `xml:"key,attr"`
		Guid                  string `xml:"guid,attr"`
		Studio                string `xml:"studio,attr"`
		Type                  string `xml:"type,attr"`
		Title                 string `xml:"title,attr"`
		ContentRating         string `xml:"contentRating,attr"`
		Summary               string `xml:"summary,attr"`
		AudienceRating        string `xml:"audienceRating,attr"`
		Year                  string `xml:"year,attr"`
		Tagline               string `xml:"tagline,attr"`
		Thumb                 string `xml:"thumb,attr"`
		Art                   string `xml:"art,attr"`
		Duration              string `xml:"duration,attr"`
		OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
		AddedAt               string `xml:"addedAt,attr"`
		UpdatedAt             string `xml:"updatedAt,attr"`
		AudienceRatingImage   string `xml:"audienceRatingImage,attr"`
		PrimaryExtraKey       string `xml:"primaryExtraKey,attr"`
		ViewCount             string `xml:"viewCount,attr"`
		LastViewedAt          string `xml:"lastViewedAt,attr"`
		ChapterSource         string `xml:"chapterSource,attr"`
		TitleSort             string `xml:"titleSort,attr"`
		SkipCount             string `xml:"skipCount,attr"`
		ViewOffset            string `xml:"viewOffset,attr"`
		Media                 []struct {
			Text                  string `xml:",chardata"`
			ID                    string `xml:"id,attr"`
			Duration              string `xml:"duration,attr"`
			Bitrate               string `xml:"bitrate,attr"`
			Width                 string `xml:"width,attr"`
			Height                string `xml:"height,attr"`
			AspectRatio           string `xml:"aspectRatio,attr"`
			AudioChannels         string `xml:"audioChannels,attr"`
			AudioCodec            string `xml:"audioCodec,attr"`
			VideoCodec            string `xml:"videoCodec,attr"`
			VideoResolution       string `xml:"videoResolution,attr"`
			Container             string `xml:"container,attr"`
			VideoFrameRate        string `xml:"videoFrameRate,attr"`
			VideoProfile          string `xml:"videoProfile,attr"`
			OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
			AudioProfile          string `xml:"audioProfile,attr"`
			Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
			Part                  struct {
				Text                  string `xml:",chardata"`
				ID                    string `xml:"id,attr"`
				Key                   string `xml:"key,attr"`
				Duration              string `xml:"duration,attr"`
				File                  string `xml:"file,attr"`
				Size                  string `xml:"size,attr"`
				Container             string `xml:"container,attr"`
				VideoProfile          string `xml:"videoProfile,attr"`
				AudioProfile          string `xml:"audioProfile,attr"`
				Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
				OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
				HasThumbnail          string `xml:"hasThumbnail,attr"`
			} `xml:"Part"`
		} `xml:"Media"`
		Genre []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Genre"`
		Director []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Director"`
		Writer []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Writer"`
		Country []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Country"`
		Role []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag,attr"`
		} `xml:"Role"`
	} `xml:"Video"`
}

// var plex_server = &server
// var plex_token = &token

// duplicatesCmd represents the duplicates command
var duplicatesCmd = &cobra.Command{
	Use:     "duplicates",
	Aliases: []string{"dups"},
	Short:   "Finds all duplicate videos",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Server:", server, "Token:", token)
		// fmt.Println("duplicates called")

		var url_libraries = "http://" + server + "/library/sections/all?X-Plex-Token=" + token

		// fmt.Println(url_libraries)

		// fmt.Println("Connecting to Plex")
		response, err := http.Get(url_libraries)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		var library Library
		xml.Unmarshal(responseData, &library)

		// for each library
		for _, library := range library.Directory {

			fmt.Println("Library:", library.Title, ", Type:", library.Type, ", Key:", library.Key)

			if library.Type == "movie" {
				// fmt.Println("This is a movie")
				var url_movie = "http://" + server + "/library/sections/" + library.Key + "/all?duplicate=1&X-Plex-Token=" + token
				// fmt.Println(url_movie)
				response, err := http.Get(url_movie)

				if err != nil {
					fmt.Print(err.Error())
					os.Exit(1)
				}

				responseData, err := ioutil.ReadAll(response.Body)

				if err != nil {
					log.Fatal(err)
				}

				var media Media
				xml.Unmarshal(responseData, &media)

				process_media(media)

			}

			if library.Type == "show" {
				//fmt.Println("This is a show")
				var url_show = "http://" + server + "/library/sections/" + library.Key + "/search?type=4&duplicate=1&X-Plex-Token=" + token
				// fmt.Println(url_show)

				response, err := http.Get(url_show)

				if err != nil {
					fmt.Print(err.Error())
					os.Exit(1)
				}

				responseData, err := ioutil.ReadAll(response.Body)

				if err != nil {
					log.Fatal(err)
				}

				var media Media
				xml.Unmarshal(responseData, &media)

				process_media(media)

			}

		}

	},
}

func init() {
	rootCmd.AddCommand(duplicatesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// duplicatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// duplicatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func process_media(media Media) {

	for _, video := range media.Video {
		fmt.Println(video.Title, video.Key)
		// fmt.Println("    ", video.Media)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "Size", "Width", "Codec"})
		t.SortBy([]table.SortBy{
			{Name: "Codec", Mode: table.Dsc},
			{Name: "Width", Mode: table.Asc},
			{Name: "Size", Mode: table.Asc},
		})

		// width := 0
		for _, media := range video.Media {

			t.AppendRows([]table.Row{
				{media.ID, media.Part.Size, media.Width, media.VideoCodec},
			})

			// fmt.Println("     ID:", media.ID, "Size:", media.Part.Size, ", Width:", media.Width, ", Codec:", media.VideoCodec)
			// fmt.Println(score)

		}
		t.SetStyle(table.StyleLight)
		t.Render()

	}

}
