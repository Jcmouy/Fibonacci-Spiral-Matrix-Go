# FIBONACCI SPIRAL MATRIX GO

# OBJECTIVE

The objective of this project is to create a matrix, in which the numbers inside the matrix will be those of the Fibonacci sequence. For this, the user is asked as a starting point to insert the number of rows and columns, from this the matrix is created. The backend is built with Java Spring Boot, and in its frontend Angular, which is embedded in the application. The database is SQLite.

## TECHNOLOGIES
* Go Web Framework ([gin-gonic](https://github.com/gin-gonic/gin))
* Containerize ([docker](https://www.docker.com/))
* Swagger ([swaggo](https://github.com/swaggo/swag))
* Database
    * [SQLite](https://www.sqlite.org/index.html)
* Dependency Injection ([google wire](https://github.com/google/wire))
* Unit/Integration Tests ([testify](https://github.com/stretchr/testify))
* Tracing ([opentracing](https://github.com/opentracing/opentracing-go))
* Logger ([logrus](https://github.com/sirupsen/logrus))
* Error Wrapper ([pkg errors](https://github.com/pkg/errors))
* WebUI ([Angular 9](https://angular.io/))

## FOLDER STRUCTURE

### `cmd` (application run)
Main application executive folder.

### `internal` (application codes)
Private application and library code. 

### `test` (integration tests)
Application integration test folder.

### `web` (web ui)
Web application specific components: static web assets, server side templates and SPAs.

### `docs` (openapi docs)
Open api (swagger) docs files.

# MANUAL INSTALLATION

* The first thing you have to do is clone the repository:


    git clone https://github.com/Jcmouy/Fibonacci-Spiral-Matrix-Go.git

- To start the application the only thing that should be done is run makefile command.


    make docker-start

- This command builds all docker services, and then you should have the following endpoints.

# ENDPOINTS

Application URLS:

| Application | URL                                      | Purpose                                  |
|-------------|------------------------------------------|------------------------------------------|
| Angular UI  | http://localhost:5000                    | Fibonacci Spiral Matrix APP Project      |
| Swagger UI  | http://localhost:8080/swagger/index.html | Fibonacci Spiral Matrix API OpenAPI Docs |
| Jaeger UI   | http://localhost:16686                   | Opentracing Dashboard                    |

![2022-05-28 20_55_25-Swagger UI](https://user-images.githubusercontent.com/10815551/170846432-0e58c46c-5ee5-403f-8a72-87f60fd23290.png)

![2022-05-28 20_56_13-Swagger UI](https://user-images.githubusercontent.com/10815551/170846436-1e7d450f-149d-479d-a611-3108e796e1ed.png)

![2022-05-28 20_56_27-Swagger UI](https://user-images.githubusercontent.com/10815551/170846441-e77e7ec7-d347-4168-a20f-862f348ea483.png)

![2022-05-28 20_56_51-Swagger UI](https://user-images.githubusercontent.com/10815551/170846442-a805ffdd-8301-4ca0-815c-267f176dd6e1.png)

## Local Development
### Configuration
As mentioned before for the tracing we use Opentracing, and as tracer we use Jaeger which is a distributed tracing system that natively supports Opentracing. It has the following configuration:
  ```
    # JAEGER
    JAEGER_AGENT_HOST=localhost
    JAEGER_AGENT_PORT=6831
    JAEGER_SAMPLER_PARAM=1
    JAEGER_SAMPLER_TYPE=probabilistic
    JAEGER_SERVICE_NAME=fibonacci-spiral-matrix-go
    JAEGER_DISABLED=false
  ```  

### Dependency Injection

The project uses google wire for compile time dependency injection.
Docker compose files generates automatically **wire_gen.go** in containers but, it must be created manually for local development.

Wire dependency file is `/internal/wired/wire_gen.go`

    make wire

This command generates **wire_gen.go** with redis provider. When `wire_gen.go` file is checked, the following change will be seen.

  ```go
  // Injectors from wire.go:
  
func InitializeFiboSpiralMatrixHandler() (handler.FiboSpiralMatrixHandler, error) {
fiboSpiralMatrixService := provider.ProvideFiboSpiralMatrixService()
fiboSpiralMatrixHandler := provider3.ProvideFiboSpiralMatrixHandler(fiboSpiralMatrixService)
return fiboSpiralMatrixHandler, nil
}

func InitializeAuthHandler() (handler.AuthHandler, error) {
authRepository := provider2.ProvideAuthRepository()
authService := provider.ProvideAuthService(authRepository)
authHandler := provider3.ProvideAuthHandler(authService)
return authHandler, nil
}

// wire.go:

var AuthService = wire.NewSet(provider.ProvideAuthService, provider2.ProvideAuthRepository)

func InitializeHealthHandler() handler.HealthHandler {
return handler.HealthHandler{}
}
  ```

### Swagger

The command that generates the open api document to `/docs` folder.

    make swag

# CAPTURES

![welcome](https://user-images.githubusercontent.com/10815551/163720975-2c6ab92b-8baa-4bd4-8943-f0d0c4f85b3d.gif)

![register](https://user-images.githubusercontent.com/10815551/163720994-6fb204bf-1520-47a0-a38d-9f2e802e04bf.gif)

![login](https://user-images.githubusercontent.com/10815551/163721001-efe2d3b4-2ad6-4c44-9608-c84ccc8766c7.gif)

![spiral](https://user-images.githubusercontent.com/10815551/163813825-d6581da1-a755-4edd-8132-2777918d8657.gif)




