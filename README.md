 Cartesian API
=============

Create an API server in [go](https://golang.org/). It will deal with a series of points represented as (x,y) coordinates on a simple 2-dimensional plane. Take a look at https://en.wikipedia.org/wiki/Cartesian_coordinate_system if you need a refresher on this concept.

It must have an api route at `/api/points` that accepts a `GET` request with the following parameters, and returns a JSON list of points that are within `distance` from `x,y`, using the Manhattan distance method. The points should be returned in order of increasing distance from the search origin.
- `x` integer (required). This represents the `x` coordinate of the search origin.
- `y` integer (required). This represents the `y` coordinate of the search origin.
- `distance` integer (required). This represents the Manhattan distance; points within `distance` from `x` and `y` are returned, points outside are filtered out.

The Manhattan distance is measured "block-wise", as the distance in blocks between any two points in the plane (e.g. 2 blocks down and 3 blocks over for a total of 5 blocks). It is defined as the sum of the horizontal and vertical distances between points on a grid. Formally, where `p1 = (x1, y1)` and `p2 = (x2, y2)`, `distance(p1,p2) = |x1-x2| + |y1-y2|`.

On startup, the API server should read a list of points from `data/points.json`.


### How to run:

Local: At the root of project execute:

```
make run
```

Using docker:

1 - Create docker image. Type de command:

```
docker build -f Dockerfile -t civi/cartesian-api:v1.0 --build-arg PORT=9001 .
```

2 - Run de container

```
docker run -p 9001:9001 --env-file=config/.env -d civi/cartesian-api:v1.0
```

Using the curl command call de api

```
curl -X GET -H'Content-Type:application/json' 'http://localhost:9001/api/points?x=10&y=10&distance=50'
```

Important: 

Don't put secrets ou security params into .env files and push them to git. 
This project there aren't secrets or any sensitive data at .env file