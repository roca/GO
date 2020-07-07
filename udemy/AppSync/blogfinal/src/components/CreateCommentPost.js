import React , { Component } from 'react';
import { API, graphqlOperation, Auth} from 'aws-amplify';
import { createComment } from '../graphql/mutations';

class CreateCommentPost extends Component {

    state = {
        commentOwnerId: "",
        commentOwnerUsername: "",
        content: ""
    }

    componentWillMount = async () => {
        await Auth.currentUserInfo()
                .then(user => {
                    this.setState({
                        commentOwnerId: user.attributes.sub,
                        commentOwnerUsername: user.username
                    })
                })
    }

    handleChangeContent = event => this.setState({content: event.target.value})

    handleAddComment = async event => {
        event.preventDefault();

        const input = {
            commentPostId: this.props.postId,
            commentOwnerId: this.state.commentOwnerId,
            commentOwnerUsername: this.state.commentOwnerUsername,
            content: this.state.content,
            createdAt: new Date().toISOString()
        }

        await API.graphql(graphqlOperation(createComment, { input }));

        this.setState({ content: ""}); // clear field
    }

    render() {
        return (
            <form className="add-comment"
                onSubmit={(event) => this.handleAddComment(event)} >
                <textarea
                    type="text" placeholder="Add your comment to this post"
                    name="content"
                    rows="3"
                    cols="40"
                    required
                    value={this.state.content}
                    onChange={(event) => this.handleChangeContent(event)}/>
                <input 
                    className="btn" 
                    style={{ fontSize: '19px'}}
                    type="submit" 
                    value="Add Comment"/>
            </form>
        )
    }
}

export default CreateCommentPost;