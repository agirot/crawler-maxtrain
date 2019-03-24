# Crawler-MaxTrain !

Personal tool to crawl TG\*V M@X API and send SMS (FREE mobile api) if schedule match.
Alert are stored in memory (TG*V M@X API allow a short date-range)

#### Edit config file to check a weekday with schedule

```
[
  {
   "from": "FRPAR",
    "to": "FRANE",
    "day": "Friday",
    "start_watch": "17:30",
    "end_watch": "23:00"
  },
  {
    "from": "FRANE",
    "to": "FRPAR",
    "day": "Monday",
    "start_watch": "7:00",
    "end_watch": "9:30"
  }
]
```

#### Command line args

```
  -config string
    	config file path (default in bin location) (default "config.json")
  -sms_key string (free mobile api)
    	List days to send alert
  -sms_user string (free mobile api)
    	List days to send alert
``