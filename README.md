# Recipe Count Analyzer  

A small golang app to analyze recipes and their popularity. 
It also gives delivery stats for selected postcode and timeslot and performs a search among recipes by selected keyword

As a data uses the following json stucture 

```yaml
  {
    "postcode": "10224",
    "recipe": "Creamy Dill Chicken",
    "delivery": "Wednesday 1AM - 7PM"
  },
  {
    "postcode": "10208",
    "recipe": "Speedy Steak Fajitas",
    "delivery": "Thursday 7AM - 5PM"
  }]
```

Output has the following structure

```yaml
{
    "unique_recipe_count": 15,
    "count_per_recipe": [
        {
            "recipe": "Speedy Steak Fajitas",
            "count": 1
        },
        {
            "recipe": "Tex-Mex Tilapia",
            "count": 3
        },
        {
            "recipe": "Mediterranean Baked Veggies",
            "count": 1
        }
    ],
    "busiest_postcode": {
        "postcode": "10120",
        "delivery_count": 1000
    },
    "count_per_postcode_and_time": {
        "postcode": "10120",
        "from": "11AM",
        "to": "3PM",
        "delivery_count": 500
    },
    "match_by_name": [
        "mushroom recipe", "veggie recipe"
    ]
}
```

## Build and run

1. In cli
```yaml
  cd /RecipeCountRnalyzer
  
  go build github.com/abondar24/RecipeCountAnalyzer
  
  ./RecipeCountAnalyzer --file=<path-to-json> --recipe=<keyword-to-find-matches> --postcode=<postcode-find> --time="{h}AM - {h}PM"
```
Please note that --file flag is mandatory and for the other flags default values are set up.
```yaml
  --recipe=veggie
  --postcode=10120
  --time=10AM - 2PM
```
2. As a docker container
```yaml
docker build -t <tag-name> .


docker run -it -v <path-to-json>:/json <tag> --file /json/<json-name>.json <other-flags>
```
