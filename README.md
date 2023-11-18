# playground

Web Service will return a random name and a random Chuck Norris Nerdy Joke, using the retrieved Random Name instead of "Chuck Norris".

### Example:

```
$ curl "http://localhost:5002"
Hasina Tanweer’s OSI network model has only one layer - Physical..
```

## PreRequisits
Needs Go Installed
- Clone github.com/RebGov/playground (branch is main)


---
## First Time Set UP

Run command

``` 
bin/run.sh
```

Will start the service. Please either run following curl comand or open `http://localhost:5002` in browser.

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



---
Copywrite &copy; 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>
