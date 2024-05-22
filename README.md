# Go URL Shortener

 URL shortener project using Golang and MySQL.
 
 to run `main.go` in your terminal and server will run in `http://localhost:5000/`.

```bash
go run main.go

## Usage

Open your favorite REST API Tester such as [Postman](https://www.postman.com/), [Insomnia](https://insomnia.rest/), [Paw](https://paw.cloud/), [cURL](https://curl.se/), etc.
### Generate Short URL From Input
Set URL to `http://localhost:5000/` and METHOD to `POST`. Send this JSON type in body, in this example we'll shorten [Facebook](https://www.facebook.com/) URL.

```json
{
    "long_url" : "https://www.google.com/"
}
```
![Cuplikan](https://github.com/firlydani18/Url_shorter/assets/52233444/431f5da7-2bcd-4259-9507-6a34c5abfd06)


### Getting Original URL From Short URL

Set URL to `http://localhost:5000/` and METHOD to `GET`. Send this JSON type in body, in this example we'll getting `http://localhost:5000/HUTSntZ` original URL.

```json
{
    "short_url" : "http://localhost:5000/HUTSntZ"
}
```
<img width="554" alt="Cuplikan layar 2024-05-22 183725" src="https://github.com/firlydani18/Url_shorter/assets/52233444/f5357ea1-3105-4229-8039-3ef920232d99">
