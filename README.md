# yarfu

Yet another rest file uploader

rest api microservice that accepts image file uploads

currently only jpeg, png and gif formats are supported

 - all uploaded images can be reached via /public/filename url
 - all valid uploaded images will have 100x100 thumb image via /public/thumb\_filename url
 - multiple files, multipart/form-data requests supported
 - JSON requests with BASE64 encoded images requests supported
 - uploading by url supported
 - Dockerfile included

service can be tested via curl:

curl -vvv -X POST http://127.0.0.1:8080/api/v1/images -F "files[]=@1.png" -F "files[]=@2.png" -H "Content-Type: multipart/form-data"

(where 127.0.0.1:8080 is a service address)

assets folder contains valid file for testing uploads via  BASE64/JSON

valid json can be generated by hand and should have format:

{ "name": "filename.ext", "image64": "base64_encoded_string" }

or my testing utility github.com/sibyakin/base64er can be used

generated file can be uploaded with curl:

curl -vvv -X POST -H "Content-Type: application/json" --data @doge.json http://127.0.0.1:8080/api/v1/images/json

uploading by url example:

curl -vvv -X GET "http://127.0.0.1:8080/api/v1/images/url?url=https://pythonprogramming.net/static/images/mainlogowhitethick.jpg"
