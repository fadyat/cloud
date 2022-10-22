## MongoDB tutorial

- Enter in container

```bash
mongo mongodb://localhost:27017
```

- Show databases

```bash
show dbs;
```

- Use database (create or switch to database)

```bash
use mydb;
```

- Create collection

```bash
db.createCollection("help");
db.createCollection("help", { capped : true, autoIndexId : true, size : 6142800, max : 10000 } );
```

- Drop

```bash
db.dropDatabase(); # Drop current database
db.hello.drop(); # Drop collection
```

- Help

```bash
db.help();
```

- Show collections

```bash
show collections;
```

- Show stats

```bash
db.stats(); # Show stats for current database
db.hello.stats(); # Show stats for collection hello
```

- Insert

```bash
db.hello.insert({name: "John", age: 30, status: "A"});
db.hello.insertMany([{name: "John", age: 30, status: "A"}, {name: "Peter", age: 40, status: "A"}]);
```

- Find

```bash
db.hello.find(); # Find all
db.hello.find({name: "John"}); # Find all documents where name is John
db.hello.find({name: "John", age: 30}).pretty(); # Pretty print
db.hello.find({name: "John", age: 30}).limit(1); # Limit to 1
db.hello.find({name: "John"}, {age: 1}); # include only age field
db.hello.find({name: "John"}, {age: 0}); # exclude age field
```

- Update

```bash
db.hello.update({name: "John"}, {name: "John", age: 40, status: "A"});
db.hello.update({name: "John"}, {$set: {age: 40}});
db.hello.update({name: "John"}, {$unset: {age: 1}}); # Remove field age where name is John
db.hello.update({name: "John"}, {$inc: {age: 1}}); # Increment age by 1 where name is John
db.hello.update({name: "John"}, {$pull: {hobbies: "Sports"}}); # Remove Sports from hobbies array where name is John
db.hello.update({name: "John"}, {$push: {hobbies: "Sports"}}); # Add Sports to hobbies array where name is John
```

- Delete

```bash
db.hello.remove({name: "John"});
db.hello.remove({name: "John"}, 1); # Remove only one document
```