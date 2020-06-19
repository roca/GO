const graphql = require('graphql');

const {
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLID,
    GraphQLString,
    GraphQLInt,
    GraphQLBoolean,
    GraphQLFloat,
    GraphQLNonNull
} = graphql


//Scalar Types
/*
    String      GraphQLString
    int         GraphQLInt
    Float       GraphQLFloat
    Boolean     GraphQLBoolean
    ID          GraphQLID

*/

const Person = new GraphQLObjectType({
    name: 'Person',
    description: 'Represents a Person Type',
    fields: () => ({
        id: {type: GraphQLID},
        name: {type: new GraphQLNonNull(GraphQLString)},
        age: {type: GraphQLInt},
        isMarried: {type: GraphQLBoolean},
        gpa: {type: GraphQLFloat},
        height: {type: new GraphQLNonNull(GraphQLFloat)},
        weight: {type: GraphQLFloat},
        bmi: {
            type: GraphQLFloat,
            resolve(parent,args){
                return parent.weight/(parent.height * parent.height); 
            }
        },
        justAType: {
            type: Person,
            resolve(parent,args){
                return parent; 
            }
        }
    })
})

// RootQuery
const RootQuery = new GraphQLObjectType({
    name: 'RootQueryType',
    description: 'Description',
    fields: () => ({
            person: {
                type: Person,
                resolve(parent,args){
                    let personObj = {
                        name: 'Antonio',
                        age: 35,
                        isMarried: true,
                        gpa: 4.0,
                        weight: 54.4311,
                        height: 1.6002
                    };
                    return personObj;
                }
            }
    })
});


module.exports = new GraphQLSchema({
    query: RootQuery,
});
