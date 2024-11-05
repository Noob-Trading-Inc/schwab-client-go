openssl req -x509 -nodes -new -sha256 -days 1024 -newkey rsa:2048 -keyout localhost.key -out localhost.pem -subj "/C=US/CN=LOCALHOST-CA"
openssl x509 -outform pem -in localhost.pem -out localhost.crt
rm localhost.pem