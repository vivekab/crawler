## News Format
```json
[
    {
        "source":"x.com",
        "tags":[],
        "url":"x.com/dogecoin_news1",
        "news_time":"11:59pm",
        "news_status":"LOADED"
    },
    {
        "source":"x.com",
        "tags":[],
        "url":"x.com/dogecoin_news2",
        "news_time":"11:45pm",
        "news_status":"PROCESSED"
    },
    {
        "source":"x.com",
        "tags":["dogecoin"],
        "url":"x.com/dogecoin_news3",
        "news_time":"11:29pm",
        "news_status":"UPDATED"
    },
]
```


## Tags Format
Contains recent news for a given tag for the past 3 days
```json
[
    {
        "name":"dogecoin"
    },
]
```


## Solution:
### Notification of the recent news obtained from the data sources.
Cron runs and posts that are "PROCESSED" prepares a notification with the format of following<br>
<br>
Source news:
```json
{
    "x.com":[
        {},
        {}
    ]
}
```
Tag News:
```json
{
    "dogecoin":[],
    "shiba":[],
    "bitcoin":[]
}
```


### Get the latest for the news for the new user
Api with the endpoint `/latest-news` gives out the news for the past three days.
<br>

Response:
```json
{
    "data":[{
        "source":"x.com",
        "tags":[],
        "url":"x.com/dogecoin_news1",
        "news_time":"11:59pm",
        "news_status":"LOADED"
    },
    {
        "source":"x.com",
        "tags":[],
        "url":"x.com/dogecoin_news2",
        "news_time":"11:45pm",
        "news_status":"PROCESSED"
    },
    {
        "source":"x.com",
        "tags":["dogecoin"],
        "url":"x.com/dogecoin_news3",
        "news_time":"11:29pm",
        "news_status":"UPDATED"
    }]
}
```

## Questions
