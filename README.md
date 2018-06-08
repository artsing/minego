# minego
#### get users
>GET localhost:9000/users

#### create user
>POST localhost:9000/users/id

Response Body is a Json String
```
{name:"Bob", email:"bob@yeah.net", website:"bob.me"}
```

#### edit user
>PATCH localhost:9000/users/id

Response Body is a Json String
```
{name:"Bob", email:"bob@yeah.net", website:"bob.me"}
```

#### delete user
>DELETE localhost:9000/users/id