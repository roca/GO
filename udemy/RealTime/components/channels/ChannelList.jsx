import React, {Component} from 'react';
import Channel from './Channel.jsx';

class ChannelList extends Component {
    render() {
        const {channels, setChannel} = this.props;
        return (
            <ul>{
                channels.map( (channel,index) => {
                   
                       <Channel key={index} channel={channel} setChannel={setChannel} />
                   
                })
            }</ul>
        );
    }
}

ChannelList.propTypes = {
    channels: React.PropTypes.array.isRequired,
    setChannel: React.PropTypes.func.isRequired
}

export default ChannelList;
