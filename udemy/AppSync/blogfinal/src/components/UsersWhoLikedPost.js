import React , { Component } from 'react';


class UsersWhoLikedPost extends Component {

     render() {
        const allUsers = this.props.data
        return allUsers.map((user) => {
            return (
                    <div key={user}>
                        <span style={{fontStyle:"bold", color:"#ged"}}>
                            <a href={"mailto:" + user}>{user}</a>
                        </span>
                    </div>
            )

        })
    }
}


export default UsersWhoLikedPost;