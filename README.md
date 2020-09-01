# Contacts (Golang)
A simple REST API for studies made in *Golang* simulating a basic contact crud.

## Structure

```
Contact {
  ID : string
  Firstname : string
  Lastname : string
  Address{
    City : string
    State : string
  }
}
```

## Endpoints

- (GET) contacts

- (GET) contacts/:id
  - (id) The searched person's id
  
  
- (POST) contacts
  - data : *Contact*


- (DELETE) contacts/:id
  - (id) The id of the person to be deleted
