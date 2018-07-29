# Rebin
> Rebuilt of the concept of [fiche](https://github.com/solusipse/fiche) running as termbin.com.

# Goals
#### 1.0.0
The goal of the project for its first version is that the service can receive data via tcp and serve that data over http.



# How it should work
##### tcp
- `nc *.*.*.* 1337` -> id of post

##### http
###### routes:
- GET / -> description website
- GET /{slug} -> post
- GET /{=! slug} -> 404
- *GET /{slug}/stats* -> return [timestamp, *views*, *size*]

###### POST
1. user sends post via tcp
    - `echo test data | nc server 1337`
2. server reveives data on port 1337
3. server generates id and timestamp for post
3. server saves data to database
    - {post, gen_id, timestamp}
4. server returns url with gen_id to user
    - `http://server/gen_id`

###### GET
1. user sends http GET request with gen_id
    - `curl http://server/gen_id`
2. server checks if gen_id is in database -> else return 404
3. server looks up gen_id in database
4. server returns post
