package server

import (
    "encoding/json"
    "fmt"
    "http-multiplexor/utils"
    "net/http"
)

type Movie struct {
     Id int64 `json:"id"`
     Title string `json:"title"`
}

type RequestBody struct {
    Queries []string `json:"queries"`
}

type TMDiResponse struct {
    Results []Movie `json:"results"`
}

type ResponsePayload struct {
    Data []Movie `json:"data"`
    Total int64 `json:"total"`
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var body RequestBody
        var response_p ResponsePayload

        err := json.NewDecoder(r.Body).Decode(&body)

        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        for _, q := range body.Queries {
            var tmdbi_r TMDiResponse

            url := utils.GetURL(q)

            resp, _ := http.Get(url)
            json.NewDecoder(resp.Body).Decode(&tmdbi_r)

            for _, movie := range tmdbi_r.Results {
                is_exist := false

                for _, m := range response_p.Data {
                    if m.Id == movie.Id {
                        is_exist = true
                        break
                    }
                }

                if !is_exist {
                    response_p.Data = append(response_p.Data, movie)
                    response_p.Total = response_p.Total + 1
                }
            }
        }

        j, _ := json.Marshal(response_p)

        w.Header().Set("Content-Type", "application/json")
        w.Write(j)
    } else {
        w.WriteHeader(http.StatusNotFound)

        fmt.Fprintf(w, "Not Found")
    }
}
