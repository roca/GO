import React, {Component} from 'react';

class User extends Comment{
    render() {
        return (
            <li>
                {this.compareDocumentPosition.user.name}
            </li>
        );
    }
}

User.propTypes = {
    user: React.PropTypes.object.isRequired
};

export default User;