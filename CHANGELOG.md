# Changelog

All notable changes to this project will be documented in this file.

## [unreleased]

### 🚀 Features

- *(currency-wallet)* Add reading environment variables from the .env file
- *(currency-wallet)* Add entry point
- *(currency-wallet)* Add logger initialization
- *(currency-wallet)* Add database configuration
- *(currency-wallet)* Add the database configuration to the .env file
- *(currency-wallet)* Add storage package
- *(currency-wallet)* Add connection to postgres database
- *(currency-wallet)* Add database initialization
- *(currency-wallet)* Add docker-compose.yml
- *(currency-wallet)* Add migrations
- *(currency-wallet)* Add sl package for structured error logging with slog
- *(currency-wallet)* Add improved error logging
- *(currency-wallet)* Add model package
- *(currency-wallet)* Add handler initialization at server startup
- *(currency-wallet)* Add repository init
- *(currency-wallet)* Add service
- *(currency-wallet)* Add handler
- *(currency-wallet)* Add a user initialization handler
- *(currency-wallet)* Add http-handler
- *(currency-wallet)* Add a user balance handler
- *(currency-wallet)* Add a request handler to exchange
- *(currency-wallet)* Add a handler for user's wallet transactions
- *(currency-wallet)* Add handler functions for endpoints
- *(currency-wallet)* Add auth, balance, exchange, wallet packages to the repository package
- *(currency-wallet)* Add repository package
- *(currency-wallet)* Add postgres  package
- *(currency-wallet)* Add service  package
- *(currency-wallet)* Add handler dependency on service
- *(currency-wallet)* Update init
- *(currency-wallet)* Add CreateUser method for AuthPostgres
- *(currency-wallet)* Add AuthRepository interface
- *(currency-wallet)* Add password hashing in CreateUser service
- *(currency-wallet)* Add repository argument to NewService constructor
- *(currency-wallet)* Update Repository constructor
- *(currency-wallet)* Add handler error handling function
- *(currency-wallet)* Add a method for user registration and update the constructor
- *(currency-wallet)* Add arguments for AuthHandler initialization
- *(currency-wallet)* Update initialization of the handler, service and repository
- *(currency-wallet)* Add additional argument output when handling handler errors
- *(currency-wallet)* Add a method that returns the user index from the database
- *(currency-wallet)* Add the GetUser method to the AuthRepository interface
- *(currency-wallet)* Add GenerateToken method for user login in the service layer
- *(currency-wallet)* Add a field discriminator to link the structure to the User table database
- *(currency-wallet)* Add a handler method for the /api/v1/login endpoint
- *(currency-wallet)* Add JWT token processing
- *(currency-wallet)* Add user identity middleware for authentication
- *(currency-wallet)* Add user authentication
- *(currency-wallet)* Add implementation of the ParseToken method of the AuthService interface
- *(currency-wallet)* Add a stub to check endpoint /api/v1/balance
- *(migrations)* Add the user_balances table
- *(repository)* Add Database interface and transaction method
- *(repository)* Implement registration and authentication requests through transactions
- *(repository)* Add GetBalance method to fetch user balance
- *(service)* Add GetBalance service method
- *(service)* Update initialization of BalanceService
- *(repository)* Add BalanceRepository interface
- *(middleware)* Add logging
- *(handler)* Add a balance handler
- Update Deposit
- *(handler)* Add a user account replenishment handler
- *(service)* Add business logic for account replenishment
- *(repository)* Add an interface for wallet operations
- *(repository)* Add interface initialization over wallet operations
- *(repository)* Add deposit functionality for wallets
- *(model)* Add UserID field to Wallet struct
- *(handler)* Improve Deposit method with error handling and balance response
- *(handler)* Enhance Deposit and Withdraw methods with improved error handling and logging
- *(service)* Add Withdraw method to WalletService
- *(repository)* Add Withdraw method to WalletRepository and implement in Postgres
- *(logger)* Add PrettyHandler for structured logging with color output
- *(exchange)* Add gRPC service for currency exchange
- Update dependencies and add new modules
- *(config)* Add support for config path via command line argument

### 🐛 Bug Fixes

- *(currency-wallet)* Fix the database port type
- *(currency-wallet)* Change the PortEndpoint field type from int to string
- *(currency-wallet)* Add port selection from the configuration file
- *(currency-wallet)* Change POST method to GET for endpoint /api/v1/login
- *(currency-wallet)* Change GET method to POST for endpoint /api/v1/login
- *(currency-wallet)* Fix the return value of the user registration method in all layers
- *(repository)* Pass the Database interface as an argument
- *(currency-wallet)* Update imports

### 🚜 Refactor

- *(currency-wallet)* Change repository structure and add init function to package main
- *(currency-wallet)* Move password hash and jwt-token generation to the security package
- *(logger)* Simplify response handling and logging functions
- *(currency-wallet)* Improve error handling and logging in authentication and wallet operations
- *(server)* Update log message and improve server exit handling
- *(server)* [**breaking**] Simplify config loading and improve logging
- *(logger)* Move the logger implementation to a common module
- *(assets)* Remove assets
- *(currency-wallet)* Simplify configuration logging

### 🎨 Styling

- *(currency-wallet)* Format code with go fmt

### ⚙️ Miscellaneous Tasks

- *(currency-wallet)* Update dependencies
- *(currency-wallet)* Define table constants
- *(currency-wallet)* Update go.mod and go.sum
- Update go.mod
- *(protos)* Update module path and add required dependencies

<!-- generated by git-cliff -->
