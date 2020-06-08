const graphql = require('graphql');

const {
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLID,
    GraphQLString,
    GraphQLInt
} = graphql

// Create types
const UserType = new GraphQLObjectType({
   name: 'User',
   description: 'Documentation for user...',
   fields: () => ({
       id: {type: GraphQLString},
       name: {type: GraphQLString},
       age: {type: GraphQLInt}
   })
});


// RootQuery

const RootQuery = new GraphQLObjectType({
    name: 'RootQueryType',
    description: 'Description',
    fields: () => ({
        user: {
            type: UserType,
            args: {
                id: {type: GraphQLString}
            },
            resolve(parent, args) {
                 let user = {
                     id: '345',
                     name: 'John',
                     age: 25
                 };
                 
                 // we resolve with data
                 // get and return data from a data source
                 return user
            }
        }
    })
});

module.exports = new GraphQLSchema({
    query: RootQuery   
});