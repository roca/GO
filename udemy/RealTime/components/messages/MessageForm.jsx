import React, {Component} from 'react';

class MessageForm extends Component {
    onSubmit(e) {
        const {addMessage} = this.props;
        e.preventDefault();
        const node = this.refs.message;
        const messageName = node.value;
        addMessage(messageName);
        node.value = '';
    }
    render() {
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
            <div className='form-group'>
                <input  
                    className='form-control'
                    placeholder='Add Message'
                    type='text'
                    ref='message'
                />
            </div>
            </form>
        );
    }
}

MessageForm.propTypes = {
    addMessage: React.PropTypes.func.isRequired
}

export default MessageForm;