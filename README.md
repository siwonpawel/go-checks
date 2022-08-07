# go-checks  

Validate common polish numbers.

Supported:  

- NIP (Numer Identyfikacji Podatkowej)
- REGON (Rejestr Gospodarki Narodowej)
- PESEL (Powszechny Elektroniczny System Ewidencji Ludności)

## How to run

> makefile is prepared to execute on podman / podman-compose if you're running docker please replace _podman-compose_ with _docker-compose_

Execute this command to build and start service 
```bash
make up
```


## How to validate data

```sh
curl --request POST \
  --url http://localhost:8080/validate \
  --header 'Content-Type: application/json' \
  --data '{
	"value": "96030700272"
}'
```

## Example response

```json
[
	{
		"country": "PL",
		"identifier": "NIP",
		"isValid": false,
		"Warnings": [
			"Missing country identifier at beginning."
		],
		"Error": "Invalid input length.",
		"CleanedInput": "96030700272"
	},
	{
		"country": "PL",
		"identifier": "PESEL",
		"isValid": true,
		"Warnings": [
			"Missing country identifier at beginning."
		],
		"Error": "",
		"CleanedInput": "96030700272"
	},
	{
		"country": "PL",
		"identifier": "REGON",
		"isValid": false,
		"Warnings": [
			"Missing country identifier at beginning."
		],
		"Error": "Invalid input length.",
		"CleanedInput": "96030700272"
	}
]
```