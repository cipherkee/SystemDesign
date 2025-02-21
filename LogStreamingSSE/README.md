# SSE
1. Only for text. Video and audio cannot be supported
2. Unidirectional - only server can send using this protocol
3. Persistent connection with unidirectional. SSE is a separate protocol on top of HTTP
4. Eg: 
    1. stock market ticker
    2. Deployment logs
    3. Advertisements
    4. Updates on twitter feeds

**How to run**
1. Start the main function. This starts the server.
2. Open the html file on web browser. This listens to the messages from the server.
3. You should see the timestamps as logs
