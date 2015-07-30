Encounter Tracker API
=====================

The API that powers https://github.com/ateoto/encounter-tracker

Development
-----------

To run a local copy of the api with Vagrant: 

```bash
git clone https://github.com/ateoto/encounter-tracker-api.git && cd encounter-tracker-api
vagrant up
```

To update and refresh the api/migrations:

```bash
git pull && vagrant provision
```

