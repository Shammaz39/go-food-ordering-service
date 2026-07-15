# Go Food Ordering Service



A backend food ordering service built with Go to demonstrate REST API development, database persistence, asynchronous processing, Kafka event publishing, and containerization.



## Tech Stack



- Go

- Fiber

- GORM

- PostgreSQL

- Apache Kafka

- Docker

- Docker Compose



## Features



- Create food orders through REST APIs

- Retrieve all orders

- Retrieve an order with its status timeline

- Persist order and order event data using PostgreSQL and GORM

- Generate unique order IDs using UUID

- Process order status progression asynchronously using goroutines

- Publish order and status events to Apache Kafka

- Configure database and Kafka connectivity using environment variables

- Run application infrastructure using Docker Compose



## Order Lifecycle



Orders progress through the following statuses:



PLACED → PREPARING → COOKING → OUT_FOR_DELIVERY → DELIVERED



Order status progression is processed asynchronously using a Go goroutine.



## REST APIs



### Create Order



POST /api/v1/orders



Example request:



&#x20;   {

&#x20;     "customer_name": "Ibrahim",

&#x20;     "address": "Bangalore",

&#x20;     "item": "Burger",

&#x20;     "size": "Large"

&#x20;   }



### Get All Orders



GET /api/v1/orders



### Get Order by ID



GET /api/v1/orders/{id}



The response includes the order details and status timeline.



## Project Structure



&#x20;   handlers/   - HTTP request handlers

&#x20;   kafka/      - Kafka producer implementation

&#x20;   models/     - GORM data models

&#x20;   services/   - Business logic and asynchronous order processing

&#x20;   main.go     - Application bootstrap and route configuration



## Environment Variables



The application uses the following environment variables:



&#x20;   DATABASE_URL

&#x20;   PORT

&#x20;   KAFKA_BROKER

&#x20;   KAFKA_TOPIC



Refer to `.env.example` for sample local configuration.



## Running with Docker



Build and start the services:



&#x20;   docker compose up --build



The application is available at:



&#x20;   http://localhost:3000



## Go Concepts Used



- Structs

- Pointers

- Packages

- Goroutines

- Error handling

- Slices

- Functions

- Modular package organization



## Note



This project was built to strengthen hands-on Go backend development skills and explore asynchronous processing and event publishing using Kafka.

