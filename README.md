# Accounting-notebook
To launch the project use: "go run main.go"

## Endpoints

GET    v1/capital/:id       Get the actual capital for de user      
POST   v1/credit/:id        Increase money into an account      
POST   v1/debit/:id         Decrease money from an account

## Bodies
Both post methods use the same body for the test. The UserID is hardcode.

{
	"UserID":1234,
	"Amount":12.3
}

## Examples

GET localhost:8080/v1/capital/1234 First call should get 0

POST localhost:8080/v1/credit/1234
{
	"UserID":1234,
	"Amount":12.3
}
Should return:
{
  "UserID": 1234,
  "PersonalInfo": {
    "FName": "Homero",
    "LName": "Simpson",
    "Birth": "",
    "DocumentType": "",
    "DocumentNumber": "",
    "Address": {
      "Street": "",
      "Number": "",
      "City": "",
      "State": "",
      "Country": ""
    }
  },
  "Capital": 12.3
}

POST localhost:8080/v1/debit/1234
{
	"UserID":1234,
	"Amount":12.3
}
Could response:
{
  "Message": "Not enough credit"
}
or the diference.
