# Simple-Agenda

Simple-Agenda is an API using a go idiomatic way as desing patern, this API is used to save important appoinments.


## Installation

We are using [go mod](https://blog.golang.org/using-go-modules) to manage all dependencies on this project. So you are free to run:

```bash
foo@bar:~$ go build 
```

And all dependencies will be resolved automatically, however you could install dependencies manually as following:
```bash
foo@bar:~$ go mod tidy 
```

## Usage

Before start make sure to edit `.env` file with our configuration, and the database contains the `appointments` table (you could use `.sql` as base to create table in our database).

### Endpoints

| HTTP METHOD | ENDPOINT            | DESCRIPTION                           |
| ----------- | ------------------- | ------------------------------------- |
| POST        | `/appointments`     | Create a new Appointment              |
| GET         | `/appointments/:id` | Get especified Appointment by ID      |
| GET         | `/appointments`     | Get all Appointments using pagination |
| PUT         | `/appointments/:id` | Updates an Appointment by ID          |
| DELETE      | `/appointments/:id` | Deletes an Appointment by ID          |

### Request Structure

#### For POST endpoint the request is:

##### Headers:

- **Content-Type** : `application/json`

##### Body:

The body is a json formated, with the following params:

- **name**: Represent the name of this appointment (required, string);
- **date**: Represent the date when is this appointment as DD/MM/YYYY format (required, string);
- **hour**: Represent the hour when is this appointment as HH:MM format (optional, string);
- **local**: Represent the local where is this appointment (optional, string);

#### For PUT endpoint the request is:

##### Headers:

- **Content-Type** : `application/json`

##### Body:

The body is a json formated, with the following params:

- **name**: Represent the name of this appointment (required, string);
- **date**: Represent the date when is this appointment as DD/MM/YYYY format (required, string);
- **hour**: Represent the hour when is this appointment as HH:MM format (optional, string);
- **local**: Represent the local where is this appointment (optional, string);

##### Param

The param is an path param (eg: `/foo/:bar`), accepts an ID of an appointment. 

#### For GET endpoints the requests are:

**For a simple appointment:**

##### Param

The param is an path param (eg: `/foo/:bar`), accepts an ID of an appointment. 

**For multiple appointments:**

##### Query

The query is a url query (eg: `/?foo=bar`), who accepts the following params:

- **page**: represent an automatic pagination (optional, int);
- **offset**: represent the number of items to skip before select (optional, int, 0);
- **limit**: represent the number of items to recive (optional, int, 10);

When no param is passed the default is assumed.

#### For DELETE endpoint the request is:

##### Param

The param is an path param (eg: `/foo/:bar`), who accepts an ID of an appointment. 

## Exemples

### Create an appointment:
```bash
curl --request POST \
  --url http://localhost:3000/api/appointments \
  --header 'content-type: application/json' \
  --data '{
	"name": "ea7ae120-bbb8-49e3-bde4-1736d7c055b9",
	"date": "17/10/2020"
}'

# Response
{
  "error": null,
  "message": "created",
  "result": {
    "id": 1,
    "name": "ea7ae120-bbb8-49e3-bde4-1736d7c055b9",
    "date": "17/10/2020"
  }
}
```

### Read an appointment:
```bash
curl --request GET \
  --url http://localhost:3000/api/appointments/1

# Response
{
  "id": 1,
  "name": "ea7ae120-bbb8-49e3-bde4-1736d7c055b9",
  "date": "17/10/2020"
}
```

### Read multiple appointments:
```bash
curl --request GET \
  --url 'http://localhost:3000/api/appointments?page=0'

# Response
{
  "_metadata": {
    "limit": 10,
    "next": 1,
    "offset": 0,
    "previous": 0,
    "this": 0
  },
  "result": [
    {
      "id": 1,
      "name": "ea7ae120-bbb8-49e3-bde4-1736d7c055b9",
      "date": "17/10/2020"
    }
  ]
}
```

### Update an appointment:
```bash
curl --request PUT \
  --url http://localhost:3000/api/appointments/1 \
  --header 'content-type: application/json' \
  --data '{
	"name": "cool name",
	"date": "15/02/2017",
	"hour": "12:00",
	"local": "somewhere"
}'

# Response
{
  "error": null,
  "message": "updated",
  "result": {
    "id": 1,
    "name": "cool name",
    "date": "15/02/2017",
    "hour": "12:00",
    "local": "somewhere"
  }
}
```

### Delete an appointment:
```bash
curl --request DELETE \
  --url http://localhost:3000/api/appointments/1

# Response 
# No content if has no errors

# Testing
curl --request GET \
  --url http://localhost:3000/api/appointments/1

{
  "error": "record not found",
  "message": "not found data with the given id"
}
```


## License
[MIT](https://choosealicense.com/licenses/mit/)