# Weeconn

A simple fast URL shortener for human.

## Technologies

* Go programming language
* Redis

## API Documentation

### URL Shortener

API to create a 7-Length ID for a given URL.

* **URL**

  `/link/shorten`

* **Method:**

  `POST`
  
* **Headers**

    `content-type` : `application/x-www-form-urlencoded`
  
* **Data Params**

   **Required:**
  
   `url=[string]` : Long URL to be shorten

* **Success Response:**
  
  * **Code:** 200 OK <br />
    **Content:** `{ long_url : "https://longurl.com/", id : "XXXXXXX" }`
 
* **Error Response:**  

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{ error : "bad url" }`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{ error : "internal server error" }`

* **Sample Call:**

    `curl -X POST -d "url=https://cloudflare.com/longlong" https://[ADDR:PORT]/link/shorten` 
  

### Short ID Redirect

Permanent Redirect short IDs to the actual URL.

* **URL**

  `/{ShortID}`

* **Method:**

  `GET`
  
* **Success Response:**
  
  * **Code:** 301 MOVED PERMANENTLY <br />
    **Location:** `[ACTUAL_URL]`
 
* **Error Response:**  

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "link not found" }`

  OR

  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `{ error : "internal server error" }`

* **Sample Call:**

    `curl https://[ADDR:PORT]/XXXXXXX` 
  
 
### Health Check
 
 Check service health.
 
 * **URL**
 
   `/status/health`
 
 * **Method:**
 
   `HEAD`
   
 * **Success Response:**
   
   * **Code:** 200 OK <br />
  
 * **Error Response:**
 
    No response
 
 * **Sample Call:**
 
     `curl -I https://[ADDR:PORT]/status/health` 
