# Recommendation Systems

This repository contains two separate implementations of a recommendation system: one in Python and another in Go. Both systems use a dataset of video game reviews to recommend similar items based on user co-occurrence. The project includes RESTful APIs that allow users to request recommendations for specific items (ASINs).

## Table of Contents
- [Project Structure](#project-structure)
- [Data Source](#data-source)
- [Getting Started](#getting-started)
    - [Python Implementation](#python-implementation)
    - [Go Implementation](#go-implementation)
- [Usage](#usage)
    - [Python API](#python-api)
    - [Go API](#go-api)
- [Configuration](#configuration)
- [License](#license)

## Project Structure
```
recommendation-systems/
├── python-recommender/
│   ├── app.py
│   ├── requirements.txt
│   ├── Dockerfile
│   ├── data/
│   │   └── Video_Games.jsonl
│   └── models/
├── go-recommender/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── data/
│   │   └── Video_Games.jsonl
│   └── models/
└── README.md
```

## Data Source

The dataset used in this project is the **Video Game Reviews** dataset from the Amazon product review dataset collection. It is provided in JSON Lines format (`.jsonl`) and contains fields such as `user_id`, `asin`, and `rating`.

### Example Data Format

Each line in the `Video_Games.jsonl` file is a separate JSON object:

```json
{"user_id": "A3SGXH7AUHU8GW", "asin": "B00004SB92", "rating": 5.0, "title": "Great Game!", "text": "I loved this game. It was really fun to play.", "timestamp": "2023-01-01"}
{"user_id": "A1CQ7RK4P0DLT1", "asin": "B00004SB92", "rating": 4.0, "title": "Good Game", "text": "Enjoyable gameplay but a bit short.", "timestamp": "2023-01-02"}
```
+ **user_id :** The ID of the user who wrote the review.
+ **asin :** The Amazon Standard Identification Number (ASIN) of the product.
+ **rating :** The rating given by the user (e.g., 1-5 stars).
### Data Download
The data can be downloaded from the [UCSD Amazon Product Review Dataset](https://cseweb.ucsd.edu/~jmcauley/datasets.html#amazon_reviews) page. You may need to extract the specific category or subset you are interested in, such as “Video Games.”

## Getting Started

### Python Implementation

#### Prerequistes
+ **Python 3.6:** Ensure you have Python installed on your system.
+ **Docker:** Docker is used to containerize the application.
#### Installation
1. **Navigate to the Python Directory**
```bash
cd recommendation-systems/python-recommender 
```
2. **Build the Docker Image**
Use Docker to build the Python application image:
```bash
docker build -t python-recommender .
```
3. **Run the Docker Container**
To run the Python application using Docker:
```bash
docker run -p 5000:5000 python-recommender
```
The application will start and listen for HTTP requests on port 5000.

### Go Implementation
### Prerequistes
	
+ **Go 1.18+:** Make sure you have Go installed on your system. You can download it from the official Go website.
+ **Docker:** Docker is used to containerize the application. You can install Docker from the official Docker website.

### Installation

1. **Navigate to the Go Directory**
```bash
cd recommendation-systems/go-recommender
```
2. **Initiatize the Go Module**
Ensure your Go module is initialized:
```bash
go mod init github.com/mackcesar/recommendation-systems/go-recommender
```
3. **Build the Docker Image**
Use Docker to build the Go application image:
```bash
docker build -t go-recommender .
```
4. **Run the Docker Container**
```bash
docker run -p 8080:8080 go-recommender
```
The application will start and listen for HTTP requests on port 8080.

### Usage
### Python API
##### Python API Endpoint
The recommendation API endpoint is accessible at:
```bash
GET /recommend/{asin}
```
Replace {asin} with the actual ASIN of the product you want recommendations for.
**Response**
```json
{"recommendations": ["B00005N5PF", "B00002E328", "B00004U5VK", "B00005N5PM", "B00004YKZT"]}
```

### Go API
#### Go API Endpoint
The recommendation API endpoint is accessible at:

```bash
GET /recommend/{asin}
```
Replace {asin} with the actual ASIN of the product you want recommendations for.

#### Example Request
You can use curl to test the API:
```bash
curl http://localhost:8080/recommend/B00004SB92
```
#### Response:
```json
{
  "recommendations": ["B00005N5PF", "B00002E328", "B00004U5VK", "B00005N5PM", "B00004YKZT"]
}
```
### Configuration

Changing the Data Source

If you wish to change the data source for either implementation, follow these steps:

1.	**Download the Desired Dataset**
Obtain the JSON Lines dataset you want to use. Make sure it follows a similar structure to the example data format.

2. **Place the Dataset**
Replace the existing dataset file (data/Video_Games.jsonl) with your new dataset file. Ensure the file is in the data/ directory and matches the expected format.

3. **Update the Code (if needed)**
If the JSON structure of your new data differs, you may need to update the Review struct in the Go code or the data loading logic in the Python code to match the new data fields


### License
This project is licensed under the [MIT License](https://github.com/git/git-scm.com/blob/main/MIT-LICENSE.txt). See the LICENSE file for more information.
