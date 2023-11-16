# playground

Web Service will return a random name and a random Chuck Norris Nerdy Joke, using the retrieved Random Name instead of "Chuck Norris".

### Example:

```
$ curl "http://localhost:5002"
Hasina Tanweer’s OSI network model has only one layer - Physical..
```

## PreRequisits
Needs Go Installed

---
## First Time Set UP

Run command

``` 
bin/run.sh
```

Will start the server and open the browser or can run the following curl after starting the browser:

``` 
curl "http://localhost:5002"
```
---
### To be done
Items required:
- DockerFile creation with docker-compose; and ci/cd pipeline for environment deployment
- config/APP new App() creation to pass in variables
- Start Service using config.App()

Improvements: 
- Adding interfaces so can create mock testing on APIs

Notes:
- Branch `main` took 3 hours.
- Please see branch `simple-service` (changes in branch add 1 hour) for changes
  - include adding golangci-lint, some unit tests, simplfy service, updated readme and instructions.


---
Copywrite &copy; 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>
