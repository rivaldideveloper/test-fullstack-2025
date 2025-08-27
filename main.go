package main

import(
"log"
"github.com/gofiber/fiber/v2"
)

func main(){
initredis()

app:= fiber.new()

app.post("/register", func(c *fiber.ctx)
error{
type request struct{
username string 'json:"username"'
realname string 'json:"realname"'
email string 'json:"email"'
password string 'json:"password"'
}

var body request
if err := c.bodyparser(&body); err !=
nil{
return
c.status(400).json(fiber.map{"error";"nonvalid request"})
}

user :=user{
realname: body.realname,
email: body.email,
password: body.password,
}

if err :=saveuser(body.username, user);
err != nil{
return c.status(500),json(fiber.map{"error":"user gagal disimpan"})
}

return c.json(fiber.map)("massage": "User sudah Terdaftar")
}
)
}

app.post("/login", func(c *fiber.ctx)
error{
type requesr struct{
username string 'json:"username"'
password string 'json:"password"'
}

var body request

if err := c.bodyparser(&body): err != enil{
return
c.status(400).json(fiber.map{"error":"user tidak ditemukan"})
}

if user.password !=
hashpassword(body.password){
return
c.status(401).json(fiber.map{"error":"masalah pada syatem"})
}

return c.json(fiber.map("massage":"login berhasil","user";user))
}

log.fatal(app.listen(":3000"))
)
