package CalendarAPI

import (
	"encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "net/url"
        "os"
        "os/user"
        "path/filepath"
        "time"

        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/calendar/v3"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
        cacheFile, err := tokenCacheFile()
        if err != nil {
                log.Fatalf("Unable to get path to cached credential file. %v", err)
        }
        tok, err := tokenFromFile(cacheFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(cacheFile, tok)
        }
        return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("calendar-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// readsFile reads in the credentials
func readFile() []byte{
        b, err := ioutil.ReadFile("clientSecretCalendar.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }
        return b
}

// parseFile Creates an oauth2.Config using the client_secret.json 
func parseFile(b []byte) *oauth2.Config{
        config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        return config
}

func CreateEventCalendar(loc, desc string, date time.Time){
        ctx := context.Background()
        b := readFile()
        config := parseFile(b)
        client := getClient(ctx, config)
        srv, err := calendar.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve calendar Client %v", err)
        }

        event := &calendar.Event{
                Summary: "CarLog Reminder",
                Location: loc, //taller
                Description: desc, //Solucion
                Start: &calendar.EventDateTime{
                        DateTime: date.Format(time.RFC3339),
                },
                End: &calendar.EventDateTime{
                       DateTime: date.AddDate(0,0,1).Format(time.RFC3339),
                },
        }

        calendarId := "primary"
        event, err = srv.Events.Insert(calendarId, event).Do()
        if err != nil {
                log.Fatalf("Unable to create event. %v\n", err)
        }
        fmt.Printf("Event created: %s\n", event.HtmlLink)
}


