# Weeconn

A simple fast URL shortener for human with view count statistics.

## Technologies

* Go programming language
* Redis
* InfluxDB

## API Documentation

### URL Shortener

API to create a 7-Length ID for a given URL.

* **Method and URL**

    `POST` `/link/shorten`

* **Headers**

    **(Required)** `content-type` : `application/x-www-form-urlencoded`
  
* **Data Params**

    **(Required)** `url=[string]` : Long URL to be shorten

* **Success Response:**

    **Code:** `200 OK`; **Content:** `{"long_url" : [string], "id": [string]}`
 
* **Error Response:**  

    **Code:** `400 BAD REQUEST`; **Content:** `{ error : "bad url" }`

    OR

    **Code:** `500 INTERNAL SERVER ERROR`; **Content:** `{"error": "internal server error"}`

* **Sample Call:**
  
        curl -X POST -d "url=https://cloudflare.com/longlong" https://[ADDR:PORT]/link/shorten

### Short ID Redirect

Permanent Redirect short IDs to the actual URL.

* **Method and URL**

    `GET` `/{ShortID}`
  
* **Success Response:**
  
    **Code:** `301 MOVED PERMANENTLY`; **Location:** `<ACTUAL_URL>`
 
* **Error Response:**  

    **Code:** `404 NOT FOUND`; **Content:** `{"error" : "link not found"}`

    OR

    **Code:** `500 INTERNAL SERVER ERROR`; **Content:** `{"error": "internal server error"}`

* **Sample Call:**

    `curl https://[ADDR:PORT]/XXXXXXX` 
  
### View count
 
 Check an ID view count statistics.
 
* **Method and URL**
 
    `POST` `/status/id`
   
* **URL Params**
   
    **(Optional)** `period=[int]` : View count for last [period] days
   
* **Data Params**

    **(Required)** `id=[string]` : Short ID to check view count
   
* **Success Response:**
        
    **Code:** `200 OK`; **Content:** `{"id": [string], "period": [int], "seen": [int]}`
  
* **Error Response:**
 
    **Code:** `500 INTERNAL SERVER ERROR`; **Content:** `{"error": "internal server error"}`
 
* **Sample Call:**
 
        curl -X POST -d "id=xxxxxxx" https://[ADDR:PORT]/status/id=period=7

### Health Check
 
Check service health.
 
* **Method and URL**
 
    `HEAD` `/status/health`
   
* **Success Response:**
   
    **Code:** 200 OK
  
* **Error Response:**
 
    No response
 
* **Sample Call:**
 
     `curl -I https://[ADDR:PORT]/status/health` 
