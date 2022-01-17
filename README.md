## Boilerplate API

Boilerplate API template includes all the common packages and setup used for API development in this Company.

### Development

- Copy `.env.example` to `.env` and update according to requirement.
- Create `serviceAccountKey.json` file for firebase admin sdk.
- To run `docker-compose up` (with default configuration will run at `5000` and adminer runs at `5001`)

#### Run Boilerplate CLI 🖥

- Run `docker-compose exec web sh`
- After running type `./__debug_bin cli` you will start cli application.
- Choose the commands to run afterwards.

#### Migration Commands 🛳

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |
| `make crud`         | Create crud template                           |

### Implemented Feature

- Dependency Injection (go-fx)
- Routing (gin web framework)
- Environment Files
- Logging (file saving on production) zap
- Middlewares (cors)
- Middleware transaction
- Database Setup (mysql)
- Models Setup and Automigrate (gorm)
- Repositories
- Implementing Basic CRUD Operation
- CRUD Scaffold Generator
- Migration Runner Implementation
- Live code refresh
- Firebase basic setup.
- GCP Cloud Storage bucket setup.
- Twilio Basic setup.
- Gmail Api Setup.

**For Debugging 🐞** Debugger runs at `5002`. Vs code configuration is at `.vscode/launch.json` which will attach debugger to remote application.
