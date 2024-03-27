# HOUSES

simple api dedicated to house information


## Routes

/properties - allows GET method and returns list of properties limited by 10 

/properties/{id} - allows only GET method and return a property of {id} see [GET method](#get-method)


/properties - allows POST method see [POST method](#post-method)



## GET method

```bash
curl -u admin:admin localhost:8080/properties
```

```bash
curl -u admin:admin localhost:8080/properties/3
```

## POST method

### properties

property_type_id (**required**) - refference to a property ("flat", "land" .etc) 

address_id (**required**) - refference to an address from addresses table (unique address only)  

price (**required**) - price of the property

rooms - amount of rooms 

area - area of the property

description - description of the property

Without the **requred** fields the server throws an error of bad request

## Example

```bash
curl -H "Content-Type: application/json" \ 
-X POST \
--data '{ "property_type_id": 1, 
"address_id": 1, 
"price": 1200000, 
"rooms": 1, 
"area": 45.5,  
"description": "Студия в новом доме с ремонтом и мебелью" }' \
-u admin:admin localhost:8080/properties
```
