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

    addChannel(namel) {
        console.log('name:' + name);
        let {channels} = this.state;
        channels.push({id: channels.length, name});
        this.setState({channels});
        // TODO: Send to server
    }

    render() {
        return (
            <ChannelSection 
                {...this.state} 
                setChannel={this.setChannel.bind(this)} 
                addChannel={this.addChannel.bind(this)}/>
        );
    }
}


export default App;