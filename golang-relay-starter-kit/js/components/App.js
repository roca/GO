import React from 'react';
import Relay from 'react-relay';

// Your React component
class App extends React.Component {
  render() {
    console.log(this.props.latestPost);
    return (
      <div>
        <div>
          {this.props.latestPost.id} <br/>
          <h1>{this.props.latestPost.text}</h1>
          {this.props.latestPost.author.name}
        </div>
        <h1>Current Author: {this.props.currentAuthor.name}</h1>
      </div>
    );
  }
}

// Your Relay container.
// Compose your React components with a declaration of
// the GraphQL query fragments that fetch their data.
export default Relay.createContainer(App, {
  fragments: {
    currentAuthor: () => Relay.QL `
          fragment on Author {
            id
            name
          }
    `,
    latestPost: () => Relay.QL`
      fragment on Post {
        id
        text
        author {
          id
          name
        }
      }
    `,
  },
});