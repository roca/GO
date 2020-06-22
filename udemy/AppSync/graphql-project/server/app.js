
require('dotenv').config()
const express = require('express');
const graphqlHTTP = require('express-graphql');
const mongoose = require('mongoose');
const cors = require('cors');

const schema = require('./schema/schema');
const testSchema = require('./schema/types_schema');

const app = express();

const uri = "mongodb+srv://" + process.env.DB_USER + ":" + encodeURI(process.env.DB_PASS) + "@cluster0-xetiz.mongodb.net/gq-course?retryWrites=true&w=majority";
mongoose.connect(uri, {useNewUrlParser: true, useUnifiedTopology: true });
const db = mongoose.connection;

db.on('error', console.error.bind(console, 'connection error:'));
db.once('open', function() {
    console.log("Mongo URI: " + uri);
    console.log('Yes! We are connected to the Mongo database.')
});


app.use(cors());
app.use('/graphql',graphqlHTTP({
    graphiql: true,
    schema: schema
}));

app.listen(4000,() => { // http://localhost:4000/
    console.log('Listening for requests on my awesome port 4000');
});

