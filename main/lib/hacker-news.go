package lib

import (
    "github.com/bndr/gopencils"
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

var api = gopencils.Api("https://hacker-news.firebaseio.com/v0")

func TopStories() ([]*Story, error) {
    ids, err := topStoryIds()
    if (err != nil) {
        return nil, err
    }

    runtime.GOMAXPROCS(len(*ids))
    stories := []*Story{}

    ch := make(chan *Story)
    for i, id := range *ids {
        go func(i, id int) {
            story, _ := topStory(id)
            ch <- story
        }(i, id)
    }

    for {
        select {
        case story := <-ch:
            stories = append(stories, story)
            if len(stories) == len(*ids) {
                return stories
            }
        }
    }
    return stories, nil
}

func topStoryIds() (ids *Ids, err error) {
    ids = new(Ids)
    _, err = api.Res("topstories.json", ids).Get()
    return
}

func topStory(id int) (story *Story, err error) {
    story = new(Story)
    _, err = api.Res("item").Res(fmt.Sprintf("%d.json", id), story).Get()
    return
}