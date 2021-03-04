<h1 align="center">Asteroid</h1> 
<p align="center">
    <img src="./.github/logo.png" width="20%" />
</p>
<h2 align="center">An Intuitive database built for rapid prototyping and ease of use!</h2>

## Installation 💻

1. Clone the repo - `git clone https://github.com/pranavnt/Asteroid && cd Asteroid`
2. Install dependencies - `go get`
3. Build asteroid - `go build asteroid.go`

## Documentation 📝

### Auth

#### SignUp

When a user signs up, send a POST request to `/signUp` with the following JSON:

```json
{
  "username": "username",
  "password": "password"
}
```

Your username and password will be saved (password will be encrypted), and you will be returned with your userid!

#### Login

When a user trys to login, send a POST request to `/login` with the following JSON:

```json
{
  "username": "username",
  "password": "password"
}
```

If it is successful, you will be returned with your userid, and if not, you'll be returned with a message saying that your password was incorrect or you haven't signed up.

### Create

#### Create a collection

In order to create a collection, run the command `asteroid create collection collectionName` into terminal, as a collection with `collectionName` will be created. This can only be done from the command line for now, but I'm hoping to allow devs to be able to do this through an API in the future.

#### Create a document

Send a POST request to `/api/collection/{collectionName}/create/document` where `collectionName` is the name of the collection the document should be created in.
You should post the json you want to be stored in that document, with the key of `uid` being the user id. Here is sample JSON that could be posted:

```json
{
  "uid": "userid",
  "data1": "hello",
  "key2": "world"
}
```

### Read

#### Get Document by ID

To get a document by its id, send a GET request to `/api/collection/{collectionName}/document/read/{_id}`, where `_id` is the document id and `collectionName` is the name of the collection.

#### Get Documents by User ID

To get all documents from a collection from a given user, send a GET request to `/api/collection/{collectionName}/document/read/user/{uid}`, where `uid` is the user id, and `collectionName` is the name of the collection.

#### Get All Documents from collection

To get all documents from a collection, send a GET request to `/api/collection/{collectionName}/all`, where `collectionName` is the name of the collection.

### Update

#### Update element in collection

To update an element in a collection, send a POST request to `/api/collection/todos/document/update/{_id}`, where `_id` is the document id. You need to post the new JSON that should be stored in that document along with the user id. Here is sample JSON that can be posted:

```json
{
  "uid": "userid",
  "data1": "world",
  "key2": "hello"
}
```

#### Update collection name

To update the name of a collection, run the command `asteroid update collection oldName newName`, where `oldName` is the name of the collection, and `newName` is the new name!

### Delete

#### Delete document from collection

To delete a document from a collection, send a GET request to `/api/collection/todos/document/delete/{uid}/{_id}`. Here, `uid` will be the userid and `_id` will be the document id.

#### Delete Collection

To delete a collection, run the command `asteroid delete collection collectionName` into terminal, and the collection with the name `collectionName` will be deleted. You can only do this through the CLI right now, but I'm hoping to make this accessible via an API in the future.

## Design Goals 🎨

### Developer Experience

d

### Simplicity

When building full stack apps, I often get tired of building an API, which interacts with a database for really simple operations.
