package lib

import (
    "github.com/Falkenfighter/GoRest"
    "fmt"
    "runtime"
)

type Ids []int
type Story struct {
    By          string
    Descendants int
    Id          int
    Kids        []int
    Score       int
    Text        string
    Time        int
    Title       string
    Type        string
    Url         string
}

func (s *Story) String() string {
    return fmt.Sprintf(s.Title)
}

var client = GoRest.MakeClient("https://hacker-news.firebaseio.com/v0")

func TopStories() ([]*Story, error) {
    ids, err := topStoryIds()
    if (err != nil) {
        return nil, err
    }

    runtime.GOMAXPROCS(len(*ids))
    stories := []*Story{}

    ch := make(chan *Story)
    for i, id := range *ids {
        go getStory(i, id, ch, 0)
    }

    for {
        select {
        case story := <-ch:
            stories = append(stories, story)
            if len(stories) == len(*ids) {
                return stories, nil
            }
        }
    }
    return stories, nil
}

func topStoryIds() (ids *Ids, err error) {
    ids = new(Ids)
    _, err = client.Path("topstories.json").Get(ids)
    return
}

func topStory(id int) (story *Story, err error) {
    story = new(Story)
    _, err = client.Path("item", fmt.Sprintf("%d.json", id)).Get(story)
    return
}

func getStory(i, id int, ch chan *Story, acc int) {
    story, err := topStory(id)
    if err != nil {
        if acc >= 2 {
            ch <- new(Story)
            return
        }
        getStory(i, id, ch, acc + 1)
    } else {
        ch <- story
    }
}