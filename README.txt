1)- Install Openssl in desktop (https://slproweb.com/products/Win32OpenSSL.html)

2)- Add bin path into enviroment variable.

3)- openssl req -x509 -days 365 -newkey rsa:2048 -keyout my-key.pem -out my-cert.pem

	* Enter password
	* Verify Password
	* Enter all the detail as per they say

4)- PKCS file command openssl pkcs12 -export -in my-cert.pem -inkey my-key.pem -out akash-test-cert.pfx

5)- Enter password that you have used before.

6)- Create new Password.