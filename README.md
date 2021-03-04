<h1 align="center">Asteroid</h1> 
<p align="center">
    <img src="./.github/logo.png" width="20%" />
</p>
<h2 align="center">An Intuitive database built for rapid prototyping and ease of use!</h2>

## Design Goals 🎨

### Developer Experience

### Simplicity

## Documentation 📝

### Auth

#### Login

#### SignUp

### Create

#### Create a collection

In order to create a collection, run the command `asteroid create collection collectionName` into terminal, as a collection with `collectionName` will be created. This can only be done from the command line for now, but I'm hoping to allow devs to be able to do this through an API in the future.

#### Create a document

Send a POST request to `localhost:5555/api/collection/collectionName/create/document` where `collectionName` is the name of the collection the document should be created in.
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

#### Get Documents by User ID

### Update

#### Update element in collection

#### Update collection name

### Delete

#### Delete documents from collection

#### Delete Collection
