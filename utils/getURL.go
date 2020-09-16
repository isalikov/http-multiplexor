package utils

import (
    "os"
    "strings"
)

func GetURL(query string) string {
    host := "https://api.themoviedb.org/3/search/movie?api_key={key}&query={query}"

    key := os.Getenv("KEY")

    if key == "" {
        key = "959c860179081d4921da15281aa06df9"
    }

    res := strings.Replace(host, "{key}", key, 1)
    res = strings.Replace(res, "{query}", query, 1)

    return res
}
