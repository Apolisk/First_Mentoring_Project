# Passgen

**passgen** is my first ever pet project as a result of @demget's mentoring. It allows a user to generate a set of unique passwords with the given configuration.

## Usage
### Build docker image 
```dockerfile
docker build -t passgen:1.0 .
```
### Run docker container
```dockerfile
docker run IMAGE_ID passgen -c -n -l -s -p -o
```
### Parameters

| Flag | Description                                                                        |
|------|------------------------------------------------------------------------------------|
| -c   | Specify the numbers of passwords to generate  ```By default 1```                   |
| -n   | Specify the desired password length ```By default 10```                            |
| -l   | **If specified** password will consists of Upper && Lower Letters ```a-z A-Z```    |
| -s   | **If specified** password will consists of special symbols  ```!@#$%^&*()/?{}[]``` |                   |
| -p   | Specify the number of threads to use ```By default  1```                           |
| -o   | Specify where output file will be stored ```By default current directory```        |

