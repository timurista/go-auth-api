## Theory

jwt === json web token
jwt means of exchanging info between 2 parties. Digitally signed. -- verifiable.

```
{ Base64 encoded Header }.{ Base64 encoded Payload }.{ Signature }
```

payload claims
- regisreted
- public
- private

in our case payload will carry claims of email
but it can also carry expiration

`https://tools.ietf.org/html/rfc7519` is the location for good information of jwt token

Signatue from header payload and a secret. Process is digitally signing a secret string. Signature cannot be decrypted without secret string.




