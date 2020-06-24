const graphql = require('graphql');
var _ = require('lodash');
const User = require('../models/user');
const Post = require('../models/post');
const Hobby = require('../models/hobby');

var RemoveItem = (mObj,args) => {
    let item = mObj.findByIdAndRemove(args.id).exec();
    if(!item){
        throw new Error("Error deleting Item");
    }
    return item;
}

var createNewUserItem = (args) => {
    let user = new User({
        name: args.name,
        age: args.age,
        profession: args.profession
    });
    user.save();
    return user;
}

var updateUserItem = (args) => {
    let updatedUser = User.findByIdAndUpdate(
        args.id,
        {
            $set: {
                name: args.name,
                age: args.age,
                profession: args.profession
            }
        },
        {new: true} // send back the updated objectType
    );
    return updatedUser;
}

var createNewPostItem = (args) => {
    let post = new Post({
        comment: args.comment,
        userId: args.userId
    });
    post.save();
    return post;
}
var updatePostItem = (args) => {
    let updatedPost = Post.findByIdAndUpdate(
        args.id,
        {
            $set: {
                comment: args.comment,
                userId: args.userId
            }
        },
        {new: true} // send back the updated objectType
    );
    return updatedPost;
}

var createNewHobbyItem = (args) => {
    let hobby = new Hobby({
        title: args.title,
        description: args.description,
        userId: args.userId
    });
    hobby.save();
    return hobby;
}
var updateHobbyItem = (args) => {
    let updatedHobby = Hobby.findByIdAndUpdate(
        args.id,
        {
            $set: {
                title: args.title,
                description: args.description,
                userId: args.userId
            }
        },
        {new: true} // send back the updated objectType
    );
    return updatedHobby;
}

// dummy data
// var userData = [
//     {id: '1', name: 'Bond', age:36, profession: 'Programmer'},
//     {id: '13', name: 'Anna', age:26, profession: 'Baker'},
//     {id: '211', name: 'Bella', age:16, profession: 'Medical Doctor'},
//     {id: '19', name: 'Gina', age:26, profession: 'Painter'},
//     {id: '150', name: 'Georgina', age:36, profession: 'Teacher'}
// ];

// var hobbyData = [
//    {id: '1', title: 'Programing', description: 'Using computer to make the world a better place', userId: '150'},
//    {id: '2', title: 'Rowing', description: 'Sweat and feel better before eating donuts', userId: '211'},
//    {id: '3', title: 'Swimming', description: 'Get in the water and learn to become the water', userId: '211'},
//    {id: '4', title: 'Fencing', description: 'A hobby for fancy people', userId: '13'},
//    {id: '5', title: 'Hiking', description: 'Wear hiking boot and explore the world', userId: '150'}
// ]

// var postData = [
//     {id: '1', comment: 'Building a Mind', userId: '1'},
//     {id: '2', comment: 'GraphQL is Amazing', userId: '1'},
//     {id: '3', comment: 'How to Change the World', userId: '19'},
//     {id: '4', comment: 'How to Change the World', userId: '211'},
//     {id: '5', comment: 'How to Change the World', userId: '1'}
// ]

const {
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLID,
    GraphQLString,
    GraphQLInt,
    GraphQLList,
    GraphQLNonNull
} = graphql

// Create types
const UserType = new GraphQLObjectType({
   name: 'User',
   description: 'Documentation for user...',
   fields: () => ({
       id: {type: GraphQLID},
       name: {type: GraphQLString},
       age: {type: GraphQLInt},
       profession: {type: GraphQLString},
       posts: {
           type: new GraphQLList(PostType),
           resolve: (parent, args) =>  Post.find({userId: parent.id})
       },
       hobbies: {
           type: new GraphQLList(HobbyType),
           resolve: (parent, args) =>  Hobby.find({userId: parent.id})
       }
   })
});

const HobbyType = new GraphQLObjectType({
    name: 'Hobby',
    description: 'Hobby description',
    fields: () => ({
        id: {type: GraphQLID},
        title: {type: GraphQLString},
        description: {type: GraphQLString},
        user: {
            type: UserType,
            resolve: (parent, args) =>  User.findById(parent.userId)
        }
    })
});

