# Pleiades Backend

## Setup

Install MongoDB and note the URI.

Create the .env file based on the .env.example file.

Environment Variables:
<table>
  <tr>
   <td>Variable</td>
   <td>Example</td>
   <td>Purpose</td>
  </tr>
  <tr>
   <td>MONGODB_URI</td>
   <td>mongodb://localhost:27017</td>
   <td>The URI of the MongoDB server to connect to.</td>
  </tr>
  <tr>
   <td>HOST</td>
   <td>localhost:5173/BASE_PATH</td>
   <td>The URL of the frontend.</td>
  </tr>
  <tr>
   <td>PROTOCOL</td>
   <td>http://</td>
   <td>The protocol to use when connecting to the frontend.</td>
  </tr>
  <tr>
   <td>MJ_APIKEY_PUBLIC</td>
   <td>abcde...</td>
   <td>API Key for Mailjet account</td>
  </tr>
  <tr>
   <td>MJ_APIKEY_PRIVATE</td>
   <td>abcde...</td>
   <td>API Secret Key for Mailjet account</td>
  </tr>
</table>

## Runnning
`go run .`

## Testing

`go test .`