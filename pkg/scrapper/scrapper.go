package scrapper

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/keshavsharma98/Blog-Aggregator/internal/database"
)

// RssScraper performs RSS scraping from multiple feeds using GO ROUTINES and updates the provided database with the fetched data.
//
// Parameters:
// - db: The database connection to use for storing the scraped data.
// - concurrency: The number of concurrent goroutines to use during scraping.
// - duration: The duration to wait between consecutive scrapes.
func RssScraper(db *database.Queries, concurrency int, duration time.Duration) {
	log.Printf("Started scraping on %v goroutines every %s duration \n", concurrency, duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		feeds_arr, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("Error while fetching the feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, f := range feeds_arr {
			wg.Add(1)

			go scrapFeed(wg, db, f)
		}
		wg.Wait()
	}
}

func scrapFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	err := db.MarkFetched(context.Background(), database.MarkFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:     time.Now().UTC(),
	})
	if err != nil {
		log.Println("Error while marking feed as fetched: ", err)
		return
	}

	rssFeed, err := scrapFeedURL(feed.Url)
	if err != nil {
		log.Printf("Error while scrapping feed url: %v\n", err)
		return
	}

	for _, f := range rssFeed.Channel.Item {
		// desc := sql.NullString{}
		// if f.Description != "" {
		// 	desc.String = f.Description
		// 	desc.Valid = true
		// }

		pub_date, err := time.Parse(time.RFC1123Z, f.PubDate)
		if err != nil {
			log.Printf("could not parse date %v with err %v\n", f.PubDate, err)
			return
		}

		err = db.CreatePosts(context.Background(), database.CreatePostsParams{
			ID:          uuid.New(),
			Title:       f.Title,
			Url:         f.Link,
			Description: &f.Description,
			PublishedAt: pub_date,
			FeedID:      feed.ID,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("post creation failed with err :", err)
			return
		}
	}
	log.Printf("Scrapped %v posts from the feed ->  %s\n", len(rssFeed.Channel.Item), feed.Name)
}

func scrapFeedURL(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
