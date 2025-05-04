package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchFeed(t *testing.T) {
	sever := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:atom="http://www.w3.org/2005/Atom" version="2.0">
<channel>
<title>Lane's Blog</title>
<link>https://wagslane.dev/</link>
<description>Recent content on Lane's Blog</description>
<generator>Hugo -- gohugo.io</generator>
<language>en-us</language>
<lastBuildDate>Sun, 08 Jan 2023 00:00:00 +0000</lastBuildDate>
<atom:link href="https://wagslane.dev/index.xml" rel="self" type="application/rss+xml"/>
<item>
<title>The Zen of Proverbs</title>
<link>https://wagslane.dev/posts/zen-of-proverbs/</link>
<pubDate>Sun, 08 Jan 2023 00:00:00 +0000</pubDate>
<guid>https://wagslane.dev/posts/zen-of-proverbs/</guid>
<description><![CDATA[20 rules of thumb for writing better software. Optimize for simplicity first Write code for humans, not computers Reading is more important than writing Any style is fine, as long as it's black There should be one way to do it, but seriously this time Hide the sharp knives]]></description>
</item>
<item>
<title>College: A Solution in Search of a Problem</title>
<link>https://wagslane.dev/posts/college-a-solution-in-search-of-a-problem/</link>
<pubDate>Sat, 17 Dec 2022 00:00:00 +0000</pubDate>
<guid>https://wagslane.dev/posts/college-a-solution-in-search-of-a-problem/</guid>
<description><![CDATA[College has been prescribed almost universally by the parents of the last ~40 years as the solution to life's problems.]]></description>
</item>
</channel>
</rss>`))
    }))
	defer sever.Close()
	feed, err := fetchFeed(context.Background(), sever.URL);
	if err != nil {
		t.Errorf("Fail to get feed with error: %v", err)
	}
	if feed.Channel.Title != "Lane's Blog" {
        t.Errorf("Expected title 'Lane's Blog', got %s", feed.Channel.Title)
    }

    if len(feed.Channel.Item) != 2 {
        t.Errorf("Expected 2 items, got %d", len(feed.Channel.Item))
    }
}
