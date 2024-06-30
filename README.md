# Netflix_Clone

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 15.2.6.

## Front End

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.


## 🛠 Back End 
![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) 

This project uses fiber framework 
[Click here](https://docs.gofiber.io) for official documentation 
```bash
  go get github.com/gofiber/fiber/v2
```

All server-related configurations can be specified in the `.yaml` file located inside the `env` directory

Install MongoDB drivers
```bash
  go get go.mongodb.org/mongo-driver/mongo
```

## Documentation

The go server is configured to run on port 9800.The database connection files are in the connection.go file.Change the Database name and collection name accordingly

## Execution
```bash
  go run ./main.go
```
