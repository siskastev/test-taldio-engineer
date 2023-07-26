# Technical Test Taldio - API for xyz company Job Information
Welcome to the API repository for retrieving job information at xyz company.This repository contains the API for get information jobs open for hiring that will be consumed by FrontEnd Developer to display job openings on the company website or application.

## Project Scopes
* Get Master Employment Type: This endpoint allows you to retrieve the master list of employment types.

* Get Master Position Level: Use this endpoint to access the master list of position levels category.

* Get All Jobs in xyz company: This endpoint provides a list of all available jobs at xyz.

* Get Job Details by ID: Using this endpoint, you can retrieve detailed information about a specific job by providing its unique ID.

## Proof of System Running
To demonstrate that the system is up and running successfully, you can find the result images in the `result_image` folder. These images capture sample responses from various API endpoints and can be referred to as evidence of the system's functionality.

## SQL Database Structure
The SQL database structure used to store job-related information can be found in the `loker.sql` file. This file contains the necessary SQL queries and schema to create and populate the required tables.

## Installation and Setup
1. Clone the repository:
```bash
git clone https://github.com/siskastev/test-taldio-engineer.git
```
2. Navigate to the project directory svc-loker-be:
```bash
cd test-taldio-engineer/svc-loker-be
```
3. Rename or copy .env.example to .env
4. Run Docker Compose Up for build the system and wait until connection
```bash
docker-compose up
```
5. After connection, i provides postman collection in this project `Taldio.postman_collection.json`. This file allows you to quickly test the API endpoints when the Docker container is running.