const PostType = new GraphQLObjectType({
    name: 'Post',
    description: 'Post description',
    fields: () => ({
        id: {type: GraphQLID},
        comment: {type: GraphQLString},
        user: {
            type: UserType,
            resolve: (parent, args) =>  User.findById(parent.userId)
        }
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
            resolve: (parent, args) =>  User.findById({id:args.id})
        },
        users: {
            type: new GraphQLList(UserType),
            resolve: (parent,args) => User.find({})
        },
        hobby: {
            type: HobbyType,
            args: {
                id: {type: GraphQLID}
            },
            resolve: (parent, args) =>  Hobby.findById(args.id)
        },
        hobbies: {
            type: new GraphQLList(HobbyType),
            resolve: (parent,args) => Hobby.find({})
        },
        post: {
            type: PostType,
            args: {
                id: {type: GraphQLID}
            }, 
            resolve(parent,args) { return Post.findById(args.id)}
        },
        posts: {
            type: new GraphQLList(PostType),
            resolve: (parent,args) => Post.find({})
        }
    })
});

//Mutations
const Mutation = new GraphQLObjectType({
    name: 'Mutation',
    fields: {
        CreateUser: {
            type: UserType,
            args: {
                name: {type: new GraphQLNonNull(GraphQLString)},
                age: {type: new GraphQLNonNull(GraphQLInt)},
                profession: {type: GraphQLString}
            },
            resolve: (parent,args) => createNewUserItem(args)
        },
        UpdateUser: {
            type: UserType,
            args: {
                id: {type: new GraphQLNonNull(GraphQLID)},
                name: {type: new GraphQLNonNull(GraphQLString)},
                age: {type: GraphQLInt},
                profession: {type: GraphQLString}
            },
            resolve: (parent,args) => updateUserItem(args)
        },
        RemoveUser: {
            type: UserType,
            args: {
                id: {type: new GraphQLNonNull(GraphQLID)}
            },
            resolve: (parent,args) => RemoveItem(User,args)
        },
        CreatePost: {
            type: PostType,
            args: {
                comment: {type: new GraphQLNonNull(GraphQLString)},
                userId: {type: new GraphQLNonNull(GraphQLID)}
            },
            resolve: (parent,args) => createNewPostItem(args)
        },
        UpdatePost: {
            type: PostType,
            args: {
                id: {type: new GraphQLNonNull(GraphQLID)},
                comment: {type: new GraphQLNonNull(GraphQLString)},
                userId: {type: new GraphQLNonNull(GraphQLID)}
            },
            resolve: (parent,args) => updatePostItem(args)
        },
        RemovePost: {
            type: PostType,
            args: {
                id: {type: new GraphQLNonNull(GraphQLID)}
            },
            resolve: (parent,args) => RemoveItem(Post,args)
        },
        CreateHobby: {
            type: HobbyType,
            args: {
                title: {type: new GraphQLNonNull(GraphQLString)},
                description: {type: GraphQLString},
                userId: {type: new GraphQLNonNull(GraphQLID)}
            },
            resolve: (parent,args) => createNewHobbyItem(args)
        },
        UpdateHobby: {
            type: HobbyType,
            args: {
                id: {type: new GraphQLNonNull(GraphQLID)},
                title: {type: new GraphQLNonNull(GraphQLString)},
                description: {type: new GraphQLNonNull(GraphQLString)},
                userId: {type: new GraphQLNonNull(GraphQLID)}
            },
            resolve: (parent,args) => updateHobbyItem(args)
        },
        RemoveHobby: {
            type: HobbyType,
            args: {
                id: {type: new GraphQLNonNull(GraphQLID)}
            },
            resolve: (parent,args) => RemoveItem(Hobby,args)
        }
    }
});

module.exports = new GraphQLSchema({
    query: RootQuery,
    mutation: Mutation   
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
  hobby(id: 3){
    title
  }
}

query q3{
  post(id: "1"){
    comment
    user{
      name
    }
  }
}

mutation m1{
  createUser(name: "John",age: 25,profession: "Barber"){
    id
    name
    age
    profession
    hobbies {
      title
    }
  }
}

mutation m1 {
  CreateUser(name: "Mabondu",age: 78, profession: "Father") {
    id
    name
    age
    profession
  }
}

mutation m2 {
  CreatePosts(comment: "This is cool",userId: "5eecd2c0fc41853930576ecb") {
    comment
    id
  }
}


mutation m3 {
  CreateHobby(title: "Cyclist", description: "Roadie", userId: "5eecd2c0fc41853930576ecb") {
    id
    title
  }
}

query q1{
  post(id: "5eecd413592edb39a8933396"){
    id
    comment
  }
}

query q2 {
  hobbies {
    id
    title
    user{
      name
      hobbies{
        title
      }
    }
  }
}

query q3 {
  hobby(id: "5eecd4c9592edb39a8933397") {
    title
  }
}
*/