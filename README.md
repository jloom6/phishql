# phishql

I got bored and did this. I have no shame. Yes it is extremely over engineered. But when the time comes for an enterprise ready, distributed, Phish setlist data solution with 100% code coverage we'll see who's laughing.

Some would ask "Why use gRPC for this?", I ask "Why not?".

Right now the actual query API is pretty bare bones, but I was more interested in setting up the infrastructure of this this. I'm new to Docker so cut me some slack if there are very noobish things going on here. I'll continue to add functionality as I find the time.

## Installation

**Clone the repository**

This assumes that you have [git installed](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

```
git clone git@github.com:jloom6/phishql.git
```

**Bootstrap**

This assumes that you have [golang installed](https://golang.org/doc/install).

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
curl $(docker-machine ip):8080/v1/shows | jq .
```

Congrats. [You did it!](https://www.youtube.com/watch?v=wxEAyJfIUI4)

## Contact
- Email: John Loomis - [jloom6@gmail.com](mailto:jloom6@gmail.com)