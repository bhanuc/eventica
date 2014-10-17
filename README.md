# README #

This is a soon to be OpenSource event Management Portal made in GO

### To get things running ###

```
sudo apt-get install libsasl2-dev bzr build-essential mercurial git-core
Ensure mongodb is running , check using mongo command

git clone https://github.com/bhanuc/eventica.git

& then move the folder in the repo to $GOPATH/src/github.com/eventica

go get ...

cd $GOPATH/src/github.com/eventica/eventica

go run main.go
```

Your Application would be running on localhost:8000

### Contribution guidelines ###

* Writing tests (Currently test coverage is preety poor)
* Fixing Issue - (facing an issue, feel free to report or submit a pull request)
* Adding a feature - (want to expand features and make it modular, open a issue )

### Creator ###

* Bhanu P Chaudhary
* Other community or team contact
