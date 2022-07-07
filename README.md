# gonerator
A simple Go microservice boilerplate. Swagger included.

## Usage
Probably you have already done but it's good to remind one thing to you before diving into boilerplate generation. You should have already installed Go and had $GOPATH on your environment.

After the assumption made above, we are good to go creating our first microservice boilerplate.

1. After clonning the repository, it's best to land on the main directory. 
```
$ cd generator-gonerator
```

2. Install all dependencies and link
```
$ npm install
$ npm link
```

3. Run the boilerplate generator.
```
$ yo gonerator
```

After you applied the steps above, go to your project directory and run below command so as to install dependencies and run your mini API.
```
$ make
```