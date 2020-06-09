const graphql = require('graphql');
var _ = require('lodash');

var getItemByID = (objArray,id) => {
    return _.find(objArray,{id: id})
}

// dummy data
var userData = [
    {id: '1', name: 'Bond', age:36, profession: 'Secret Agent'},
    {id: '13', name: 'Anna', age:26},
    {id: '211', name: 'Bella', age:16, profession: 'Medical Doctor'},
    {id: '19', name: 'Gina', age:26},
    {id: '150', name: 'Georgina', age:36}
];

var hobbyData = [
   {id: '1', title: 'Programing', description: 'Using computer to make the world a better place'},
   {id: '2', title: 'Rowing', description: 'Sweat and feel better before eating donuts'},
   {id: '3', title: 'Swimming', description: 'Get in the water and learn to become the water'},
   {id: '4', title: 'Fencing', description: 'A hobby for fancy people'},
   {id: '5', title: 'Hiking', description: 'Wear hiking boot and explore the world'}
]

var postData = [
    {id: '1', comment: 'Building a Mind'},
    {id: '2', comment: 'GraphQL is Amazing'},
    {id: '3', comment: 'How to Change the World'}
]

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
       id: {type: GraphQLID},
       name: {type: GraphQLString},
       age: {type: GraphQLInt},
       profession: {type: GraphQLString}
   })
});

const HobbyType = new GraphQLObjectType({
    name: 'Hobby',
    description: 'Hobby description',
    fields: () => ({
        id: {type: GraphQLID},
        title: {type: GraphQLString},
        description: {type: GraphQLString}
    })
});

const PostType = new GraphQLObjectType({
    name: 'Post',
    description: 'Post description',
    fields: () => ({
        id: {type: GraphQLID},
        comment: {type: GraphQLString}
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
                id: {type: GraphQLID}
            },
            resolve(parent, args) {
                 // we resolve with data
                 // get and return data from a data source
                 return _.find(userData,{id: args.id})
            }
        },
        hobby: {
            type: HobbyType,
            args: {
                id: {type: GraphQLID}
            },
            resolve: (parent, args) =>  getItemByID(hobbyData,args.id)
        },
        post: {
            type: PostType,
            args: {
                id: {type: GraphQLID}
            }, 
            resolve(parent,args) { return getItemByID(postData,args.id)}
        }

    })
});

module.exports = new GraphQLSchema({
    query: RootQuery   
});

/*
query q1{
  user(id: "1"){
    id,
    age,
    name,
    profession

  }
}

query q2{
  hobby(id: 1){
    title
  }
}
*/