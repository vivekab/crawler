package firebase

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"github.com/vivekab/crawler/pkg/models"
)

func Test_store_GetLatest(t *testing.T) {
	ctx := context.Background()
	config := &firebase.Config{}
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEY_PATH"))
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.DatabaseWithURL(ctx, os.Getenv("FIREBASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	type args struct {
		source string
		status string
	}
	tests := []struct {
		name    string
		fields  store
		args    args
		wantM   *models.UrlData
		wantErr bool
	}{
		{
			name: "Success",
			fields: store{
				logger: logrus.New(),
				db:     client,
			},
			args: args{
				source: "dogecoinnewsnet",
				status: "LOADED",
			},
			wantM: &models.UrlData{
				Url: "x.com/dogecoin_news3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, err := tt.fields.GetLatest(tt.args.source, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.GetLatest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotM.Url != tt.wantM.Url {
				t.Errorf("store.GetLatest() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func Test_store_Insert(t *testing.T) {
	ctx := context.Background()
	config := &firebase.Config{}
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEY_PATH"))
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.DatabaseWithURL(ctx, os.Getenv("FIREBASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		source   string
		status   string
		newsData []models.UrlData
	}
	tests := []struct {
		name    string
		store   store
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "Succes",
			store: store{
				logger: logrus.New(),
				db:     client,
			},
			args: args{
				source: "dogecoinnewsnet",
				status: "LOADED",
				newsData: []models.UrlData{
					{
						Tags: []string{"dogecoin"},
						Url:  "x.com/dogecoin",
						Date: time.Date(2021, time.May, 28, 11, 42, 0, 0, time.UTC),
					},
				},
			},
			wantN: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotN, err := tt.store.Insert(tt.args.source, tt.args.status, tt.args.newsData)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("store.Insert() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
