
const express = require('express');
const graphqlHTTP = require('express-graphql');

const app = express();

app.use('/graphql',graphqlHTTP({
    graphiql: true
}));

app.listen(4000,() => { // http://localhost:4000/
    console.log('Listening for requests on my awesome port 4000');
});