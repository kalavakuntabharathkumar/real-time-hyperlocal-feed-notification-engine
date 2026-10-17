package main
import("context";"encoding/json";"log";"net/http";"os";"github.com/redis/go-redis/v9")
type Event struct{UserID string `json:"user_id"`;City string `json:"city"`;Text string `json:"text"`}
func main(){r:=redis.NewClient(&redis.Options{Addr:getenv("REDIS_ADDR","localhost:6379")});http.HandleFunc("/events",func(w http.ResponseWriter,q *http.Request){defer q.Body.Close();var e Event;if json.NewDecoder(q.Body).Decode(&e)!=nil{http.Error(w,"bad json",400);return};b,_:=json.Marshal(e);if r.XAdd(context.Background(),&redis.XAddArgs{Stream:"feed-events",Values:map[string]any{"payload":string(b)}}).Err()!=nil{http.Error(w,"redis error",500);return};w.WriteHeader(202)});log.Fatal(http.ListenAndServe(":8081",nil))}
func getenv(k,d string)string{if v:=os.Getenv(k);v!=""{return v};return d}
