import React, {Component} from 'react';
import ChannelSection from './channels/ChannelSection.jsx';

class App extends Component{
    constructor(props){
        super(props);
        this.state = {
            channels: []
        };
    }

    setChannel(activeChannel) {
        this.setState({activeChannel});
        // TODO: Get Channels Message
    }   

    addChannel(name) {
        let {channels} = this.state;
        channels.push({id: channels.length, name});
        this.setState({channels});
         // TODO: Send to server
    }

    render() {
        return (
            <ChannelSection 
                channels={this.state.channels} 
                setChannel={this.setChannel.bind(this)} 
                addChannel={this.addChannel.bind(this)}/>
        );
    }
}


export default App;