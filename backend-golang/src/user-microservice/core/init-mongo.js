db.createUser(
    {
      user: "vosuyak",
      pwd: "password",
      roles: [ { role: "readWrite", db: "victorosuyak" },
               { role: "dbOwner", db: "victorosuyak" } ]
    }
  )