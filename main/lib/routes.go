package lib
import (
    "net/http"
    "io"
    "fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hacker news's top stories")
}

func TopRated(w http.ResponseWriter, r *http.Request) {
    stories, err := TopStories()
    if err != nil {
        fmt.Fprintf(w, "Unable to get top stories %s", err)
        return
    }
    for i, story := range stories {
        fmt.Fprintf(w, "%d: %s \n", i+1, story.Title)
    }
}