import React, {Component} from 'react';

class Channel extends React.Component{
    onClick() {
        e.preventDefault();
        const {setChannel, channel} = this.props;
        setChannel(channel);
    }
    render() {
        const {channel} = this.props;
        return(
            <li>
                <a onClick={this.onClick.bind(this)}>
                    {channel.name}
                </a>
            </li>
        )
    }
}

Channel.propTypes = {
    channel: React.ProTypes.object.isRequired,
    setChannel: React.ProTypes.func.isRequired
}

export default Channel;