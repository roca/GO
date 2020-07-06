import React , { Component } from 'react';
import { API, graphqlOperation, Auth} from 'aws-amplify';
import { updatePost } from '../graphql/mutations';


class EditPost extends Component {

    state = {
        show: false,
        id: "",
        postOwnerId: "",
        postOwnerUsername: "",
        postTitle: "",
        postBody: "",
        postData: {
            postTitle: this.props.postTitle,
            postBody: this.props.postBody
        }
    }

    handleModal = () => {
        this.setState({ show: !this.state.show });
        document.body.scrollTop = 0;
        document.documentElement.scrollTop = 0;

    }

    handleChangePostTitle = event => this.setState({
       postData: {...this.state.postData, postTitle: event.target.value} 
    })
    handleChangePostBody = event => this.setState({
       postData: {...this.state.postData, postBody: event.target.value} 
    })

    handleUpdatePost = async event => {
        event.preventDefault();

        const input = {
            id: this.props.id,
            postOwnerId: this.state.postOwnerId,
            postOwnerUsername: this.state.postOwnerUsername,
            postTitle: this.state.postData.postTitle,
            postBody: this.state.postData.postBody
        }

        await API.graphql(graphqlOperation(updatePost, { input }));

        this.setState({ show: !this.state.show });
    }


    componentWillMount = async () => {
        await Auth.currentUserInfo()
                .then(user => {
                    this.setState({
                        postOwnerId: user.attributes.sub,
                        postOwnerUsername: user.username
                    })
                })
    }


     render() {
        return (
            <>
                { this.state.show && (
                    <div className="modal">
                        <button className="close" onClick={this.handleModal}>X</button>

                        <form className="add-post"
                            onSubmit={(event) => this.handleUpdatePost(event)} >
                            <input style={{ font: '19px'}}
                                type="text" placeholder="Title"
                                name="postTitle"
                                value={this.state.postData.postTitle}
                                onChange={(event) => this.handleChangePostTitle(event)}/>
                            <input style={{ height: "150px", fontSize: "19px"}}
                                type="text" placeholder="New Blog Post"
                                name="postBody"
                                value={this.state.postData.postBody}
                                onChange={(event) => this.handleChangePostBody(event)}/>
                            <input className="btn" style={{ fontSize: '19px'}}
                                type="submit" value="update post"/>
                        </form>

                    </div>
                )
                }
                <button onClick={this.handleModal}>Edit</button>
            </>
        );
    }
}


export default EditPost;