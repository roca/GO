ionic cordova plugin remove cordova-plugin-googleplus --save \
--variable REVERSED_CLIENT_ID=com.googleusercontent.apps.808648518409-h2macpr6ndoe1bn78trmdslf5jd6mg8m \
--variable WEB_APPLICATION_CLIENT_ID=808648518409-hj2d7c5gk1fc57d4dtulucg6a8qsuhh6.apps.googleusercontent.com


ionic cordova plugin add cordova-plugin-googleplus --save \
--variable REVERSED_CLIENT_ID=com.googleusercontent.apps.808648518409-h2macpr6ndoe1bn78trmdslf5jd6mg8m \
--variable WEB_APPLICATION_CLIENT_ID=808648518409-hj2d7c5gk1fc57d4dtulucg6a8qsuhh6.apps.googleusercontent.com

ionic cordova  run android --target=Pixel_API_24  -l -c --debug

https://console.firebase.google.com/

keytool -list -v -alias androiddebugkey -keystore ~/.android/debug.keystore

https://accounts.google.com/.well-known/openid-configuration
"jwks_uri": "https://www.googleapis.com/oauth2/v3/certs"

openssl s_client -showcerts -connect www.googleapis.com:443
openssl x509 -in certificate.crt -fingerprint -noout

ionic cordova build ios -- --buildFlag="-UseModernBuildSystem=0"

