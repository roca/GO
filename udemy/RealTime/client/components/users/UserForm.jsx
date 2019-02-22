import React, {Component} from 'react';

class UserForm extends Component {
    onSubmit(e) {
        const {setUserName} = this.props;
        e.preventDefault();
        const node = this.refs.user;
        const userName = node.value;
        setUserName(userName);
        node.value = '';
        //node.disabled = true;
    }
    render() {
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
            <div className='form-group'>
                <input  
                    className='form-control'
                    placeholder='Add User'
                    type='text'
                    ref='user'
                />
            </div>
            </form>
        );
    }
}

UserForm.propTypes = {
    setUserName: React.PropTypes.func.isRequired
}

export default UserForm;