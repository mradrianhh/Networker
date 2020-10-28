package main

/*
Networker is a server-client based networking package for easy communications between terminal-based programs.

The design-philosophy revolves around the idea of a dialogue - one request, one response.
Peer-to-peer.
Multiple peers-to-one peer.

Should consider splitting "Message" to one "Request"-object and one "Response"-object.
Or, make message a message parent-class with a messagetype parameter.
Request inherits it, setting the messagetype to "REQUEST".
Response inherits it, setting the messagetype to "RESPONSE".

Request structure:

<RequestCode|Carriage>

Response structure:

<ResponseCode|Carriage>
*/
