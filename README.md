# Tiny Url

## Run
### Docker
You can run the project with docker compose.
The following command will launch the database and the API on your local machine
```shell
make start-docker
```

## Usage
### Create short URL
Run the following POST request
```bash
curl http://localhost:3000/create -X POST -d '{"url": "<YOUR_URL>"}
```

It will return the shortened url in the following format:
```json
{
  "url": "http://localhost:3000/aBcDeFgH"
}
```

> [!NOTE]
> If the URL is badly formatted, the API will return an error
> ```json
> { "error": "not a valid url" }
> ```

### Access shortened url
Simply go to the url send on creation.
The format is `http://<host>/<slug>`.

> [!NOTE]
> Every access to a slug will increase the `clicked` counter

If the slug does not exist, the API will return this message:
```json
{
  "error": "not found",
  "message": "You can create your own tiny url with a POST request to http://localhost:3000/create with the field 'url'"
}
```

If a user tries to access to a shortened url which expiration date is in the past, the API will also return not found

### Extend slug expiration
Run the following POST request:
```bash
curl http://localhost:3000/extend -X POST -d '{"slug": "<YOUR_SLUG>"}'
```

The API will return the new expiration date

> [!NOTE]
> If the slug does not exist, it will return an error

### Get information about a slug 
Run the following POST request:
```bash
curl http://localhost:3000/info -X POST -d '{"slug": "<YOUR_SLUG>"}'
```