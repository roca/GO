import React from 'react';
import Relay from 'react-relay';

// Your React component
class App extends React.Component {
  render() {
    console.log(this.props.latestPost);
    return (
      <div>
        {this.props.latestPost.id}
        <h1>{this.props.latestPost.text}</h1>
        <h1>{this.props.latestPost.author.name}</h1>
      </div>
    );
  }
}

// Your Relay container.
// Compose your React components with a declaration of
// the GraphQL query fragments that fetch their data.
export default Relay.createContainer(App, {
  fragments: {
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