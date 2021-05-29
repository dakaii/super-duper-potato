### README

## how to run this application

You first need to build the application with the command below.
``` bash
make build
```

Subsequently, you need to feed the mock data to the database.
``` bash
make migrate
```

Then, you can start up the server with this command.
``` bash
make up
```

## example query
``` gql
query {
  scooter (lat: 1.39, lng: 103.9, limit:20, distance: 5000) {
      title
      lat
      lng
    }
}
```
