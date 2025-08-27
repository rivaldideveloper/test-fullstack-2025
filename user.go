package main

import(
"crypto/shal"
"encoding/hex"
"encoding/json"
"fmt"
)

type user struct{
realname string 'json:"realname"'
email string 'json:"email"'
password string 'json:"password"'
}

func hashpassword(password string) string
{
h := shal.new()
h.write([]byte(password))
return hex.encodetostring(h.sum(nil))
}

func saveuser(username string, user user)
error{
key := fmt.sprintf("login_%s", username)

jsonData, err := json.Marshal(user)
if err != nil {
return err
}

return rdb.set(ctx,key, jsonData, 0).Err()
}

func getuser(username string) (*user, error){
key := fmt.sprintf("login_%s", username)

val,err := rdb.get(ctx, key).result()
if err != nil{
return nil, err
}

var user user
err = json.unmarshal([]byte(val), &user)
if err != nil{
return nil,err
}

return nil, err
}
