# PhishQL

I got bored and did this. I have no shame. Yes it is extremely over-engineered. But when the time comes for an enterprise level, distributed, Phish setlist data solution with 100% code coverage we'll see who's laughing.

Some would ask "Why use gRPC for this?", I ask "Why not?".

This project is still very young so expect breaking changes. LOTS OF THEM.

## Installation

**Clone the repository**

This assumes that you have [git installed](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

```
git clone git@github.com:jloom6/phishql.git
```

**Bootstrap**

This assumes that you have [golang](https://golang.org/doc/install) and [dep](https://golang.github.io/dep/docs/installation.html) installed.

```
cd phishql
make bootstrap
```

**Shock and persuade my soul to ignite**

This assumes that [you get your ass handed to you everyday](https://www.youtube.com/watch?v=9PinOWOAtHk).

```
make test
```

## Run like an antelope

**Run everything in Docker**

This assumes that you have [Docker installed](https://docs.docker.com/install/).

```
make run-hard
```

## Set the gearshift for the high gear of your soul

**Call the REST API**

This assumes that you have [jq installed](https://stedolan.github.io/jq/download/). You obviously don't need it but it makes everything look nice.

```
curl -XPOST -d '{}' $(docker-machine ip):8080/v1/shows | jq .
```

Congrats. [You did it!](https://www.youtube.com/watch?v=wxEAyJfIUI4)

## The API

Swagger json can be found in [here](https://github.com/jloom6/phishql/blob/master/proto/jloom6/phishql/phishql.swagger.json), just paste that into [this swagger editor](https://editor.swagger.io/) to see example HTTP requests. You can also use the [proto file](https://github.com/jloom6/phishql/blob/master/proto/jloom6/phishql/phishql.proto) and give [gRPC](https://grpc.io/) a try on port :9090!

**Available Endpoints**

I'll be more descriptive of the endpoints eventually, the swagger docs should be sufficient for now.

|gRPC|HTTP|
|---|---|
|GetShows|/v1/shows|
|GetArtists|/v1/arists|
|GetSongs|/v1/songs|
|GetTags|/v1/tags|
|GetTours|/v1/tours|
|GetVenues|/v1/venues|

**Base Conditions**

You can search for shows with basic conditions like this

```
curl -XPOST -d '{
    "condition": {
        "base": {
            "year": 2019,
            "month": 7,
            "day": 14,
            "day_of_week": 1,
            "city": "East Troy",
            "state": "WI",
            "country": "USA",
            "song": "Ruby Waves"
        }
    }
}' $(docker-machine ip):8080/v1/shows | jq .
```

The fields in the base condition are all "anded" together. If you leave the fields out of the request they are ignored in the query. The "day_of_week" field is indexed such that Sunday is 1, Monday is 2, ..., Saturday is 7.

**Composable Conditions**

You can compose the conditions using "and" and "or" as demonstrated below. The query is for shows that occurred in the state of NJ AND occurred on a Sunday OR Saturday

```
curl -d '{
    "condition": {
        "and": {
            "conditions": [
                {
                    "base": {
                        "state": "NJ"
                    }
                },
                {
                    "or": {
                        "conditions": [
                            {
                                "base": {
                                    "day_of_week": 1
                                }
                            },
                            {
                                "base": {
                                    "day_of_week": 7
                                }
                            }
                        ]
                    }
                }
            ]
        }
    }
}' $(docker-machine ip):8080/v1/shows | jq .
```

If the entity model seems a bit... superfluous, that is because they are auto-generated from proto files using [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway).

Your hands and feet are mangoes, you're gonna be a genius anyway.

## Contact
- Email: John Loomis - [jloom6@gmail.com](mailto:jloom6@gmail.com)